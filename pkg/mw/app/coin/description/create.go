package description

import (
	"context"
	"fmt"

	descriptioncrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/description"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateCoinDescription(ctx context.Context) (*npool.CoinDescription, error) {
	lockKey := fmt.Sprintf(
		"%v:%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoinDescription,
		*h.AppID,
		*h.CoinTypeID,
		*h.UsedFor,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	h.Conds = &descriptioncrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		UsedFor:    &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistCoinDescriptionConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("description exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := descriptioncrud.CreateSet(
			cli.CoinDescription.Create(),
			&descriptioncrud.Req{
				ID:         h.ID,
				AppID:      h.AppID,
				CoinTypeID: h.CoinTypeID,
				UsedFor:    h.UsedFor,
				Title:      h.Title,
				Message:    h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCoinDescription(ctx)
}
