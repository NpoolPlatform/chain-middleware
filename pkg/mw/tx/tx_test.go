package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.Tx{
	CoinUnit:      "BTC",
	CoinLogo:      uuid.NewString(),
	CoinENV:       "test",
	FromAccountID: uuid.NewString(),
	ToAccountID:   uuid.NewString(),
	Amount:        "123.1",
	FeeAmount:     "2.01",
	State:         basetypes.TxState_TxStateCreated,
	StateStr:      basetypes.TxState_TxStateCreated.String(),
	Type:          basetypes.TxType_TxWithdraw,
	TypeStr:       basetypes.TxType_TxWithdraw.String(),
	Extra:         uuid.NewString(),
}

var req = &npool.TxReq{
	FromAccountID: &ret.FromAccountID,
	ToAccountID:   &ret.ToAccountID,
	Amount:        &ret.Amount,
	FeeAmount:     &ret.FeeAmount,
	State:         &ret.State,
	Type:          &ret.Type,
	Extra:         &ret.Extra,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithID(&ret.CoinTypeID),
		coin1.WithName(&ret.CoinName),
		coin1.WithUnit(&ret.CoinUnit),
		coin1.WithLogo(&ret.CoinLogo),
		coin1.WithENV(&ret.CoinENV),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithCoinTypeID(req.CoinTypeID),
		WithFromAccountID(req.FromAccountID),
		WithToAccountID(req.ToAccountID),
		WithAmount(req.Amount),
		WithFeeAmount(req.FeeAmount),
		WithChainTxID(req.ChainTxID),
		WithState(req.State),
		WithExtra(req.Extra),
		WithType(req.Type),
	)
	assert.Nil(t, err)

	info, err := handler.CreateTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.State = basetypes.TxState_TxStateWait
	ret.StateStr = ret.State.String()
	req.State = &ret.State

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithChainTxID(req.ChainTxID),
		WithState(req.State),
		WithExtra(req.Extra),
		WithType(req.Type),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateTx(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}

	ret.State = basetypes.TxState_TxStateWait
	ret.StateStr = ret.State.String()
	req.State = &ret.State

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithChainTxID(req.ChainTxID),
		WithState(req.State),
		WithExtra(req.Extra),
		WithType(req.Type),
	)
	assert.Nil(t, err)

	_, err = handler.UpdateTx(context.Background())
	assert.NotNil(t, err)
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
}
