package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

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

var ret = &npool.Coin{
	Name:                        uuid.NewString(),
	Unit:                        "BTC1",
	ENV:                         "test",
	ReservedAmount:              "0.000000000000000000",
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "0.000000000000000000",
	CollectFeeAmount:            "0.000000000000000000",
	HotWalletFeeAmount:          "0.000000000000000000",
	LowFeeAmount:                "0.000000000000000000",
	HotLowFeeAmount:             "0.000000000000000000",
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinUnit:                 "BTC1",
	FeeCoinENV:                  "test",
	LeastTransferAmount:         "0.000000000000000000",
	CheckNewAddressBalance:      true,
}

var req = &npool.CoinReq{
	Name: &ret.Name,
	Unit: &ret.Unit,
	ENV:  &ret.ENV,
}

//nolint
func setupCoin(t *testing.T) func(*testing.T) {
	ret.FeeCoinName = ret.Name
	return func(*testing.T) {}
}

func create(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithName(&ret.Name, true),
		WithUnit(&ret.Unit, true),
		WithENV(&ret.ENV, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.FeeCoinTypeID = info.EntID
		assert.Equal(t, info.String(), ret.String())
	}
}

func update(t *testing.T) {
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

	h1, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithName(req.Name, true),
		WithUnit(req.Unit, true),
		WithLogo(req.Logo, true),
		WithReservedAmount(req.ReservedAmount, true),
		WithHomePage(req.HomePage, true),
		WithSpecs(req.Specs, true),
		// TODO: this should be get from chain type
		WithFeeCoinTypeID(req.FeeCoinTypeID, true),
		WithWithdrawFeeByStableUSD(req.WithdrawFeeByStableUSD, true),
		WithWithdrawFeeAmount(req.WithdrawFeeAmount, true),
		WithCollectFeeAmount(req.CollectFeeAmount, true),
		WithHotWalletFeeAmount(req.HotWalletFeeAmount, true),
		WithLowFeeAmount(req.LowFeeAmount, true),
		WithHotLowFeeAmount(req.HotLowFeeAmount, true),
		WithHotWalletFeeAmount(req.HotWalletFeeAmount, true),
		WithHotWalletAccountAmount(req.HotWalletAccountAmount, true),
		WithPaymentAccountCollectAmount(req.PaymentAccountCollectAmount, true),
		WithLeastTransferAmount(req.LeastTransferAmount, true),
		WithPresale(req.Presale, true),
		WithForPay(req.ForPay, true),
		WithDisabled(req.Disabled, true),
		// TODO: this should be in create from register coin
		WithStableUSD(req.StableUSD, true),
		// TODO: this should be in create from register coin
		WithNeedMemo(req.NeedMemo, true),
		WithRefreshCurrency(req.RefreshCurrency, true),
		WithCheckNewAddressBalance(req.CheckNewAddressBalance, true),
	)
	assert.Nil(t, err)

	info, err := h1.UpdateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func _delete(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	_, err = h1.DeleteCoin(context.Background())
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
