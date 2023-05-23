package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	basecrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin"
	extracrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin/setting"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entextra "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinextra"
	entsetting "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/setting"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateCoinBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := basecrud.UpdateSet(
		tx.CoinBase.UpdateOneID(*h.ID),
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

func (h *updateHandler) updateCoinExtra(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		CoinExtra.
		Query().
		Where(
			entextra.CoinTypeID(*h.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		if _, err := extracrud.UpdateSet(
			info.Update(),
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

func (h *updateHandler) updateCoinSetting(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Setting.
		Query().
		Where(
			entsetting.CoinTypeID(*h.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info != nil {
		if _, err := settingcrud.UpdateSet(
			info.Update(),
			&settingcrud.Req{
				CoinTypeID:                  h.ID,
				FeeCoinTypeID:               h.FeeCoinTypeID,
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

	if _, err := settingcrud.CreateSet(
		tx.Setting.Create(),
		&settingcrud.Req{
			CoinTypeID:                  h.ID,
			FeeCoinTypeID:               h.FeeCoinTypeID,
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

func (h *Handler) UpdateCoin(ctx context.Context) (*npool.Coin, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid cointypeid")
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCoinBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateCoinExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateCoinSetting(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoin(ctx)
}
