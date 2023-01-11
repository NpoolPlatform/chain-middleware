package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	basecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/base"
	extracrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/extra"
	settingcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/setting"
	basemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"
	extramgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/extra"
	settingmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/setting"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entextra "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	entsetting "github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"

	"github.com/google/uuid"
)

func UpdateCoin(ctx context.Context, in *npool.CoinReq) (*npool.Coin, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "coin", "coin", "UpdateJoin")

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.CoinBase.Query().Where(entbase.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = basecrud.UpdateSet(info, &basemgrpb.CoinBaseReq{
			ID:             in.ID,
			Logo:           in.Logo,
			Presale:        in.Presale,
			ReservedAmount: in.ReservedAmount,
			ForPay:         in.ForPay,
			Disabled:       in.Disabled,
		}).Save(_ctx)
		if err != nil {
			return err
		}

		info1, err := tx.CoinExtra.Query().Where(entextra.CoinTypeID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = extracrud.UpdateSet(info1, &extramgrpb.CoinExtraReq{
			HomePage:  in.HomePage,
			Specs:     in.Specs,
			StableUSD: in.StableUSD,
		}).Save(_ctx)
		if err != nil {
			return err
		}

		info2, err := tx.Setting.Query().Where(entsetting.CoinTypeID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = settingcrud.UpdateSet(info2, &settingmgrpb.SettingReq{
			FeeCoinTypeID:               in.FeeCoinTypeID,
			WithdrawFeeByStableUSD:      in.WithdrawFeeByStableUSD,
			WithdrawFeeAmount:           in.WithdrawFeeAmount,
			CollectFeeAmount:            in.CollectFeeAmount,
			HotWalletFeeAmount:          in.HotWalletFeeAmount,
			LowFeeAmount:                in.LowFeeAmount,
			HotWalletAccountAmount:      in.HotWalletAccountAmount,
			PaymentAccountCollectAmount: in.PaymentAccountCollectAmount,
			LeastTransferAmount:         in.LeastTransferAmount,
		}).Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetCoin(ctx, in.GetID())
}
