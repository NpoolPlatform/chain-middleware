package currencyfeed

import (
	"context"
	"fmt"

	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/feed"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entfeed "github.com/NpoolPlatform/chain-manager/pkg/db/ent/currencyfeed"

	"entgo.io/ent/dialect/sql"

	"github.com/google/uuid"
)

func GetCurrencyFeed(ctx context.Context, id string) (*npool.CurrencyFeed, error) {
	var infos []*npool.CurrencyFeed

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			CurrencyFeed.
			Query().
			Where(
				entfeed.ID(uuid.MustParse(id)),
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

	infos = expand(infos)

	return infos[0], nil
}

func GetCurrencyFeeds(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.CurrencyFeed, uint32, error) {
	var infos []*npool.CurrencyFeed
	var total uint32

	ids := []uuid.UUID{}
	for _, id := range conds.GetCoinTypeIDs().GetValue() {
		ids = append(ids, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&feedmgrpb.Conds{
			ID:         conds.ID,
			CoinTypeID: conds.CoinTypeID,
			Disabled:   conds.Disabled,
		}, cli)
		if err != nil {
			return err
		}

		if len(ids) > 0 {
			stm.Where(
				entfeed.CoinTypeIDIn(ids...),
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

	infos = expand(infos)

	return infos, total, nil
}

func join(stm *ent.CurrencyFeedQuery) *ent.CurrencyFeedSelect {
	return stm.
		Select(
			entfeed.FieldID,
			entfeed.FieldCoinTypeID,
			entfeed.FieldFeedType,
			entfeed.FieldFeedSource,
			entfeed.FieldCreatedAt,
			entfeed.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entcoinbase.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entfeed.FieldCoinTypeID),
					t1.C(entcoinbase.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entcoinbase.FieldName), "coin_name"),
					sql.As(t1.C(entcoinbase.FieldLogo), "coin_logo"),
					sql.As(t1.C(entcoinbase.FieldUnit), "coin_unit"),
					sql.As(t1.C(entcoinbase.FieldEnv), "coin_env"),
				)
		})
}

func expand(infos []*npool.CurrencyFeed) []*npool.CurrencyFeed {
	for _, info := range infos {
		info.FeedType = feedmgrpb.FeedType(feedmgrpb.FeedType_value[info.FeedTypeStr])
	}
	return infos
}
