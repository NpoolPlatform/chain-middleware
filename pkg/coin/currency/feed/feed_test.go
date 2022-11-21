package currencyfeed

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/chain-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"
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

var ret = &npool.CurrencyFeed{
	ID:          uuid.NewString(),
	CoinName:    uuid.NewString(),
	CoinUnit:    uuid.NewString(),
	CoinENV:     uuid.NewString(),
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

func create(t *testing.T) {
	coinRet, err := coin1.CreateCoin(context.Background(), coin)
	assert.Nil(t, err)

	req.CoinTypeID = &coinRet.ID
	ret.CoinTypeID = coinRet.ID

	info, err := CreateCurrencyFeed(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {

}

func TestTx(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("update", update)
}
