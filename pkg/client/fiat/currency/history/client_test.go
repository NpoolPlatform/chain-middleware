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

	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"
	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/history"
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
	FiatName:        uuid.NewString(),
	FiatUnit:        uuid.NewString(),
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
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID

	h1, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithID(&ret.FiatID),
		fiat1.WithName(&ret.FiatName),
		fiat1.WithUnit(&ret.FiatUnit),
		fiat1.WithLogo(&ret.FiatLogo),
	)
	assert.Nil(t, err)

	_, err = h1.CreateFiat(context.Background())
	assert.Nil(t, err)

	h2, err := currency1.NewHandler(
		context.Background(),
		currency1.WithFiatID(req.FiatID),
		currency1.WithMarketValueHigh(req.MarketValueHigh),
		currency1.WithMarketValueLow(req.MarketValueLow),
		currency1.WithFeedType(req.FeedType),
	)
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {}
}

func getCurrencies(t *testing.T) {
	infos, total, err := GetCurrencies(context.Background(), &npool.Conds{
		FiatID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.FiatID},
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
