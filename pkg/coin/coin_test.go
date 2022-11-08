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
	Name:                        "bitcoin1",
	Unit:                        "BTC1",
	ENV:                         "test",
	ReservedAmount:              "0.000000000000000000",
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "0.000000000000000000",
	CollectFeeAmount:            "0.000000000000000000",
	HotWalletFeeAmount:          "0.000000000000000000",
	LowFeeAmount:                "0.000000000000000000",
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinName:                 "bitcoin1",
	FeeCoinUnit:                 "BTC1",
	FeeCoinENV:                  "test",
}

var req = &npool.CoinReq{
	Name: &ret.Name,
	Unit: &ret.Unit,
	ENV:  &ret.ENV,
}

func create(t *testing.T) {
	info, err := CreateCoin(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.FeeCoinTypeID = info.ID
		assert.Equal(t, info, ret)
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
	ret.HotWalletAccountAmount = amount
	ret.PaymentAccountCollectAmount = amount
	ret.FeeCoinLogo = logo

	req.ID = &ret.ID
	req.Logo = &logo
	req.WithdrawFeeByStableUSD = &feeByUSD
	req.ReservedAmount = &amount
	req.WithdrawFeeAmount = &amount
	req.CollectFeeAmount = &amount
	req.HotWalletFeeAmount = &amount
	req.LowFeeAmount = &amount
	req.HotWalletAccountAmount = &amount
	req.PaymentAccountCollectAmount = &amount

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
	t.Run("create", update)
}
