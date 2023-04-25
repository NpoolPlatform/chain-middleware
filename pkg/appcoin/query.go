package appcoin

import (
	"context"
	"encoding/json"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/appcoin"
	appcoinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin"

	entappcoin "github.com/NpoolPlatform/chain-manager/pkg/db/ent/appcoin"
	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entcoinextra "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	entappexrate "github.com/NpoolPlatform/chain-manager/pkg/db/ent/exchangerate"
	entsetting "github.com/NpoolPlatform/chain-manager/pkg/db/ent/setting"

	"github.com/google/uuid"
)

func GetCoin(ctx context.Context, id string) (*npool.Coin, error) {
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
		stm := cli.
			AppCoin.
			Query().
			Where(
				entappcoin.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	infos = expand(infos)

	return infos[0], nil
}

func GetCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Coin, uint32, error) {
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

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoins")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&appcoinmgrpb.Conds{
			ID:          conds.ID,
			AppID:       conds.AppID,
			CoinTypeID:  conds.CoinTypeID,
			ForPay:      conds.ForPay,
			Disabled:    conds.Disabled,
			IDs:         conds.IDs,
			CoinTypeIDs: conds.CoinTypeIDs,
		}, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm.
			Order(ent.Asc(entappcoin.FieldDisplayIndex)).
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}
	infos = expand(infos)

	return infos, total, nil
}

func GetCoinOnly(ctx context.Context, conds *npool.Conds) (*npool.Coin, error) {
	var infos []*npool.Coin
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoins")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&appcoinmgrpb.Conds{
			ID:         conds.ID,
			AppID:      conds.AppID,
			CoinTypeID: conds.CoinTypeID,
			ForPay:     conds.ForPay,
			Disabled:   conds.Disabled,
		}, cli)
		if err != nil {
			return err
		}

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	infos = expand(infos)

	return infos[0], nil
}

func expand(infos []*npool.Coin) []*npool.Coin {
	for _, info := range infos {
		fmt.Printf("dn: %v, st: %v\n", info.DisplayNamesStr, info.SettleTipsStr)
		_ = json.Unmarshal([]byte(info.DisplayNamesStr), &info.DisplayNames)
		if !info.CoinForPay {
			info.ForPay = info.CoinForPay
		}
		if !info.Disabled {
			info.Disabled = info.CoinDisabled
		}
		if info.MarketValue == "" {
			info.MarketValue = decimal.NewFromInt(0).String()
		}
		if info.SettleValue == "" {
			info.SettleValue = decimal.NewFromInt(0).String()
		}
		if info.MaxAmountPerWithdraw == "" {
			info.MaxAmountPerWithdraw = decimal.NewFromInt(0).String()
		}
		_ = json.Unmarshal([]byte(info.SettleTipsStr), &info.SettleTips)
	}
	return infos
}

func join(stm *ent.AppCoinQuery) *ent.AppCoinSelect { //nolint:funlen
	return stm.
		Select(
			entappcoin.FieldID,
			entappcoin.FieldAppID,
			entappcoin.FieldCoinTypeID,
			entappcoin.FieldName,
			entappcoin.FieldDisplayNames,
			entappcoin.FieldLogo,
			entappcoin.FieldForPay,
			entappcoin.FieldWithdrawAutoReviewAmount,
			entappcoin.FieldProductPage,
			entappcoin.FieldDisabled,
			entappcoin.FieldCreatedAt,
			entappcoin.FieldUpdatedAt,
			entappcoin.FieldDisplay,
			entappcoin.FieldDailyRewardAmount,
			entappcoin.FieldDisplayIndex,
			entappcoin.FieldMaxAmountPerWithdraw,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entcoinextra.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entappcoin.FieldCoinTypeID),
					t1.C(entcoinextra.FieldCoinTypeID),
				).
				OnP(
					sql.EQ(t1.C(entcoinextra.FieldDeletedAt), 0),
				).
				AppendSelect(
					sql.As(t1.C(entcoinextra.FieldHomePage), "home_page"),
					sql.As(t1.C(entcoinextra.FieldSpecs), "specs"),
					sql.As(t1.C(entcoinextra.FieldStableUsd), "stable_usd"),
				)

			t2 := sql.Table(entsetting.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(entappcoin.FieldCoinTypeID),
					t2.C(entsetting.FieldCoinTypeID),
				).
				OnP(
					sql.EQ(t2.C(entsetting.FieldDeletedAt), 0),
				).
				AppendSelect(
					sql.As(t2.C(entsetting.FieldFeeCoinTypeID), "fee_coin_type_id"),
					sql.As(t2.C(entsetting.FieldWithdrawFeeByStableUsd), "withdraw_fee_by_stable_usd"),
					sql.As(t2.C(entsetting.FieldWithdrawFeeAmount), "withdraw_fee_amount"),
					sql.As(t2.C(entsetting.FieldCollectFeeAmount), "collect_fee_amount"),
					sql.As(t2.C(entsetting.FieldHotWalletFeeAmount), "hot_wallet_fee_amount"),
					sql.As(t2.C(entsetting.FieldLowFeeAmount), "low_fee_amount"),
					sql.As(t2.C(entsetting.FieldHotLowFeeAmount), "hot_low_fee_amount"),
					sql.As(t2.C(entsetting.FieldHotWalletAccountAmount), "hot_wallet_account_amount"),
					sql.As(t2.C(entsetting.FieldPaymentAccountCollectAmount), "payment_account_collect_amount"),
					sql.As(t2.C(entsetting.FieldLeastTransferAmount), "least_transfer_amount"),
					sql.As(t2.C(entsetting.FieldNeedMemo), "need_memo"),
				)

			t3 := sql.Table(entcoinbase.Table)
			s.
				LeftJoin(t3).
				On(
					t2.C(entsetting.FieldFeeCoinTypeID),
					t3.C(entcoinbase.FieldID),
				).
				OnP(
					sql.EQ(t3.C(entcoinbase.FieldDeletedAt), 0),
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
				OnP(
					sql.EQ(t4.C(entcoinbase.FieldDeletedAt), 0),
				).
				AppendSelect(
					sql.As(t4.C(entcoinbase.FieldName), "coin_name"),
					sql.As(t4.C(entcoinbase.FieldUnit), "unit"),
					sql.As(t4.C(entcoinbase.FieldEnv), "env"),
					sql.As(t4.C(entcoinbase.FieldPresale), "presale"),
					sql.As(t4.C(entcoinbase.FieldReservedAmount), "reserved_amount"),
					sql.As(t4.C(entcoinbase.FieldDisabled), "coin_disabled"),
					sql.As(t4.C(entcoinbase.FieldForPay), "coin_for_pay"),
				)

			t5 := sql.Table(entappexrate.Table)
			s.
				LeftJoin(t5).
				On(
					s.C(entappcoin.FieldCoinTypeID),
					t5.C(entappexrate.FieldCoinTypeID),
				).
				On(
					s.C(entappcoin.FieldAppID),
					t5.C(entappexrate.FieldAppID),
				).
				OnP(
					sql.EQ(t5.C(entappexrate.FieldDeletedAt), 0),
				).
				AppendSelect(
					sql.As(t5.C(entappexrate.FieldMarketValue), "market_value"),
					sql.As(t5.C(entappexrate.FieldSettleValue), "settle_value"),
					sql.As(t5.C(entappexrate.FieldSettlePercent), "settle_percent"),
					sql.As(t5.C(entappexrate.FieldSettleTips), "settle_tips"),
					sql.As(t5.C(entappexrate.FieldSetter), "setter"),
				)
		})
}
