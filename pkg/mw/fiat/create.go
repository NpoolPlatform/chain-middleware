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

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createFiat(ctx context.Context, cli *ent.Client) error {
	if _, err := fiatcrud.CreateSet(
		cli.Fiat.Create(),
		&fiatcrud.Req{
			ID:   h.ID,
			Name: h.Name,
			Logo: h.Logo,
			Unit: h.Unit,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFiat(ctx context.Context) (*npool.Fiat, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid fiatname")
	}
	if h.Unit == nil {
		return nil, fmt.Errorf("invalid fiatunit")
	}

	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateFiat,
		*h.Name,
		*h.Unit,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &fiatcrud.Conds{
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
		Unit: &cruder.Cond{Op: cruder.EQ, Val: *h.Unit},
	}
	h.Offset = 0
	h.Limit = 2

	fiat, err := h.GetFiatOnly(ctx)
	if err != nil {
		return nil, err
	}
	if fiat != nil {
		return fiat, nil
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createFiat(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFiat(ctx)
}
