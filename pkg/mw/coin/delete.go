package coin

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	basecrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin"
	extracrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/setting"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entextra "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinextra"
	entsetting "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/setting"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteCoinBase(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := basecrud.UpdateSet(
		tx.CoinBase.UpdateOneID(*h.ID),
		&basecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteCoinExtra(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		CoinExtra.
		Query().
		Where(
			entextra.CoinTypeID(*h.ID),
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
	if _, err := extracrud.UpdateSet(
		info.Update(),
		&extracrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *deleteHandler) deleteCoinSetting(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Setting.
		Query().
		Where(
			entsetting.CoinTypeID(*h.ID),
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
	if _, err := settingcrud.UpdateSet(
		info.Update(),
		&settingcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) DeleteCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	info, err := h.GetCoin(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCoinBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteCoinExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteCoinSetting(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
