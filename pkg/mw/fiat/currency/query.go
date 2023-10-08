package currency

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	fiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat"
	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat/currency"
	entfiat "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiat"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatSelect
	infos []*npool.Currency
	total uint32
}

func (h *queryHandler) selectFiat(stm *ent.FiatQuery) {
	h.stm = stm.
		Select(entfiat.FieldCreatedAt).
		Modify(func(s *sql.Selector) {
			t := sql.Table(entfiat.Table)
			s.AppendSelect(
				sql.As(t.C(entfiat.FieldID), "fiat_id"),
				sql.As(t.C(entfiat.FieldName), "fiat_name"),
				sql.As(t.C(entfiat.FieldLogo), "fiat_logo"),
				sql.As(t.C(entfiat.FieldUnit), "fiat_unit"),
			)
		})
}

func (h *queryHandler) queryFiat(ctx context.Context, cli *ent.Client) error {
	_stm1, err := currencycrud.SetQueryConds(cli.FiatCurrency.Query(), &currencycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
	})
	if err != nil {
		return err
	}
	_info1, err := _stm1.Only(ctx)
	if err != nil {
		return err
	}

	_stm2, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), &fiatcrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: _info1.FiatID},
	})
	if err != nil {
		return err
	}

	h.selectFiat(_stm2)
	return nil
}

func (h *queryHandler) queryFiats(ctx context.Context, cli *ent.Client) error {
	stm, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), &fiatcrud.Conds{
		EntID:  h.Conds.FiatID,
		EntIDs: h.Conds.FiatIDs,
		Name:   h.Conds.FiatName,
	})
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectFiat(stm)
	return nil
}

func (h *queryHandler) queryJoinCurrency(s *sql.Selector) {
	s.GroupBy(entcurrency.FieldFiatID)
	t := sql.Table(entcurrency.Table)
	s.LeftJoin(t).
		On(
			s.C(entfiat.FieldID),
			t.C(entcurrency.FieldFiatID),
		).
		OnP(
			sql.EQ(t.C(entcurrency.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcurrency.FieldID), "id"),
			sql.As(t.C(entcurrency.FieldFeedType), "feed_type"),
			sql.As(t.C(entcurrency.FieldMarketValueHigh), "market_value_high"),
			sql.As(t.C(entcurrency.FieldMarketValueLow), "market_value_low"),
			sql.As(t.C(entcurrency.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcurrency.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCurrency(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.FeedType = basetypes.CurrencyFeedType(basetypes.CurrencyFeedType_value[info.FeedTypeStr])
		if _, err := decimal.NewFromString(info.MarketValueHigh); err != nil {
			info.MarketValueHigh = decimal.NewFromInt(0).String()
		}
		if _, err := decimal.NewFromString(info.MarketValueLow); err != nil {
			info.MarketValueLow = decimal.NewFromInt(0).String()
		}
	}
}

func (h *Handler) GetCurrency(ctx context.Context) (*npool.Currency, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiat(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 1
		handler.stm.
			Offset(0).
			Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCurrencies(ctx context.Context) ([]*npool.Currency, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiats(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
