package fiat

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
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"
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

var ret = &npool.Fiat{
	Name: uuid.NewString(),
	Logo: uuid.NewString(),
	Unit: "USD",
}

var req = &npool.FiatReq{
	Name: &ret.Name,
	Logo: &ret.Logo,
	Unit: &ret.Unit,
}

func createFiat(t *testing.T) {
	info, err := CreateFiat(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func updateFiat(t *testing.T) {
	ret.Logo = uuid.NewString()
	req.ID = &ret.ID
	req.Logo = &ret.Logo

	info, err := UpdateFiat(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func getFiat(t *testing.T) {
	info, err := GetFiat(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getFiats(t *testing.T) {
	infos, total, err := GetFiats(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
	}, 0, 2)
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

	t.Run("createFiat", createFiat)
	t.Run("updateFiat", updateFiat)
	t.Run("getFiat", getFiat)
	t.Run("getFiats", getFiats)
}
