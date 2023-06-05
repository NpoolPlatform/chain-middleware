package coinfiat

import (
	"context"
	"time"

	coinfiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) DeleteCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	info, err := h.GetCoinFiat(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := coinfiatcrud.UpdateSet(
			cli.CoinFiat.UpdateOneID(*h.ID),
			&coinfiatcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
