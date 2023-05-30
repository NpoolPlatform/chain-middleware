package currency

import (
	"context"
	"fmt"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat/currency"
	currencyhiscrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat/currency/history"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinfiatcurrency"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCurrency(ctx context.Context, tx *ent.Tx) error {
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoinFiatCurrency,
		*h.CoinTypeID,
		*h.FiatID,
		*h.FeedType,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	info, err := tx.
		CoinFiatCurrency.
		Query().
		Where(
			entcurrency.CoinTypeID(*h.CoinTypeID),
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
		h.ID = &info.ID

		if _, err := currencycrud.UpdateSet(
			info.Update(),
			&currencycrud.Req{
				MarketValueHigh: h.MarketValueHigh,
				MarketValueLow:  h.MarketValueLow,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	info, err = currencycrud.CreateSet(
		tx.CoinFiatCurrency.Create(),
		&currencycrud.Req{
			CoinTypeID:      h.CoinTypeID,
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
		tx.CoinFiatCurrencyHistory.Create(),
		&currencyhiscrud.Req{
			CoinTypeID:      h.CoinTypeID,
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
