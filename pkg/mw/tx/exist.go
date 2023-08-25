package tx

import (
	"context"
	"fmt"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	enttran "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/tran"
)

func (h *Handler) ExistTx(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Tran.
			Query().
			Where(
				enttran.ID(*h.ID),
				enttran.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistTxConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := txcrud.SetQueryConds(cli.Tran.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
