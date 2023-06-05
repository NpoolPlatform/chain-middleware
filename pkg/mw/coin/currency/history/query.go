package currencyhistory

import (
	"context"
	"fmt"

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

func (h *queryHandler) queryJoinCoin(s *sql.Selector) error {
	t := sql.Table(entcoinbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entcurrencyhis.FieldCoinTypeID),
			t.C(entcoinbase.FieldID),
		)

	if h.Conds.CoinNames != nil {
		names, ok := h.Conds.CoinNames.Val.([]string)
		if !ok {
			return fmt.Errorf("invalid coinnames")
		}
		s.OnP(sql.In(t.C(entcoinbase.FieldName), names))
	}

	s.AppendSelect(
		sql.As(t.C(entcoinbase.FieldName), "coin_name"),
		sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
		sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
		sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
	)
	return nil
}

func (h *queryHandler) queryJoin() (err error) {
	h.stm.Modify(func(s *sql.Selector) {
		if err = h.queryJoinCoin(s); err != nil {
			return
		}
	})
	return err
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
		if err := handler.queryJoin(); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
