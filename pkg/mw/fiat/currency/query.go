package currency

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	currencycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat/currency"
	entfiat "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiat"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatCurrencySelect
	infos []*npool.Currency
	total uint32
}

func (h *queryHandler) selectCurrency(stm *ent.FiatCurrencyQuery) {
	h.stm = stm.Select(
		entcurrency.FieldID,
		entcurrency.FieldFiatID,
		entcurrency.FieldFeedType,
		entcurrency.FieldMarketValueHigh,
		entcurrency.FieldMarketValueLow,
		entcurrency.FieldCreatedAt,
		entcurrency.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCurrency(cli *ent.Client) error {
	h.selectCurrency(
		cli.FiatCurrency.
			Query().
			Where(
				entcurrency.ID(*h.ID),
			),
	)
	return nil
}

func (h *queryHandler) queryCurrencies(ctx context.Context, cli *ent.Client) error {
	stm, err := currencycrud.SetQueryConds(cli.FiatCurrency.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectCurrency(stm)
	return nil
}

func (h *queryHandler) queryJoinFiat(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcurrency.FieldFiatID),
			t.C(entfiat.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entfiat.FieldName), "fiat_name"),
			sql.As(t.C(entfiat.FieldLogo), "fiat_logo"),
			sql.As(t.C(entfiat.FieldUnit), "fiat_unit"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinFiat(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.FeedType = basetypes.CurrencyFeedType(basetypes.CurrencyFeedType_value[info.FeedTypeStr])
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
		if err := handler.queryCurrency(cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCurrencies(ctx context.Context) ([]*npool.Currency, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCurrencies(ctx, cli); err != nil {
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
