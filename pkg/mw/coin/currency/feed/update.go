package currencyfeed

import (
	"context"
	"fmt"

	currencyfeedcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) UpdateFeed(ctx context.Context) (*npool.Feed, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := currencyfeedcrud.UpdateSet(
			cli.CurrencyFeed.UpdateOneID(*h.ID),
			&currencyfeedcrud.Req{
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
