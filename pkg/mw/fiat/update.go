package fiat

import (
	"context"
	"fmt"

	fiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateFiat(ctx context.Context, cli *ent.Client) error {
	if _, err := fiatcrud.UpdateSet(
		cli.Fiat.UpdateOneID(*h.ID),
		&fiatcrud.Req{
			Name: h.Name,
			Logo: h.Logo,
			Unit: h.Unit,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.updateFiat(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
