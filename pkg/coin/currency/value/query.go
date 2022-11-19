package currencyvalue

import (
	"context"
	"fmt"

	// valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entfeed "github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"
	entvalue "github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyvalue"

	// crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/value"

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

func join(stm *ent.CurrencyValueQuery) *ent.CurrencyValueSelect {
	return stm.
		Select(
			entvalue.FieldID,
			entvalue.FieldCoinTypeID,
			entvalue.FieldFeedSourceID,
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
