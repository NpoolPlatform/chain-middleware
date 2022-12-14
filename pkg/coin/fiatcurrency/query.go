package fiatcurrency

import (
	"context"
	"fmt"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	"time"

	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entfiatcurrency "github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrency"

	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/fiatcurrency"

	constuuid "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

func GetCoinFiatCurrency(ctx context.Context, coinTypeID string) (*npool.FiatCurrency, error) {
	var infos []*npool.FiatCurrency

	coin, err := coin1.GetCoin(ctx, coinTypeID)
	if err != nil {
		return nil, err
	}
	if coin.StableUSD {
		now := uint32(time.Now().Unix())

		return &npool.FiatCurrency{
			ID:                 constuuid.InvalidUUIDStr,
			FiatCurrencyTypeID: "",
			FiatCurrencyName:   "",
			CreatedAt:          now,
			UpdatedAt:          now,
			MarketValueHigh:    "1",
			MarketValueLow:     "1",
			FeedTypeStr:        currencymgrpb.FeedType_StableUSDHardCode.String(),
			FeedType:           currencymgrpb.FeedType_StableUSDHardCode,
			CoinTypeID:         coinTypeID,
			CoinName:           coin.Name,
			CoinLogo:           coin.Logo,
			CoinUnit:           coin.Unit,
			CoinENV:            coin.ENV,
		}, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			FiatCurrency.
			Query().
			Where(
				entfiatcurrency.FiatTypeID(uuid.MustParse(coinTypeID)),
			).
			Order(ent.Desc(entfiatcurrency.FieldCreatedAt)).
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

	infos = expand(infos)

	return infos[0], nil
}

func GetFiatCurrency(ctx context.Context, id string) (*npool.FiatCurrency, error) {
	var infos []*npool.FiatCurrency

	coin, err := coin1.GetCoin(ctx, id)
	if err != nil {
		return nil, err
	}
	if coin.StableUSD {
		now := uint32(time.Now().Unix())

		return &npool.FiatCurrency{
			ID:                 constuuid.InvalidUUIDStr,
			FiatCurrencyTypeID: "",
			FiatCurrencyName:   "",
			CreatedAt:          now,
			UpdatedAt:          now,
			MarketValueHigh:    "1",
			MarketValueLow:     "1",
			FeedTypeStr:        currencymgrpb.FeedType_StableUSDHardCode.String(),
			FeedType:           currencymgrpb.FeedType_StableUSDHardCode,
			CoinTypeID:         "",
			CoinName:           coin.Name,
			CoinLogo:           coin.Logo,
			CoinUnit:           coin.Unit,
			CoinENV:            coin.ENV,
		}, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			FiatCurrency.
			Query().
			Where(
				entfiatcurrency.FiatTypeID(uuid.MustParse(id)),
			).
			Order(ent.Desc(entfiatcurrency.FieldCreatedAt)).
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

	infos = expand(infos)

	return infos[0], nil
}

func GetFiatCurrencies(ctx context.Context, conds *npool.Conds) ([]*npool.FiatCurrency, error) {
	var infos []*npool.FiatCurrency

	ids := []string{}
	if conds.CoinTypeID == nil {
		ids = append(ids, conds.GetCoinTypeIDs().GetValue()...)
	} else {
		ids = append(ids, conds.GetCoinTypeID().GetValue())
	}

	coins, err := coin1.GetManyCoins(ctx, ids)
	if err != nil {
		return nil, err
	}

	ids = []string{}
	for _, coin := range coins {
		if coin.StableUSD {
			continue
		}
		ids = append(ids, coin.ID)
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, id := range ids {
			var linfos []*npool.FiatCurrency

			stm, err := crud.SetQueryConds(&fiatcurrencymgrpb.Conds{
				ID: conds.ID,
				FiatTypeID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: id,
				},
			}, cli)
			if err != nil {
				return err
			}

			stm.
				Order(ent.Asc(entfiatcurrency.FieldCreatedAt)).
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

	infos = expand(infos)

	for _, coin := range coins {
		if !coin.StableUSD {
			continue
		}

		now := uint32(time.Now().Unix())

		infos = append(infos, &npool.FiatCurrency{
			ID:              constuuid.InvalidUUIDStr,
			CoinTypeID:      coin.ID,
			CoinName:        coin.Name,
			CoinLogo:        coin.Logo,
			CoinUnit:        coin.Unit,
			CoinENV:         coin.ENV,
			CreatedAt:       now,
			UpdatedAt:       now,
			MarketValueHigh: "1",
			MarketValueLow:  "1",
			FeedTypeStr:     currencymgrpb.FeedType_StableUSDHardCode.String(),
			FeedType:        currencymgrpb.FeedType_StableUSDHardCode,
		})
	}

	return infos, nil
}

func GetHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FiatCurrency, uint32, error) {
	var infos []*npool.FiatCurrency
	var total uint32

	ids := []uuid.UUID{}
	if conds.CoinTypeID == nil {
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(&fiatcurrencymgrpb.Conds{
			ID:         conds.ID,
			FiatTypeID: conds.CoinTypeID,
			StartAt:    conds.StartAt,
			EndAt:      conds.EndAt,
		}, cli)
		if err != nil {
			return err
		}

		if len(ids) > 0 {
			stm.Where(
				entfiatcurrency.FiatTypeIDIn(ids...),
			)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm.
			Order(ent.Desc(entfiatcurrency.FieldCreatedAt)).
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

func join(stm *ent.FiatCurrencyQuery) *ent.FiatCurrencySelect {
	return stm.
		Select(
			entfiatcurrency.FieldID,
			entfiatcurrency.FieldFiatTypeID,
			entfiatcurrency.FieldFeedType,
			entfiatcurrency.FieldMarketValueHigh,
			entfiatcurrency.FieldMarketValueLow,
			entfiatcurrency.FieldCreatedAt,
			entfiatcurrency.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entcoinbase.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entfiatcurrency.FieldFiatTypeID),
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

func expand(infos []*npool.FiatCurrency) []*npool.FiatCurrency {
	for _, info := range infos {
		info.FeedType = currencymgrpb.FeedType(currencymgrpb.FeedType_value[info.FeedTypeStr])
	}
	return infos
}
