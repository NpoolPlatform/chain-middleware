package appcoin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entappexrate "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/exchangerate"

	appcoincrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin"
	appexratecrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/exrate"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppCoin(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := appcoincrud.UpdateSet(
		tx.AppCoin.UpdateOneID(*h.ID),
		&appcoincrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteExrate(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		ExchangeRate.
		Query().
		Where(
			entappexrate.AppID(*h.AppID),
			entappexrate.CoinTypeID(*h.CoinTypeID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	if info == nil {
		return nil
	}

	now := uint32(time.Now().Unix())
	if _, err := appexratecrud.UpdateSet(
		info.Update(),
		&appexratecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteCoin(ctx context.Context) (*npool.Coin, error) {
	info, err := h.GetCoin(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppCoin(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteExrate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
