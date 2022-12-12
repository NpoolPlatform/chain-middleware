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

	commonpb "github.com/NpoolPlatform/message/npool"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	"github.com/stretchr/testify/assert"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"

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
	ID:              uuid.NewString(),
	CoinName:        uuid.NewString(),
	CoinUnit:        uuid.NewString(),
	CoinENV:         "test",
	FeedType:        currencymgrpb.FeedType_CoinBase,
	FeedTypeStr:     currencymgrpb.FeedType_CoinBase.String(),
	MarketValueHigh: "12.001000000000000000",
	MarketValueLow:  "11.001000000000000000",
}

var coin = &coinmwpb.CoinReq{
	Name: &ret.CoinName,
	Unit: &ret.CoinUnit,
	ENV:  &ret.CoinENV,
}

var req = &currencymgrpb.CurrencyReq{
	ID:              &ret.ID,
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func createCurrency(t *testing.T) {
	coinRet, err := coin1.CreateCoin(context.Background(), coin)
	assert.Nil(t, err)
	assert.NotNil(t, coinRet)

	ret.CoinTypeID = coinRet.ID
	req.CoinTypeID = &coinRet.ID

	info, err := CreateCurrency(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateCurrency(t *testing.T) {
}

func getCurrency(t *testing.T) {
}

func getCurrencies(t *testing.T) {
	infos, err := GetCurrencies(context.Background(), &npool.Conds{
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.CoinTypeID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCurrency", createCurrency)
	t.Run("updateCurrency", updateCurrency)
	t.Run("getCurrency", getCurrency)
	t.Run("getCurrencies", getCurrencies)
}
