package appcoin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"

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
	coinUnit = "BTC"
	coinENV  = "test"
)

var ret = &npool.Coin{
	AppID:                       uuid.NewString(),
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

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()
	ret.FeeCoinName = ret.CoinName

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithID(&ret.CoinTypeID),
		coin1.WithName(&ret.CoinName),
		coin1.WithUnit(&coinUnit),
		coin1.WithENV(&coinENV),
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
		WithID(req.ID),
		WithAppID(req.AppID),
		WithCoinTypeID(req.CoinTypeID),
		WithName(req.Name),
		WithDisplayNames(req.DisplayNames),
		WithLogo(req.Logo),
		WithForPay(req.ForPay),
		WithProductPage(req.ProductPage),
		WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount),
		WithDailyRewardAmount(req.DailyRewardAmount),
		WithDisplay(req.Display),
		WithDisplayIndex(req.DisplayIndex),
		WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw),
		WithMarketValue(req.MarketValue),
		WithSettlePercent(req.SettlePercent),
		WithSettleTips(req.SettleTips),
		WithSetter(req.Setter),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.FeeCoinTypeID = ret.CoinTypeID
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

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID),
		WithName(req.Name),
		WithDisplayNames(req.DisplayNames),
		WithLogo(req.Logo),
		WithForPay(req.ForPay),
		WithProductPage(req.ProductPage),
		WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount),
		WithDailyRewardAmount(req.DailyRewardAmount),
		WithDisplay(req.Display),
		WithDisplayIndex(req.DisplayIndex),
		WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw),
		WithMarketValue(req.MarketValue),
		WithSettlePercent(req.SettlePercent),
		WithSettleTips(req.SettleTips),
		WithSetter(req.Setter),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func _delete(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	_, err = handler.DeleteCoin(context.Background())
	assert.Nil(t, err)
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", _delete)
}
