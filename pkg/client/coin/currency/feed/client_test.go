package currencyfeed

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
	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"
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

var ret = &npool.CurrencyFeed{
	ID:          uuid.NewString(),
	CoinName:    uuid.NewString(),
	CoinUnit:    uuid.NewString(),
	CoinENV:     "test",
	FeedType:    feedmgrpb.FeedType_CoinBase,
	FeedTypeStr: feedmgrpb.FeedType_CoinBase.String(),
	FeedSource:  uuid.NewString(),
}

var coin = &coinmwpb.CoinReq{
	Name: &ret.CoinName,
	Unit: &ret.CoinUnit,
	ENV:  &ret.CoinENV,
}

var req = &feedmgrpb.CurrencyFeedReq{
	ID:         &ret.ID,
	FeedSource: &ret.FeedSource,
	FeedType:   &ret.FeedType,
}

func createCurrencyFeed(t *testing.T) {
	coinRet, err := coin1.CreateCoin(context.Background(), coin)
	assert.Nil(t, err)
	assert.NotNil(t, coinRet)

	ret.CoinTypeID = coinRet.ID
	req.CoinTypeID = &coinRet.ID

	info, err := CreateCurrencyFeed(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateCurrencyFeed(t *testing.T) {
	disabled := true

	ret.Disabled = disabled
	req.Disabled = &disabled

	info, err := UpdateCurrencyFeed(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCurrencyFeed(t *testing.T) {
}

func getCurrencyFeeds(t *testing.T) {
	infos, total, err := GetCurrencyFeeds(context.Background(), &npool.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
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

	t.Run("createCurrencyFeed", createCurrencyFeed)
	t.Run("updateCurrencyFeed", updateCurrencyFeed)
	t.Run("getCurrencyFeed", getCurrencyFeed)
	t.Run("getCurrencyFeeds", getCurrencyFeeds)
}
