package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	basecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/base"
	basemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entextra "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	entsetting "github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"

	"github.com/google/uuid"
)

func GetCoin(ctx context.Context, id string) (*npool.Coin, error) {
	infos := []*npool.Coin{}
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return cli.
			CoinBase.
			Query().
			Where(
				entbase.ID(uuid.MustParse(id)),
			).
			Select(
				entbase.FieldID,
				entbase.FieldName,
				entbase.FieldLogo,
				entbase.FieldPresale,
				entbase.FieldForPay,
				entbase.FieldUnit,
				entbase.FieldEnv,
				entbase.FieldReservedAmount,
				entbase.FieldCreatedAt,
				entbase.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entextra.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entbase.FieldID),
						t1.C(entextra.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t1.C(entextra.FieldHomePage), "home_page"),
						sql.As(t1.C(entextra.FieldSpecs), "specs"),
					)

				t2 := sql.Table(entsetting.Table)
				s.
					LeftJoin(t2).
					On(
						s.C(entbase.FieldID),
						t2.C(entsetting.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t2.C(entsetting.FieldFeeCoinTypeID), "fee_coin_type_id"),
						sql.As(t2.C(entsetting.FieldWithdrawFeeByStableUsd), "withdraw_fee_by_stable_usd"),
						sql.As(t2.C(entsetting.FieldWithdrawFeeAmount), "withdraw_fee_amount"),
						sql.As(t2.C(entsetting.FieldCollectFeeAmount), "collect_fee_amount"),
						sql.As(t2.C(entsetting.FieldHotWalletFeeAmount), "hot_wallet_fee_amount"),
						sql.As(t2.C(entsetting.FieldLowFeeAmount), "low_fee_amount"),
						sql.As(t2.C(entsetting.FieldHotWalletAccountAmount), "hot_wallet_account_amount"),
						sql.As(t2.C(entsetting.FieldPaymentAccountCollectAmount), "payment_account_collect_amount"),
					)
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}

	return infos[0], nil
}

func GetCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Coin, total uint32, err error) {
	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := basecrud.SetQueryConds(&basemgrpb.Conds{
			ID:      conds.ID,
			Presale: conds.Presale,
			ENV:     conds.ENV,
			ForPay:  conds.ForPay,
		}, cli)
		if err != nil {
			return err
		}

		stm.
			Select(
				entbase.FieldID,
				entbase.FieldName,
				entbase.FieldLogo,
				entbase.FieldPresale,
				entbase.FieldForPay,
				entbase.FieldUnit,
				entbase.FieldEnv,
				entbase.FieldReservedAmount,
				entbase.FieldCreatedAt,
				entbase.FieldUpdatedAt,
			)

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return stm.
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entextra.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entbase.FieldID),
						t1.C(entextra.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t1.C(entextra.FieldHomePage), "home_page"),
						sql.As(t1.C(entextra.FieldSpecs), "specs"),
					)

				t2 := sql.Table(entsetting.Table)
				s.
					LeftJoin(t2).
					On(
						s.C(entbase.FieldID),
						t2.C(entsetting.FieldCoinTypeID),
					)

				s.
					AppendSelect(
						sql.As(t2.C(entsetting.FieldFeeCoinTypeID), "fee_coin_type_id"),
						sql.As(t2.C(entsetting.FieldWithdrawFeeByStableUsd), "withdraw_fee_by_stable_usd"),
						sql.As(t2.C(entsetting.FieldWithdrawFeeAmount), "withdraw_fee_amount"),
						sql.As(t2.C(entsetting.FieldCollectFeeAmount), "collect_fee_amount"),
						sql.As(t2.C(entsetting.FieldHotWalletFeeAmount), "hot_wallet_fee_amount"),
						sql.As(t2.C(entsetting.FieldLowFeeAmount), "low_fee_amount"),
						sql.As(t2.C(entsetting.FieldHotWalletAccountAmount), "hot_wallet_account_amount"),
						sql.As(t2.C(entsetting.FieldPaymentAccountCollectAmount), "payment_account_collect_amount"),
					)
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
