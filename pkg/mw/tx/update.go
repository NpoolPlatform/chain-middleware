package tx

import (
	"context"
	"fmt"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	enttran "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/tran"
)

func (h *Handler) UpdateTx(ctx context.Context) (*npool.Tx, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Tran.
			Query().
			Where(
				enttran.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		if req.State != nil {
			switch info.State {
			case basetypes.TxState_TxStateCreated.String():
				switch *req.State {
				case basetypes.TxState_TxStateWait:
				default:
					return fmt.Errorf("state is invalid: %v -> %v", info.State, req.State)
				}
			case basetypes.TxState_TxStateWait.String():
				switch *req.State {
				case basetypes.TxState_TxStateTransferring:
				default:
					return fmt.Errorf("state is invalid: %v -> %v", info.State, req.State)
				}
			case basetypes.TxState_TxStateTransferring.String():
				switch *req.State {
				case basetypes.TxState_TxStateSuccessful:
				case basetypes.TxState_TxStateFail:
				default:
					return fmt.Errorf("state is invalid: %v -> %v", info.State, req.State)
				}
			case basetypes.TxState_TxStateSuccessful.String():
				fallthrough //nolint
			case basetypes.TxState_TxStateFail.String():
				fallthrough //nolint
			default:
				return fmt.Errorf("state is invalid: %v -> %v", info.State, req.State)
			}
		}

		stm, err := txcrud.UpdateSet(
			info.Update(),
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
