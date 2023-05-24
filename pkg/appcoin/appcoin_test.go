package appcoin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	coinmw "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
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

var ret = &npool.Coin{
	AppID:                       uuid.NewString(),
	CoinName:                    coinName,
	Name:                        "My BTC1",
	DisplayNames:                []string{"123123", "2323"},
	Logo:                        uuid.NewString(),
	Unit:                        coinUnit,
	ENV:                         coinENV,
	ReservedAmount:              "0.000000000000000000",
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "0.000000000000000000",
	CollectFeeAmount:            "0.000000000000000000",
	HotWalletFeeAmount:          "0.000000000000000000",
	LowFeeAmount:                "0.000000000000000000",
	HotLowFeeAmount:             "0.000000000000000000",
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinName:                 coinName,
	FeeCoinUnit:                 coinUnit,
	FeeCoinENV:                  coinENV,
	WithdrawAutoReviewAmount:    "0.000000000000000000",
	MarketValue:                 "0.000000000000000000",
	SettleValue:                 "0.000000000000000000",
	SettlePercent:               90,
	SettleTips:                  []string{"1213", "234234"},
	Setter:                      uuid.NewString(),
	Display:                     true,
	DailyRewardAmount:           "0.000000000000000000",
	MaxAmountPerWithdraw:        "0.000000000000000000",
	LeastTransferAmount:         "0.000000000000000000",
}

var req = &npool.CoinReq{
	AppID:         &ret.AppID,
	Name:          &ret.Name,
	DisplayNames:  ret.DisplayNames,
	Logo:          &ret.Logo,
	SettlePercent: &ret.SettlePercent,
	SettleTips:    ret.SettleTips,
	Setter:        &ret.Setter,
}

func create(t *testing.T) {
	coin, err := coinmw.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)

	req.CoinTypeID = &coin.ID

	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.CoinTypeID = coin.ID
		ret.FeeCoinTypeID = coin.ID
		ret.DisplayNamesStr = info.DisplayNamesStr
		ret.SettleTipsStr = info.SettleTipsStr
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	amount := "123.700000000000000000"
	logo := uuid.NewString()

	ret.Logo = logo
	ret.WithdrawAutoReviewAmount = amount
	ret.MarketValue = amount
	ret.SettleValue = "111.330000000000000000"

	req.ID = &ret.ID
	req.Logo = &logo
	req.WithdrawAutoReviewAmount = &amount
	req.MarketValue = &amount

	info, err := UpdateCoin(context.Background(), req)
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
	t.Run("update", update)
}
