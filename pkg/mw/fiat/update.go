package fiat

import (
	"context"
	"fmt"

	fiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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

	lockKey := fmt.Sprintf(
		"%v:%v",
		basetypes.Prefix_PrefixSetFiat,
		*h.Name,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &fiatcrud.Conds{
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
	}
	h.Offset = 0
	h.Limit = 2

	fiat, err := h.GetFiatOnly(ctx)
	if err != nil {
		return nil, err
	}
	if fiat != nil {
		return nil, fmt.Errorf("fiat exist")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
