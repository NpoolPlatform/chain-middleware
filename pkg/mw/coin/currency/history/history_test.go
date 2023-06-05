package currencyhistory

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/history"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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

var ret = &currencymwpb.Currency{
	CoinName:        "My BTC1",
	CoinLogo:        uuid.NewString(),
	CoinUnit:        "BTC",
	CoinENV:         "test",
	MarketValueHigh: "0.000000000000000000",
	MarketValueLow:  "0.000000000000000000",
	FeedType:        basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinGecko.String(),
}

var req = &currencymwpb.CurrencyReq{
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithID(&ret.CoinTypeID),
		coin1.WithName(&ret.CoinName),
		coin1.WithLogo(&ret.CoinLogo),
		coin1.WithUnit(&ret.CoinUnit),
		coin1.WithENV(&ret.CoinENV),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	h2, err := currency1.NewHandler(
		context.Background(),
		currency1.WithCoinTypeID(req.CoinTypeID),
		currency1.WithMarketValueHigh(req.MarketValueHigh),
		currency1.WithMarketValueLow(req.MarketValueLow),
		currency1.WithFeedType(req.FeedType),
	)
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func getMany(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		}),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetCurrencies(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, uint32(2), total)
	assert.Equal(t, 2, len(infos))
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("getMany", getMany)
}
