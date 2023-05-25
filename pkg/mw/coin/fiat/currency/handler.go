package currency

import (
	"context"
	"fmt"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat/currency"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID              *uint32
	CoinTypeID      *uuid.UUID
	FiatID          *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
	Conds           *currencycrud.Conds
	Offset          int32
	Limit           int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		Conds: &currencycrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ID = id
		return nil
	}
}

func WithCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithFiatID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FiatID = &_id
		return nil
	}
}

func WithFeedType(feedType *basetypes.CurrencyFeedType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feedType == nil {
			return nil
		}
		switch *feedType {
		case basetypes.CurrencyFeedType_CoinGecko:
		case basetypes.CurrencyFeedType_CoinBase:
		case basetypes.CurrencyFeedType_StableUSDHardCode:
		default:
			return fmt.Errorf("invalid feedtype")
		}
		h.FeedType = feedType
		return nil
	}
}

func WithMarketValueHigh(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		_value, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.MarketValueHigh = &_value
		return nil
	}
}

func WithMarketValueLow(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		_value, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.MarketValueLow = &_value
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Limit = limit
		return nil
	}
}
