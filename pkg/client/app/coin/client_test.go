package appcoin

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
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
	"github.com/stretchr/testify/assert"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"

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
var logo = uuid.NewString()

var ret = &npool.Coin{
	EntID:                       uuid.NewString(),
	AppID:                       uuid.NewString(),
	CoinName:                    name,
	Name:                        name,
	DisplayNamesStr:             "[]",
	DisplayNames:                []string{"tttttttt", "tttt1"},
	Logo:                        logo,
	Unit:                        unit,
	Presale:                     false,
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
	WithdrawAutoReviewAmount:    "0.000000000000000000",
	MarketValue:                 "0.000000000000000000",
	SettleValue:                 "0.000000000000000000",
	SettlePercent:               80,
	SettleTipsStr:               "[]",
	SettleTips:                  []string{"tttttttttttt", "ttt"},
	Setter:                      uuid.NewString(),
	Display:                     true,
	DailyRewardAmount:           "0.000000000000000000",
	MaxAmountPerWithdraw:        "0.000000000000000000",
	LeastTransferAmount:         "0.000000000000000000",
}

var chainType = uuid.NewString()
var chainAtomicUnit = uuid.NewString()
var chainUnitExp = uint32(1)
var gasType = basetypes.GasType_FixedGas
var chainID = uuid.NewString()
var chainNickname = uuid.NewString()
var chainNativeCoinName = uuid.NewString()

var req = &npool.CoinReq{
	AppID:                    &ret.AppID,
	Name:                     &ret.Name,
	DisplayNames:             ret.DisplayNames,
	Logo:                     &ret.Logo,
	ForPay:                   &ret.ForPay,
	WithdrawAutoReviewAmount: &ret.WithdrawAutoReviewAmount,
	MarketValue:              &ret.MarketValue,
	SettlePercent:            &ret.SettlePercent,
	SettleTips:               ret.SettleTips,
	Setter:                   &ret.Setter,
	DailyRewardAmount:        &ret.DailyRewardAmount,
	MaxAmountPerWithdraw:     &ret.MaxAmountPerWithdraw,
}

func setupAppCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	ret.FeeCoinLogo = ret.Logo
	req.CoinTypeID = &ret.CoinTypeID

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithEntID(&ret.CoinTypeID, true),
		coin1.WithName(&ret.CoinName, true),
		coin1.WithUnit(&ret.Unit, true),
		coin1.WithLogo(&ret.Logo, true),
		coin1.WithENV(&ret.ENV, true),
		coin1.WithChainType(&chainType, true),
		coin1.WithChainNativeUnit(&ret.Unit, true),
		coin1.WithChainAtomicUnit(&chainAtomicUnit, true),
		coin1.WithChainUnitExp(&chainUnitExp, true),
		coin1.WithGasType(&gasType, true),
		coin1.WithChainID(&chainID, true),
		coin1.WithChainNickname(&chainNickname, true),
		coin1.WithChainNativeCoinName(&chainNativeCoinName, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func createCoin(t *testing.T) {
	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.FeeCoinTypeID = info.FeeCoinTypeID
		ret.DisplayNamesStr = info.DisplayNamesStr
		ret.SettleTipsStr = info.SettleTipsStr
		ret.CheckNewAddressBalance = info.CheckNewAddressBalance
		assert.Equal(t, ret, info)
	}
}

func updateCoin(t *testing.T) {
	amount := "123.700000000000000000"
	index := uint32(1)

	ret.WithdrawAutoReviewAmount = amount
	ret.MarketValue = amount
	ret.SettleValue = "98.960000000000000000"
	ret.DisplayIndex = index

	req.ID = &ret.ID
	req.WithdrawAutoReviewAmount = &amount
	req.MarketValue = &amount
	req.DisplayIndex = &index

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
		if assert.Equal(t, len(infos), 1) {
			assert.Equal(t, infos[0], ret)
		}
	}
}

func deleteCoin(t *testing.T) {
	info, err := DeleteCoin(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}

	info, err = GetCoin(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	teardown := setupAppCoin(t)
	defer teardown(t)

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
	t.Run("deleteCoin", deleteCoin)
}
