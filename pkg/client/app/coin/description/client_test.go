package description

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"
	"github.com/stretchr/testify/assert"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"

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

var ret = &npool.CoinDescription{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinLogo:   uuid.NewString(),
	CoinUnit:   "BTC",
	CoinENV:    "test",
	UsedForStr: basetypes.UsedFor_ProductPage.String(),
	UsedFor:    basetypes.UsedFor_ProductPage,
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
}

var req = &npool.CoinDescriptionReq{
	AppID:   &ret.AppID,
	UsedFor: &ret.UsedFor,
	Title:   &ret.Title,
	Message: &ret.Message,
}

func setupAppCoinDescription(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithEntID(&ret.CoinTypeID, true),
		coin1.WithName(&ret.CoinName, true),
		coin1.WithLogo(&ret.CoinLogo, true),
		coin1.WithUnit(&ret.CoinUnit, true),
		coin1.WithENV(&ret.CoinENV, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	h2, err := appcoin1.NewHandler(
		context.Background(),
		appcoin1.WithAppID(&ret.AppID, true),
		appcoin1.WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	_, err = h2.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
		_, _ = h2.DeleteCoin(context.Background())
	}
}

func createCoinDescription(t *testing.T) {
	info, err := CreateCoinDescription(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func updateCoinDescription(t *testing.T) {
	ret.Title = uuid.NewString()
	ret.Message = uuid.NewString()

	req.ID = &ret.ID
	req.Title = &ret.Title
	req.Message = &ret.Message

	info, err := UpdateCoinDescription(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCoinDescription(t *testing.T) {
	info, err := GetCoinDescription(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getCoinDescriptions(t *testing.T) {
	infos, total, err := GetCoinDescriptions(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}, 0, 100)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, total, uint32(1))
		if assert.Equal(t, len(infos), 1) {
			assert.Equal(t, infos[0], ret)
		}
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost

	teardown := setupAppCoinDescription(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createCoinDescription", createCoinDescription)
	t.Run("updateCoinDescription", updateCoinDescription)
	t.Run("getCoinDescription", getCoinDescription)
	t.Run("getCoinDescriptions", getCoinDescriptions)
}
