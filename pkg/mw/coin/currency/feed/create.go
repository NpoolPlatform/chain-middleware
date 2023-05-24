package currencyfeed

import (
	"context"

	currencyfeedcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateFeed(ctx context.Context) (*npool.Feed, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := currencyfeedcrud.CreateSet(
			cli.CurrencyFeed.Create(),
			&currencyfeedcrud.Req{
				ID:           h.ID,
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
