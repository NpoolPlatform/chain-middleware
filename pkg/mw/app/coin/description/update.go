package description

import (
	"context"
	"fmt"

	descriptioncrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) UpdateCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := descriptioncrud.UpdateSet(
			cli.CoinDescription.UpdateOneID(*h.ID),
			&descriptioncrud.Req{
				Title:   h.Title,
				Message: h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCoinDescription(ctx)
}
