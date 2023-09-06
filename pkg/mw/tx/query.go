package tx

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	enttx "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/tran"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	entcoinbase "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinbase"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stm   *ent.TranSelect
	infos []*npool.Tx
	total uint32
}

func (h *queryHandler) selectTx(stm *ent.TranQuery) {
	h.stm = stm.Select(
		enttx.FieldID,
		enttx.FieldCoinTypeID,
		enttx.FieldFromAccountID,
		enttx.FieldToAccountID,
		enttx.FieldAmount,
		enttx.FieldFeeAmount,
		enttx.FieldState,
		enttx.FieldChainTxID,
		enttx.FieldType,
		enttx.FieldExtra,
		enttx.FieldCreatedAt,
		enttx.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTx(cli *ent.Client) error {
	h.selectTx(
		cli.Tran.
			Query().
			Where(
				enttx.ID(*h.ID),
			),
	)
	return nil
}

func (h *queryHandler) queryTxs(ctx context.Context, cli *ent.Client) error {
	stm, err := txcrud.SetQueryConds(cli.Tran.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)

	h.selectTx(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.
		LeftJoin(t).
		On(
			s.C(enttx.FieldCoinTypeID),
			t.C(entcoinbase.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
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
		info.Type = basetypes.TxType(basetypes.TxType_value[info.TypeStr])
		info.State = basetypes.TxState(basetypes.TxState_value[info.StateStr])
		amount, err := decimal.NewFromString(info.Amount)
		if err != nil {
			info.Amount = decimal.NewFromInt(0).String()
		} else {
			info.Amount = amount.String()
		}
		amount, err = decimal.NewFromString(info.FeeAmount)
		if err != nil {
			info.FeeAmount = decimal.NewFromInt(0).String()
		} else {
			info.FeeAmount = amount.String()
		}
	}
}

func (h *Handler) GetTx(ctx context.Context) (*npool.Tx, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTx(cli); err != nil {
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

func (h *Handler) GetTxs(ctx context.Context) ([]*npool.Tx, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTxs(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(enttx.FieldUpdatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
