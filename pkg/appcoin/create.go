package appcoin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	appcoinmgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin"
	appexratemgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin/exrate"
	appcoinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
	appexratemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"
)

func CreateCoin(ctx context.Context, in *npool.CoinReq) (*npool.Coin, error) {
	var id string
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "coin", "coin", "CreateJoin")

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := appcoinmgrcrud.CreateSet(
			tx.AppCoin.Create(),
			&appcoinmgrpb.AppCoinReq{
				ID:                       in.ID,
				AppID:                    in.AppID,
				CoinTypeID:               in.CoinTypeID,
				Name:                     in.Name,
				DisplayNames:             in.DisplayNames,
				Logo:                     in.Logo,
				ForPay:                   in.ForPay,
				ProductPage:              in.ProductPage,
				WithdrawAutoReviewAmount: in.WithdrawAutoReviewAmount,
				DailyRewardAmount:        in.DailyRewardAmount,
				Display:                  in.Display,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		id = info.ID.String()

		_, err = appexratemgrcrud.CreateSet(
			tx.ExchangeRate.Create(),
			&appexratemgrpb.ExchangeRateReq{
				AppID:         in.AppID,
				CoinTypeID:    in.CoinTypeID,
				MarketValue:   in.MarketValue,
				SettlePercent: in.SettlePercent,
				SettleTips:    in.SettleTips,
				Setter:        in.Setter,
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
