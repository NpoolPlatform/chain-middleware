package currency

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"

	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"

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

var ret = &npool.Currency{
	FiatName:        "My BTC1",
	FiatLogo:        uuid.NewString(),
	FiatUnit:        "BTC",
	MarketValueHigh: "0.000000000000000000",
	MarketValueLow:  "0.000000000000000000",
	FeedType:        basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinGecko.String(),
}

var req = &npool.CurrencyReq{
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupFiat(t *testing.T) func(*testing.T) {
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID
	ret.FiatName = uuid.NewString()

	h1, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithID(&ret.FiatID),
		fiat1.WithName(&ret.FiatName),
		fiat1.WithLogo(&ret.FiatLogo),
		fiat1.WithUnit(&ret.FiatUnit),
	)
	assert.Nil(t, err)

	_, err = h1.CreateFiat(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithFiatID(req.FiatID),
		WithMarketValueHigh(req.MarketValueHigh),
		WithMarketValueLow(req.MarketValueLow),
		WithFeedType(req.FeedType),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCurrency(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, ret)
	}
}

func TestFiat(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupFiat(t)
	defer teardown(t)

	t.Run("create", create)
}
