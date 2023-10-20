package feed

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
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"
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

var ret = &npool.Feed{
	FiatName:     uuid.NewString(),
	FiatUnit:     "BTC",
	FiatLogo:     uuid.NewString(),
	FeedType:     basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:  basetypes.CurrencyFeedType_CoinGecko.String(),
	FeedFiatName: "bitcoin",
}

var req = &npool.FeedReq{
	FeedType:     &ret.FeedType,
	FeedFiatName: &ret.FeedFiatName,
}

func setupFeed(t *testing.T) func(*testing.T) {
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

	return func(*testing.T) {}
}

func createFeed(t *testing.T) {
	info, err := CreateFeed(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func updateFeed(t *testing.T) {
	ret.Disabled = true
	req.ID = &ret.ID
	req.Disabled = &ret.Disabled

	info, err := UpdateFeed(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getFeeds(t *testing.T) {
	infos, total, err := GetFeeds(context.Background(), &npool.Conds{
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		FiatID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.FiatID},
	}, 0, 100)
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

	teardown := setupFeed(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createFeed", createFeed)
	t.Run("updateFeed", updateFeed)
	t.Run("getFeeds", getFeeds)
}
