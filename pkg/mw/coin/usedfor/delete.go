package coinusedfor

import (
	"context"
	"time"

	coinusedforcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/usedfor"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) DeleteCoinUsedFor(ctx context.Context) (*npool.CoinUsedFor, error) {
	info, err := h.GetCoinUsedFor(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := coinusedforcrud.UpdateSet(
			cli.CoinUsedFor.UpdateOneID(info.ID),
			&coinusedforcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
