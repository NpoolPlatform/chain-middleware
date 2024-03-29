package currency

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
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"
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

var ret = &npool.Currency{
	EntID:           uuid.NewString(),
	FiatName:        uuid.NewString(),
	FiatUnit:        uuid.NewString(),
	FeedType:        basetypes.CurrencyFeedType_CoinBase,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinBase.String(),
	MarketValueHigh: "12.001000000000000000",
	MarketValueLow:  "11.001000000000000000",
}

var req = &npool.CurrencyReq{
	ID:              &ret.ID,
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupCurrency(t *testing.T) func(*testing.T) {
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID

	h1, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithEntID(&ret.FiatID, true),
		fiat1.WithName(&ret.FiatName, true),
		fiat1.WithUnit(&ret.FiatUnit, true),
		fiat1.WithLogo(&ret.FiatLogo, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateFiat(context.Background())
	assert.Nil(t, err)

	h2, err := currency1.NewHandler(
		context.Background(),
		currency1.WithFiatID(req.FiatID, true),
		currency1.WithMarketValueHigh(req.MarketValueHigh, true),
		currency1.WithMarketValueLow(req.MarketValueLow, true),
		currency1.WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	info, err := h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	ret.ID = info.ID
	ret.EntID = info.EntID
	ret.CreatedAt = info.CreatedAt
	ret.UpdatedAt = info.UpdatedAt

	return func(*testing.T) {}
}

func getCurrency(t *testing.T) {
	info, err := GetCurrency(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getCurrencies(t *testing.T) {
	infos, total, err := GetCurrencies(context.Background(), &npool.Conds{
		FiatID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.FiatID},
	}, 0, 100)
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, ret, infos[0])
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	teardown := setupCurrency(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("getCurrency", getCurrency)
	t.Run("getCurrencies", getCurrencies)
}
