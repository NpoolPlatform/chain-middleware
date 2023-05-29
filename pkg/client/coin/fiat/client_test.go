package coinfiat

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
	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"
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

var ret = &npool.CoinFiat{
	CoinName:    uuid.NewString(),
	CoinUnit:    uuid.NewString(),
	CoinENV:     "test",
	FiatName:    uuid.NewString(),
	FiatLogo:    uuid.NewString(),
	FiatUnit:    "USD",
	FeedType:    basetypes.CurrencyFeedType_CoinBase,
	FeedTypeStr: basetypes.CurrencyFeedType_CoinBase.String(),
}

var req = &npool.CoinFiatReq{
	FeedType: &ret.FeedType,
}

func setupCoinFiat(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID

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

	h2, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithID(&ret.FiatID),
		fiat1.WithName(&ret.FiatName),
		fiat1.WithLogo(&ret.FiatLogo),
		fiat1.WithUnit(&ret.FiatUnit),
	)
	assert.Nil(t, err)

	_, err = h2.CreateFiat(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func createCoinFiat(t *testing.T) {

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

func deleteCoinFiat(t *testing.T) {

}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	teardown := setupCoinFiat(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCoinFiat", createCoinFiat)
	t.Run("getCoinFiats", getCoinFiats)
	t.Run("deleteCoinFiat", deleteCoinFiat)
}
