package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	basecrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin"
	extracrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/setting"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCoinBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := basecrud.CreateSet(
		tx.CoinBase.Create(),
		&basecrud.Req{
			ID:             h.ID,
			Name:           h.Name,
			Logo:           h.Logo,
			Presale:        h.Presale,
			Unit:           h.Unit,
			ENV:            h.ENV,
			ReservedAmount: h.ReservedAmount,
			ForPay:         h.ForPay,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createCoinExtra(ctx context.Context, tx *ent.Tx) error {
	if _, err := extracrud.CreateSet(
		tx.CoinExtra.Create(),
		&extracrud.Req{
			CoinTypeID: h.ID,
			HomePage:   h.HomePage,
			Specs:      h.Specs,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createCoinSetting(ctx context.Context, tx *ent.Tx) error {
	if _, err := settingcrud.CreateSet(
		tx.Setting.Create(),
		&settingcrud.Req{
			CoinTypeID:                  h.ID,
			FeeCoinTypeID:               h.ID,
			WithdrawFeeByStableUSD:      h.WithdrawFeeByStableUSD,
			WithdrawFeeAmount:           h.WithdrawFeeAmount,
			CollectFeeAmount:            h.CollectFeeAmount,
			HotWalletFeeAmount:          h.HotWalletFeeAmount,
			LowFeeAmount:                h.LowFeeAmount,
			HotLowFeeAmount:             h.HotLowFeeAmount,
			HotWalletAccountAmount:      h.HotWalletAccountAmount,
			PaymentAccountCollectAmount: h.PaymentAccountCollectAmount,
			LeastTransferAmount:         h.LeastTransferAmount,
			NeedMemo:                    h.NeedMemo,
			RefreshCurrency:             h.RefreshCurrency,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.Name == nil {
		return nil, fmt.Errorf("invalid coinname")
	}
	if h.Unit == nil {
		return nil, fmt.Errorf("invalid coinunit")
	}
	if h.ENV == nil {
		return nil, fmt.Errorf("invalid coinenv")
	}

	lockKey := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoin,
		*h.Name,
		*h.ENV,
	)
	if err := redis2.TryLock(lockKey, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(lockKey)
	}()

	h.Conds = &basecrud.Conds{
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
		ENV:  &cruder.Cond{Op: cruder.EQ, Val: *h.ENV},
	}
	h.Offset = 0
	h.Limit = 2

	coin, err := h.GetCoinOnly(ctx)
	if err != nil {
		return nil, err
	}
	if coin != nil {
		if coin.Unit != *h.Unit {
			return nil, fmt.Errorf("invalid coinunit")
		}
		return coin, nil
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createCoinBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createCoinExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createCoinSetting(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
