package currencyhistory

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	historycrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/currency/history"
	entcoinbase "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinbase"
	entcurrencyhis "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyhistory"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.CurrencyHistorySelect
	infos []*npool.Currency
	total uint32
}

func (h *queryHandler) selectCurrencyHistory(stm *ent.CurrencyHistoryQuery) {
	h.stm = stm.Select(
		entcurrencyhis.FieldID,
		entcurrencyhis.FieldCoinTypeID,
		entcurrencyhis.FieldFeedType,
		entcurrencyhis.FieldMarketValueHigh,
		entcurrencyhis.FieldMarketValueLow,
		entcurrencyhis.FieldCreatedAt,
		entcurrencyhis.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCurrencyHistories(ctx context.Context, cli *ent.Client) error {
	stm, err := historycrud.SetQueryConds(cli.CurrencyHistory.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectCurrencyHistory(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t1 := sql.Table(entcoinbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcurrencyhis.FieldCoinTypeID),
			t1.C(entcoinbase.FieldID),
		).
		AppendSelect(
			sql.As(t1.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t1.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t1.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t1.C(entcoinbase.FieldEnv), "coin_env"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoin(s)
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

func (h *Handler) GetCurrencies(ctx context.Context) ([]*npool.Currency, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCurrencyHistories(_ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Order(ent.Desc(entcurrencyhis.FieldCreatedAt)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
