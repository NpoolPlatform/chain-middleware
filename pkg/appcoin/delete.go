package appcoin

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entappcoin "github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	entdescription "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coindescription"
	entappexrate "github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"

	appcoinmgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin"
	appdescriptionmgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin/description"
	appexratemgrcrud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin/exrate"
	appcoinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"
	appdescriptionmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	appexratemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/exrate"

	"github.com/google/uuid"
)

func DeleteCoin(ctx context.Context, id string) (*npool.Coin, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	ret, err := GetCoin(ctx, id)
	if err != nil {
		return nil, err
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "DeleteJoin")
	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.AppCoin.Query().Where(entappcoin.ID(uuid.MustParse(id))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = appcoinmgrcrud.UpdateSet(
			info,
			&appcoinmgrpb.AppCoinReq{
				DeletedAt: &now,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		info1, err := tx.ExchangeRate.Query().Where(
			entappexrate.AppID(info.AppID),
			entappexrate.CoinTypeID(info.CoinTypeID),
		).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		_, err = appexratemgrcrud.UpdateSet(
			info1,
			&appexratemgrpb.ExchangeRateReq{
				DeletedAt: &now,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		infos, err := tx.CoinDescription.Query().Where(
			entdescription.AppID(info.AppID),
			entdescription.CoinTypeID(info.CoinTypeID),
		).ForUpdate().All(_ctx)
		if err != nil {
			return err
		}

		for _, info2 := range infos {
			_, err = appdescriptionmgrcrud.UpdateSet(
				info2,
				&appdescriptionmgrpb.CoinDescriptionReq{
					DeletedAt: &now,
				},
			).Save(_ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}
