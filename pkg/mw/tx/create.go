package tx

import (
	"context"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateTx(ctx context.Context) (*npool.Tx, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := txcrud.CreateSet(
			cli.Tran.Create(),
			&txcrud.Req{
				ID:            h.ID,
				CoinTypeID:    h.CoinTypeID,
				FromAccountID: h.FromAccountID,
				ToAccountID:   h.ToAccountID,
				Amount:        h.Amount,
				FeeAmount:     h.FeeAmount,
				ChainTxID:     h.ChainTxID,
				State:         h.State,
				Extra:         h.Extra,
				Type:          h.Type,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTx(ctx)
}

func (h *Handler) CreateTxs(ctx context.Context) ([]*npool.Tx, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			info, err := txcrud.CreateSet(tx.Tran.Create(), req).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, info.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &txcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetTxs(ctx)
	return infos, err
}
