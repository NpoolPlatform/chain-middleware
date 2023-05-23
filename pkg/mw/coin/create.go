package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	basecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/base"
	extracrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/setting"
	basemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
	extramgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"
	settingmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
)

func (h *Handler) CreateCoin(ctx context.Context) (*npool.Coin, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := basecrud.CreateSet(
			tx.CoinBase.Create(),
			&basemgrpb.CoinBaseReq{
				ID:             h.ID,
				Name:           h.Name,
				Logo:           h.Logo,
				Presale:        h.Presale,
				Unit:           h.Unit,
				ENV:            h.ENV,
				ReservedAmount: h.ReservedAmount,
				ForPay:         h.ForPay,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		_, err = extracrud.CreateSet(
			tx.CoinExtra.Create(),
			&extramgrpb.CoinExtraReq{
				CoinTypeID: &id,
				HomePage:   in.HomePage,
				Specs:      in.Specs,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		_, err = settingcrud.CreateSet(
			tx.Setting.Create(),
			&settingmgrpb.SettingReq{
				CoinTypeID:                  &id,
				FeeCoinTypeID:               &id,
				WithdrawFeeByStableUSD:      in.WithdrawFeeByStableUSD,
				WithdrawFeeAmount:           in.WithdrawFeeAmount,
				CollectFeeAmount:            in.CollectFeeAmount,
				HotWalletFeeAmount:          in.HotWalletFeeAmount,
				LowFeeAmount:                in.LowFeeAmount,
				HotLowFeeAmount:             in.HotLowFeeAmount,
				HotWalletAccountAmount:      in.HotWalletAccountAmount,
				PaymentAccountCollectAmount: in.PaymentAccountCollectAmount,
				LeastTransferAmount:         in.LeastTransferAmount,
				NeedMemo:                    in.NeedMemo,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetCoin(ctx, id)
}
