package description

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"

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

var ret = &npool.CoinDescription{
	ID:         uuid.NewString(),
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

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithID(&ret.CoinTypeID),
		coin1.WithName(&ret.CoinName),
		coin1.WithLogo(&ret.CoinLogo),
		coin1.WithUnit(&ret.CoinUnit),
		coin1.WithENV(&ret.CoinENV),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	h2, err := appcoin1.NewHandler(
		context.Background(),
		appcoin1.WithAppID(&ret.AppID),
		appcoin1.WithCoinTypeID(&ret.CoinTypeID),
	)
	assert.Nil(t, err)

	_, err = h2.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
		_, _ = h2.DeleteCoin(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID),
		WithAppID(req.AppID),
		WithCoinTypeID(req.CoinTypeID),
		WithTitle(req.Title),
		WithMessage(req.Message),
		WithUsedFor(req.UsedFor),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoinDescription(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.Title = uuid.NewString()
	ret.Message = uuid.NewString()

	req.Title = &ret.Title
	req.Message = &ret.Message

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithTitle(req.Title),
		WithMessage(req.Message),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoinDescription(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
}
