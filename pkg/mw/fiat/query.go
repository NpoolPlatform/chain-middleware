package fiat

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	fiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/fiat"
	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	entfiat "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiat"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatSelect
	infos []*npool.Fiat
	total uint32
}

func (h *queryHandler) selectFiat(stm *ent.FiatQuery) {
	h.stm = stm.Select(
		entfiat.FieldID,
		entfiat.FieldName,
		entfiat.FieldLogo,
		entfiat.FieldUnit,
		entfiat.FieldCreatedAt,
		entfiat.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryFiat(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	h.selectFiat(
		cli.Fiat.
			Query().
			Where(
				entfiat.ID(*h.ID),
			),
	)
	return nil
}

func (h *queryHandler) queryFiats(ctx context.Context, cli *ent.Client) error {
	stm, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), h.Conds)
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

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetFiat(ctx context.Context) (*npool.Fiat, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiat(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.Offset(0).Limit(2)
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

	return handler.infos[0], nil
}

func (h *Handler) GetFiats(ctx context.Context) ([]*npool.Fiat, uint32, error) {
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

	return handler.infos, handler.total, nil
}

func (h *Handler) GetFiatOnly(ctx context.Context) (info *npool.Fiat, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiats(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.Offset(0).Limit(2)
		return handler.scan(_ctx)
	})

	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
