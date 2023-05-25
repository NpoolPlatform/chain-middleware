package tx

import (
	"context"
	"fmt"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) UpdateTx(ctx context.Context) (*npool.Tx, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := txcrud.UpdateSet(
			cli.Tran.UpdateOneID(*h.ID),
			&txcrud.Req{
				ChainTxID: h.ChainTxID,
				State:     h.State,
				Extra:     h.Extra,
			},
		)
		if err != nil {
			return err
		}
		if _, err := stm.Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}
