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

	commonpb "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
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

var coinReq = &coinmwpb.CoinReq{
	Name: &ret.Name,
	Unit: &ret.Unit,
	ENV:  &ret.ENV,
}

var name = uuid.NewString()
var unit = uuid.NewString()
var logo = uuid.NewString()

var ret = &npool.Coin{
	ID:                          uuid.NewString(),
	AppID:                       uuid.NewString(),
	Name:                        name,
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
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinName:                 name,
	FeeCoinUnit:                 unit,
	FeeCoinENV:                  "main",
	WithdrawAutoReviewAmount:    "0.000000000000000000",
	MarketValue:                 "0.000000000000000000",
	SettleValue:                 "0.000000000000000000",
	SettlePercent:               80,
	Setter:                      uuid.NewString(),
}

var req = &npool.CoinReq{
	ID:                       &ret.ID,
	AppID:                    &ret.AppID,
	Name:                     &ret.Name,
	Logo:                     &ret.Logo,
	ForPay:                   &ret.ForPay,
	WithdrawAutoReviewAmount: &ret.WithdrawAutoReviewAmount,
	MarketValue:              &ret.MarketValue,
	SettlePercent:            &ret.SettlePercent,
	Setter:                   &ret.Setter,
}

func createCoin(t *testing.T) {
	coin1, err := coincrud.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)

	req.CoinTypeID = &coin1.ID
	ret.CoinTypeID = coin1.ID

	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.FeeCoinTypeID = info.FeeCoinTypeID
		assert.Equal(t, ret, info)
	}
}

func updateCoin(t *testing.T) {
	amount := "123.700000000000000000"

	ret.WithdrawAutoReviewAmount = amount
	ret.MarketValue = amount
	ret.SettleValue = "98.960000000000000000"

	req.ID = &ret.ID
	req.WithdrawAutoReviewAmount = &amount
	req.MarketValue = &amount

	info, err := UpdateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCoin(t *testing.T) {
	info, err := GetCoin(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getCoins(t *testing.T) {
	infos, total, err := GetCoins(context.Background(), &npool.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
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

	_, err = GetCoin(context.Background(), ret.ID)
	assert.NotNil(t, err)
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

	t.Run("createCoin", createCoin)
	t.Run("updateCoin", updateCoin)
	t.Run("getCoin", getCoin)
	t.Run("getCoins", getCoins)
	t.Run("deleteCoin", deleteCoin)
}
