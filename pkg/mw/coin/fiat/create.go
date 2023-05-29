package coinfiat

import (
	"context"

	coinfiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := coinfiatcrud.CreateSet(
			cli.CoinFiat.Create(),
			&coinfiatcrud.Req{
				ID:         h.ID,
				CoinTypeID: h.CoinTypeID,
				FiatID:     h.FiatID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoinFiat(ctx)
}
