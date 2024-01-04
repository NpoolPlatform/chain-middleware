package coinusedfor

import (
	"context"
	"fmt"

	coinusedforcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/usedfor"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) CreateCoinUsedFor(ctx context.Context) (*npool.CoinUsedFor, error) {
	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoinUsedFor,
		*h.CoinTypeID,
		*h.UsedFor,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &coinusedforcrud.Conds{
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		UsedFor:    &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistCoinUsedForConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coinusedfor exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := coinusedforcrud.CreateSet(
			cli.CoinUsedFor.Create(),
			&coinusedforcrud.Req{
				CoinTypeID: h.CoinTypeID,
				UsedFor:    h.UsedFor,
				Priority:   h.Priority,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoinUsedFor(ctx)
}
