package currency

import (
	"context"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency"
	currencyhiscrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/history"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currency"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Currency.
		Query().
		Where(
			entcurrency.CoinTypeID(*h.CoinTypeID),
			entcurrency.FeedType(h.FeedType.String()),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		info, err := currencycrud.UpdateSet(
			info.Update(),
			&currencycrud.Req{
				MarketValueHigh: h.MarketValueHigh,
				MarketValueLow:  h.MarketValueLow,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID

		return nil
	}

	info, err = currencycrud.CreateSet(
		tx.Currency.Create(),
		&currencycrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	h.ID = &info.ID

	return nil
}

func (h *createHandler) createCurrencyHistory(ctx context.Context, tx *ent.Tx) error {
	if _, err := currencyhiscrud.CreateSet(
		tx.CurrencyHistory.Create(),
		&currencyhiscrud.Req{
			CoinTypeID:      h.CoinTypeID,
			FeedType:        h.FeedType,
			MarketValueHigh: h.MarketValueHigh,
			MarketValueLow:  h.MarketValueLow,
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
		if err := handler.createCurrency(ctx, tx); err != nil {
			return err
		}
		if err := handler.createCurrencyHistory(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCurrency(ctx)
}
