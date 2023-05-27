package currency

import (
	"context"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat/currency"
	currencyhiscrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat/currency/history"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		FiatCurrency.
		Query().
		Where(
			entcurrency.FiatID(*h.FiatID),
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
		info, err = currencycrud.UpdateSet(
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
		tx.FiatCurrency.Create(),
		&currencycrud.Req{
			FiatID:          h.FiatID,
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
		tx.FiatCurrencyHistory.Create(),
		&currencyhiscrud.Req{
			FiatID:          h.FiatID,
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