//nolint:dupl
package appcoin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entappcoin "github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entcoinextra "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	entappexrate "github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"
	entsetting "github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"

	"github.com/google/uuid"
)

func GetCoin(ctx context.Context, id string) (*npool.Coin, error) { //nolint
	var infos []*npool.Coin
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
			AppCoin.
			Query().
			Where(
				entappcoin.ID(uuid.MustParse(id)),
			).
			Select(
				entappcoin.FieldID,
				entappcoin.FieldAppID,
				entappcoin.FieldCoinTypeID,
				entappcoin.FieldName,
				entappcoin.FieldLogo,
				entappcoin.FieldForPay,
				entappcoin.FieldWithdrawAutoReviewAmount,
				entappcoin.FieldCreatedAt,
				entappcoin.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entcoinextra.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t1.C(entcoinextra.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t1.C(entcoinextra.FieldHomePage), "home_page"),
						sql.As(t1.C(entcoinextra.FieldSpecs), "specs"),
					)

				t2 := sql.Table(entsetting.Table)
				s.
					LeftJoin(t2).
					On(
						s.C(entappcoin.FieldCoinTypeID),
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

				t3 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t3).
					On(
						t2.C(entsetting.FieldFeeCoinTypeID),
						t3.C(entcoinbase.FieldID),
					).
					AppendSelect(
						sql.As(t3.C(entcoinbase.FieldName), "fee_coin_name"),
						sql.As(t3.C(entcoinbase.FieldLogo), "fee_coin_logo"),
						sql.As(t3.C(entcoinbase.FieldUnit), "fee_coin_unit"),
						sql.As(t3.C(entcoinbase.FieldEnv), "fee_coin_env"),
					)

				t4 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t4).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t4.C(entcoinbase.FieldID),
					).
					AppendSelect(
						sql.As(t4.C(entcoinbase.FieldUnit), "unit"),
						sql.As(t4.C(entcoinbase.FieldEnv), "env"),
						sql.As(t4.C(entcoinbase.FieldPresale), "presale"),
						sql.As(t4.C(entcoinbase.FieldReservedAmount), "reserved_amount"),
					)

				t5 := sql.Table(entappexrate.Table)
				s.
					LeftJoin(t5).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t5.C(entappexrate.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t5.C(entappexrate.FieldMarketValue), "market_value"),
						sql.As(t5.C(entappexrate.FieldSettleValue), "settle_value"),
						sql.As(t5.C(entappexrate.FieldSettlePercent), "settle_percent"),
					)
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	return infos[0], nil
}

func GetCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coin, uint32, error) { //nolint
	var infos []*npool.Coin
	var err error
	var total uint32

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	ids := []uuid.UUID{}
	for _, id := range conds.GetIDs().GetValue() {
		ids = append(ids, uuid.MustParse(id))
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoins")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppCoin.
			Query()

		if conds.ID != nil {
			stm.Where(
				entappcoin.ID(uuid.MustParse(conds.GetID().GetValue())),
			)
		}
		if len(ids) > 0 {
			stm.Where(
				entappcoin.IDIn(ids...),
			)
		}
		if conds.AppID != nil {
			stm.Where(
				entappcoin.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
			)
		}
		if conds.CoinTypeID != nil {
			stm.Where(
				entappcoin.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())),
			)
		}
		if conds.ForPay != nil {
			stm.Where(
				entappcoin.ForPay(conds.GetForPay().GetValue()),
			)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return stm.
			Select(
				entappcoin.FieldID,
				entappcoin.FieldAppID,
				entappcoin.FieldCoinTypeID,
				entappcoin.FieldName,
				entappcoin.FieldLogo,
				entappcoin.FieldForPay,
				entappcoin.FieldWithdrawAutoReviewAmount,
				entappcoin.FieldCreatedAt,
				entappcoin.FieldUpdatedAt,
			).
			Offset(int(offset)).
			Limit(int(limit)).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entcoinextra.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t1.C(entcoinextra.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t1.C(entcoinextra.FieldHomePage), "home_page"),
						sql.As(t1.C(entcoinextra.FieldSpecs), "specs"),
					)

				t2 := sql.Table(entsetting.Table)
				s.
					LeftJoin(t2).
					On(
						s.C(entappcoin.FieldCoinTypeID),
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

				t3 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t3).
					On(
						t2.C(entsetting.FieldFeeCoinTypeID),
						t3.C(entcoinbase.FieldID),
					).
					AppendSelect(
						sql.As(t3.C(entcoinbase.FieldName), "fee_coin_name"),
						sql.As(t3.C(entcoinbase.FieldLogo), "fee_coin_logo"),
						sql.As(t3.C(entcoinbase.FieldUnit), "fee_coin_unit"),
						sql.As(t3.C(entcoinbase.FieldEnv), "fee_coin_env"),
					)

				t4 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t4).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t4.C(entcoinbase.FieldID),
					).
					AppendSelect(
						sql.As(t4.C(entcoinbase.FieldUnit), "unit"),
						sql.As(t4.C(entcoinbase.FieldEnv), "env"),
						sql.As(t4.C(entcoinbase.FieldPresale), "presale"),
						sql.As(t4.C(entcoinbase.FieldReservedAmount), "reserved_amount"),
					)

				t5 := sql.Table(entappexrate.Table)
				s.
					LeftJoin(t5).
					On(
						s.C(entappcoin.FieldCoinTypeID),
						t5.C(entappexrate.FieldCoinTypeID),
					).
					AppendSelect(
						sql.As(t5.C(entappexrate.FieldMarketValue), "market_value"),
						sql.As(t5.C(entappexrate.FieldSettleValue), "settle_value"),
						sql.As(t5.C(entappexrate.FieldSettlePercent), "settle_percent"),
					)
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
