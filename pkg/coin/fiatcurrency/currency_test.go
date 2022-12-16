package fiatcurrency

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	"github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	fiatcurrencytypecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/fiatcurrencytype"
	fiatcurrencytypepb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrencytype"

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

var ret = &npool.FiatCurrency{
	ID:              uuid.NewString(),
	CoinName:        uuid.NewString(),
	CoinUnit:        uuid.NewString(),
	CoinENV:         uuid.NewString(),
	FeedType:        currency.FeedType_CoinBase,
	FeedTypeStr:     currency.FeedType_CoinBase.String(),
	MarketValueHigh: "12.001000000000000000",
	MarketValueLow:  "11.001000000000000000",
}

var fiat = &fiatcurrencytypepb.FiatCurrencyTypeReq{
	Name: &ret.CoinName,
}

var req = &fiatcurrencymgrpb.FiatCurrencyReq{
	ID:              &ret.ID,
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func create(t *testing.T) {
	coinRet, err := fiatcurrencytypecrud.Create(context.Background(), fiat)
	assert.Nil(t, err)

	id := coinRet.ID.String()
	req.FiatCurrencyTypeID = &id
	ret.CoinTypeID = coinRet.ID.String()

	info, err := CreateFiatCurrency(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("update", update)
}
