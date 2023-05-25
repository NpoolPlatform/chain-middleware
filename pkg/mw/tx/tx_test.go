package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	coinmw "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

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

var (
	coinName = "BTC111"
	coinUnit = "BTC"
	coinENV  = "test"
)

var coinReq = &coinmwpb.CoinReq{
	Name: &coinName,
	Unit: &coinUnit,
	ENV:  &coinENV,
}

var ret = &npool.Tx{
	CoinName:      coinName,
	CoinUnit:      coinUnit,
	CoinENV:       coinENV,
	FromAccountID: uuid.NewString(),
	ToAccountID:   uuid.NewString(),
	Amount:        "123.1",
	FeeAmount:     "2.01",
	State:         txmgrpb.TxState_StateCreated,
	StateStr:      txmgrpb.TxState_StateCreated.String(),
	Type:          basetypes.TxType_TxWithdraw,
	TypeStr:       basetypes.TxType_TxWithdraw.String(),
	Extra:         uuid.NewString(),
}

var req = &txmgrpb.TxReq{
	FromAccountID: &ret.FromAccountID,
	ToAccountID:   &ret.ToAccountID,
	Amount:        &ret.Amount,
	FeeAmount:     &ret.FeeAmount,
	State:         &ret.State,
	Type:          &ret.Type,
	Extra:         &ret.Extra,
}

func create(t *testing.T) {
	coin, err := coinmw.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)
	assert.NotNil(t, coin)

	req.CoinTypeID = &coin.ID

	info, err := CreateTx(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.CoinTypeID = coin.ID

		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	state := txmgrpb.TxState_StateTransferring

	ret.State = state

	req.ID = &ret.ID
	req.State = &state

	_, err := UpdateTx(context.Background(), req)
	assert.NotNil(t, err)

	state = txmgrpb.TxState_StateWait

	ret.State = state
	ret.StateStr = state.String()
	req.State = &state

	info, err := UpdateTx(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("create", update)
}
