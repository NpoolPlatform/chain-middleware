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

	commonpb "github.com/NpoolPlatform/message/npool"
	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	"github.com/stretchr/testify/assert"

	coincrud "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"

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

var name = uuid.NewString()
var unit = uuid.NewString()
var env = "main"

var coinReq = &coinmwpb.CoinReq{
	Name: &name,
	Unit: &unit,
	ENV:  &env,
}

var ret = &npool.CoinDescription{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinName:   name,
	CoinUnit:   unit,
	CoinENV:    env,
	UsedFor:    descmgrpb.UsedFor_ProductPage,
	UsedForStr: descmgrpb.UsedFor_ProductPage.String(),
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
}

var req = &descmgrpb.CoinDescriptionReq{
	ID:      &ret.ID,
	AppID:   &ret.AppID,
	UsedFor: &ret.UsedFor,
	Title:   &ret.Title,
	Message: &ret.Message,
}

func createCoinDescription(t *testing.T) {
	coin1, err := coincrud.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)

	req.CoinTypeID = &coin1.ID
	ret.CoinTypeID = coin1.ID

	info, err := CreateCoinDescription(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func updateCoinDescription(t *testing.T) {
	title := uuid.NewString()
	message := uuid.NewString()

	ret.Title = title
	ret.Message = message

	req.ID = &ret.ID
	req.Title = &title
	req.Message = &message

	info, err := UpdateCoinDescription(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getCoinDescription(t *testing.T) {
	info, err := GetCoinDescription(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getCoinDescriptions(t *testing.T) {
	infos, total, err := GetCoinDescriptions(context.Background(), &descmgrpb.Conds{
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

	t.Run("createCoinDescription", createCoinDescription)
	t.Run("updateCoinDescription", updateCoinDescription)
	t.Run("getCoinDescription", getCoinDescription)
	t.Run("getCoinDescriptions", getCoinDescriptions)
}
