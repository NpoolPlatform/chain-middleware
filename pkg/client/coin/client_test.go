package coin

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/stretchr/testify/assert"

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
var chainType = uuid.NewString()
var chainAtomicUnit = uuid.NewString()
var chainUnitExp = uint32(1)
var gasType = basetypes.GasType_FixedGas
var chainID = uuid.NewString()
var chainNickname = uuid.NewString()
var chainNativeCoinName = uuid.NewString()
var ret = &npool.Coin{
	Name:                        name,
	Presale:                     false,
	Unit:                        unit,
	ENV:                         "main",
	ForPay:                      false,
	ReservedAmount:              "0.000000000000000000",
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "0.000000000000000000",
	CollectFeeAmount:            "0.000000000000000000",
	HotWalletFeeAmount:          "0.000000000000000000",
	LowFeeAmount:                "0.000000000000000000",
	HotLowFeeAmount:             "0.000000000000000000",
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinName:                 name,
	FeeCoinUnit:                 unit,
	FeeCoinENV:                  "main",
	LeastTransferAmount:         "0.000000000000000000",
	CheckNewAddressBalance:      true,
}

var req = &npool.CoinReq{
	Name:                &ret.Name,
	Unit:                &ret.Unit,
	ENV:                 &ret.ENV,
	ChainType:           &chainType,
	ChainAtomicUnit:     &chainAtomicUnit,
	ChainUnitExp:        &chainUnitExp,
	GasType:             &gasType,
	ChainID:             &chainID,
	ChainNickname:       &chainNickname,
	ChainNativeCoinName: &chainNativeCoinName,
	ChainNativeUnit:     &ret.Unit,
}

func createCoin(t *testing.T) {
	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.FeeCoinTypeID = info.FeeCoinTypeID
		assert.Equal(t, ret, info)
	}
}

func updateCoin(t *testing.T) {
	feeByUSD := false
	amount := "123.700000000000000000"
	logo := uuid.NewString()

	ret.Logo = logo
	ret.WithdrawFeeByStableUSD = feeByUSD
	ret.ReservedAmount = amount
	ret.WithdrawFeeAmount = amount
	ret.CollectFeeAmount = amount
	ret.HotWalletFeeAmount = amount
	ret.LowFeeAmount = amount
	ret.HotLowFeeAmount = amount
	ret.HotWalletAccountAmount = amount
	ret.PaymentAccountCollectAmount = amount
	ret.LeastTransferAmount = amount
	ret.FeeCoinLogo = logo

	req.ID = &ret.ID
	req.Logo = &logo
	req.WithdrawFeeByStableUSD = &feeByUSD
	req.ReservedAmount = &amount
	req.WithdrawFeeAmount = &amount
	req.CollectFeeAmount = &amount
	req.HotWalletFeeAmount = &amount
	req.LowFeeAmount = &amount
	req.HotLowFeeAmount = &amount
	req.HotWalletAccountAmount = &amount
	req.PaymentAccountCollectAmount = &amount
	req.LeastTransferAmount = &amount

	info, err := UpdateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCoin(t *testing.T) {
	info, err := GetCoin(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getCoins(t *testing.T) {
	infos, total, err := GetCoins(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 100)
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCoin", createCoin)
	t.Run("updateCoin", updateCoin)
	t.Run("getCoin", getCoin)
	t.Run("getCoins", getCoins)
}
