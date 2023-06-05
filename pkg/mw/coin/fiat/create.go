package coinfiat

import (
	"context"
	"fmt"

	coinfiatcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/fiat"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) CreateCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoinFiat,
		*h.CoinTypeID,
		*h.FiatID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &coinfiatcrud.Conds{
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		FiatID:     &cruder.Cond{Op: cruder.EQ, Val: *h.FiatID},
	}
	exist, err := h.ExistCoinConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coinfiat exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := coinfiatcrud.CreateSet(
			cli.CoinFiat.Create(),
			&coinfiatcrud.Req{
				CoinTypeID: h.CoinTypeID,
				FiatID:     h.FiatID,
				FeedType:   h.FeedType,
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

	return h.GetCoinFiat(ctx)
}
