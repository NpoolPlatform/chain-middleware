package coinfiat

import (
	"context"

	coinfiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) CreateCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := coinfiatcrud.CreateSet(
			cli.CoinFiat.Create(),
			&coinfiatcrud.Req{
				CoinTypeID: h.CoinTypeID,
				FiatID:     h.FiatID,
				FeedType:   h.FeedType,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoinFiat(ctx)
}
