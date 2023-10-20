package currencyfeed

import (
	"context"
	"fmt"

	currencyfeedcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/feed"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateFeed(ctx context.Context) (*npool.Feed, error) {
	if h.FeedCoinName == nil {
		return nil, fmt.Errorf("invalid feedcoinname")
	}

	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoinCurrencyFeed,
		*h.CoinTypeID,
		*h.FeedType,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &currencyfeedcrud.Conds{
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		FeedType:   &cruder.Cond{Op: cruder.EQ, Val: *h.FeedType},
	}
	exist, err := h.ExistFeedConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coinfeed exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := currencyfeedcrud.CreateSet(
			cli.CurrencyFeed.Create(),
			&currencyfeedcrud.Req{
				EntID:        h.EntID,
				CoinTypeID:   h.CoinTypeID,
				FeedType:     h.FeedType,
				FeedCoinName: h.FeedCoinName,
				Disabled:     h.Disabled,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFeed(ctx)
}
