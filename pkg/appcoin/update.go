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

	entappcoin "github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	entappexrate "github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"

	appcoinmgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin"
	appexratemgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin/exrate"
	appcoinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
	appexratemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"

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
		info, err := tx.AppCoin.Query().Where(entappcoin.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = appcoinmgrcrud.UpdateSet(
			info,
			&appcoinmgrpb.AppCoinReq{
				Name:                     in.Name,
				Logo:                     in.Logo,
				ForPay:                   in.ForPay,
				WithdrawAutoReviewAmount: in.WithdrawAutoReviewAmount,
				Disabled:                 in.Disabled,
				DailyRewardAmount:        in.DailyRewardAmount,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		info1, err := tx.ExchangeRate.Query().Where(
			entappexrate.AppID(uuid.MustParse(in.GetAppID())),
			entappexrate.CoinTypeID(uuid.MustParse(in.GetCoinTypeID())),
		).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = appexratemgrcrud.UpdateSet(
			info1,
			&appexratemgrpb.ExchangeRateReq{
				MarketValue:   in.MarketValue,
				SettlePercent: in.SettlePercent,
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

	return GetCoin(ctx, in.GetID())
}
