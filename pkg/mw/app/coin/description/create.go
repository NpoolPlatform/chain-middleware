package description

import (
	"context"

	descriptioncrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := descriptioncrud.CreateSet(
			cli.CoinDescription.Create(),
			&descriptioncrud.Req{
				ID:         h.ID,
				AppID:      h.AppID,
				CoinTypeID: h.CoinTypeID,
				UsedFor:    h.UsedFor,
				Title:      h.Title,
				Message:    h.Message,
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
