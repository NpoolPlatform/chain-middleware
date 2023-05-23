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

func CreateCoin(ctx context.Context, in *npool.CoinReq) (*npool.Coin, error) {
	var id string
	var err error

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := basecrud.CreateSet(
			tx.CoinBase.Create(),
			&basemgrpb.CoinBaseReq{
				ID:             in.ID,
				Name:           in.Name,
				Logo:           in.Logo,
				Presale:        in.Presale,
				Unit:           in.Unit,
				ENV:            in.ENV,
				ReservedAmount: in.ReservedAmount,
				ForPay:         in.ForPay,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		id = info.ID.String()

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
