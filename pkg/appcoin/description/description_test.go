package description

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	coinmw "github.com/NpoolPlatform/chain-middleware/pkg/coin"
	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	coinName = "BTC111"
	coinUnit = "BTC"
	coinENV  = "test"
)

var coinReq = &coinmwpb.CoinReq{
	Name: &coinName,
	Unit: &coinUnit,
	ENV:  &coinENV,
}

var ret = &npool.CoinDescription{
	AppID:      uuid.NewString(),
	CoinName:   coinName,
	CoinUnit:   coinUnit,
	CoinENV:    coinENV,
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
	UsedFor:    descmgrpb.UsedFor_ProductPage,
	UsedForStr: descmgrpb.UsedFor_ProductPage.String(),
}

var req = &descmgrpb.CoinDescriptionReq{
	AppID:   &ret.AppID,
	Title:   &ret.Title,
	Message: &ret.Message,
	UsedFor: &ret.UsedFor,
}

func create(t *testing.T) {
	coin, err := coinmw.CreateCoin(context.Background(), coinReq)
	assert.Nil(t, err)

	req.CoinTypeID = &coin.ID

	info, err := CreateCoinDescription(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.CoinTypeID = coin.ID

		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
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

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("create", update)
}
