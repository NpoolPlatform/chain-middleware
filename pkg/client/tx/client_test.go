package tx

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/chain-middleware/pkg/testinit"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	commonpb "github.com/NpoolPlatform/message/npool"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"
	"github.com/stretchr/testify/assert"

	coincrud "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var name = uuid.NewString()
var unit = uuid.NewString()
var env = "main"

var coinReq = &coinmwpb.CoinReq{
	Name: &name,
	Unit: &unit,
	ENV:  &env,
}

var ret = &npool.Tx{
	ID:            uuid.NewString(),
	CoinName:      name,
	CoinUnit:      unit,
	CoinENV:       env,
	FromAccountID: uuid.NewString(),
	ToAccountID:   uuid.NewString(),
	Amount:        "123.345000000000000000",
	FeeAmount:     "2.001000000000000000",
	State:         txmgrpb.TxState_StateCreated,
	StateStr:      txmgrpb.TxState_StateCreated.String(),
	Extra:         uuid.NewString(),
	Type:          txmgrpb.TxType_TxWithdraw,
	TypeStr:       txmgrpb.TxType_TxWithdraw.String(),
}

var req = &txmgrpb.TxReq{
	ID:            &ret.ID,
	FromAccountID: &ret.FromAccountID,
	ToAccountID:   &ret.ToAccountID,
	Amount:        &ret.Amount,
	FeeAmount:     &ret.FeeAmount,
	State:         &ret.State,
	Extra:         &ret.Extra,
	Type:          &ret.Type,
}

func createTx(t *testing.T) {
	coin1, err := coincrud.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)

	req.CoinTypeID = &coin1.ID
	ret.CoinTypeID = coin1.ID

	info, err := CreateTx(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func updateTx(t *testing.T) {
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

func getTx(t *testing.T) {
	info, err := GetTx(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getTxs(t *testing.T) {
	infos, total, err := GetTxs(context.Background(), &txmgrpb.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], ret)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createTx", createTx)
	t.Run("updateTx", updateTx)
	t.Run("getTx", getTx)
	t.Run("getTxs", getTxs)
}
