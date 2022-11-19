package currencyvalue

import (
	"context"
	"fmt"

	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entfeed "github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"
	entvalue "github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyvalue"

	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/value"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func GetCurrency(ctx context.Context, id string) (*npool.Currency, error) {
	var infos []*npool.Currency

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			CurrencyValue.
			Query().
			Where(
				entvalue.ID(uuid.MustParse(id)),
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
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func GetCoinCurrency(ctx context.Context, coinTypeID string) (*npool.Currency, error) {
	var infos []*npool.Currency

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			CurrencyValue.
			Query().
			Where(
				entvalue.CoinTypeID(uuid.MustParse(coinTypeID)),
			).
			Order(ent.Desc(entvalue.FieldCreatedAt)).
			Limit(1)

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
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func GetCurrencies(ctx context.Context, conds *npool.Conds) ([]*npool.Currency, error) {
	var infos []*npool.Currency

	ids := []string{}
	if conds.CoinTypeID == nil {
		ids = append(ids, conds.GetCoinTypeIDs().GetValue()...)
	} else {
		ids = append(ids, conds.GetCoinTypeID().GetValue())
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, id := range ids {
			var linfos []*npool.Currency

			stm, err := crud.SetQueryConds(&valuemgrpb.Conds{
				ID: conds.ID,
				CoinTypeID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: id,
				},
			}, cli)
			if err != nil {
				return err
			}

			stm.
				Order(ent.Asc(entvalue.FieldCreatedAt)).
				Limit(1)

			if err := join(stm).Scan(_ctx, &linfos); err != nil {
				return err
			}

			infos = append(infos, linfos...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func GetHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Currency, uint32, error) {
	var infos []*npool.Currency
	var total uint32

	ids := []uuid.UUID{}
	if conds.CoinTypeID == nil {
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&valuemgrpb.Conds{
			ID:         conds.ID,
			CoinTypeID: conds.CoinTypeID,
			StartAt:    conds.StartAt,
			EndAt:      conds.EndAt,
		}, cli)
		if err != nil {
			return err
		}

		if len(ids) > 0 {
			stm.Where(
				entvalue.CoinTypeIDIn(ids...),
			)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.CurrencyValueQuery) *ent.CurrencyValueSelect {
	return stm.
		Select(
			entvalue.FieldID,
			entvalue.FieldCoinTypeID,
			entvalue.FieldMarketValueHigh,
			entvalue.FieldMarketValueLow,
			entvalue.FieldCreatedAt,
			entvalue.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entcoinbase.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entvalue.FieldCoinTypeID),
					t1.C(entcoinbase.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entcoinbase.FieldName), "coin_name"),
					sql.As(t1.C(entcoinbase.FieldLogo), "coin_logo"),
					sql.As(t1.C(entcoinbase.FieldUnit), "coin_unit"),
					sql.As(t1.C(entcoinbase.FieldEnv), "coin_env"),
				)

			t2 := sql.Table(entfeed.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(entvalue.FieldFeedSourceID),
					t2.C(entfeed.FieldID),
				).
				AppendSelect(
					sql.As(t2.C(entfeed.FieldFeedType), "feed_type"),
					sql.As(t2.C(entfeed.FieldFeedSource), "feed_source"),
				)
		})
}
