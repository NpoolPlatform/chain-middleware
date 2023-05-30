package currency

import (
	"context"
	"fmt"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency"
	currencyhiscrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/history"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currency"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx, req *currencycrud.Req) error {
	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoinCurrency,
		*h.CoinTypeID,
		*h.FeedType,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	info, err := tx.
		Currency.
		Query().
		Where(
			entcurrency.CoinTypeID(*req.CoinTypeID),
			entcurrency.FeedType(req.FeedType.String()),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		info, err := currencycrud.
			UpdateSet(info.Update(), req).
			Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID

		return nil
	}

	info, err = currencycrud.
		CreateSet(tx.Currency.Create(), req).
		Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *createHandler) createCurrencyHistory(ctx context.Context, tx *ent.Tx, req *currencycrud.Req) error {
	if _, err := currencyhiscrud.CreateSet(
		tx.CurrencyHistory.Create(),
		&currencyhiscrud.Req{
			CoinTypeID:      req.CoinTypeID,
			FeedType:        req.FeedType,
			MarketValueHigh: req.MarketValueHigh,
			MarketValueLow:  req.MarketValueLow,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCurrency(ctx context.Context) (*npool.Currency, error) {
	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		req := &currencycrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
		}

		if err := handler.createCurrency(ctx, tx, req); err != nil {
			return err
		}
		if err := handler.createCurrencyHistory(ctx, tx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCurrency(ctx)
}

func (h *Handler) CreateCurrencies(ctx context.Context) ([]*npool.Currency, error) {
	handler := &createHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range handler.Reqs {
			if err := handler.createCurrency(ctx, tx, req); err != nil {
				return err
			}
			if err := handler.createCurrencyHistory(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *req.CoinTypeID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &currencycrud.Conds{
		CoinTypeIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetCurrencies(ctx)
	return infos, err
}
