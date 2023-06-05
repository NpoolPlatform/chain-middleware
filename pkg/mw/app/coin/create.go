package appcoin

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	appcoincrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin"
	appexratecrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/exrate"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppCoin(ctx context.Context, tx *ent.Tx) error {
	if _, err := appcoincrud.CreateSet(
		tx.AppCoin.Create(),
		&appcoincrud.Req{
			ID:                       h.ID,
			AppID:                    h.AppID,
			CoinTypeID:               h.CoinTypeID,
			Name:                     h.Name,
			DisplayNames:             h.DisplayNames,
			Logo:                     h.Logo,
			ForPay:                   h.ForPay,
			ProductPage:              h.ProductPage,
			WithdrawAutoReviewAmount: h.WithdrawAutoReviewAmount,
			DailyRewardAmount:        h.DailyRewardAmount,
			Display:                  h.Display,
			DisplayIndex:             h.DisplayIndex,
			MaxAmountPerWithdraw:     h.MaxAmountPerWithdraw,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createExrate(ctx context.Context, tx *ent.Tx) error {
	if _, err := appexratecrud.CreateSet(
		tx.ExchangeRate.Create(),
		&appexratecrud.Req{
			AppID:         h.AppID,
			CoinTypeID:    h.CoinTypeID,
			MarketValue:   h.MarketValue,
			SettlePercent: h.SettlePercent,
			SettleTips:    h.SettleTips,
			Setter:        h.Setter,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateAppCoin,
		*h.AppID,
		*h.CoinTypeID,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &appcoincrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
	}
	exist, err := h.ExistCoinConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("appcoin exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppCoin(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createExrate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
