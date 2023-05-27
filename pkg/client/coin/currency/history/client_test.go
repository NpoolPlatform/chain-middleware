package currencyhistory

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

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/history"
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

var ret = &currencymwpb.Currency{
	ID:              uuid.NewString(),
	CoinName:        uuid.NewString(),
	CoinUnit:        uuid.NewString(),
	CoinENV:         "test",
	FeedType:        basetypes.CurrencyFeedType_CoinBase,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinBase.String(),
	MarketValueHigh: "12.001000000000000000",
	MarketValueLow:  "11.001000000000000000",
}

var req = &currencymwpb.CurrencyReq{
	ID:              &ret.ID,
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupCurrencyHistory(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithID(&ret.CoinTypeID),
		coin1.WithName(&ret.CoinName),
		coin1.WithUnit(&ret.CoinUnit),
		coin1.WithLogo(&ret.CoinLogo),
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

func getCurrencies(t *testing.T) {
	infos, total, err := GetCurrencies(context.Background(), &npool.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
	}, 0, 100)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(infos))
		assert.Equal(t, uint32(2), total)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	teardown := setupCurrencyHistory(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("getCurrencies", getCurrencies)
}
