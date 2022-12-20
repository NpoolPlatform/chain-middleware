package fiat

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	currencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entfiatcurrency "github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrency"
	entfiatcurrencytype "github.com/NpoolPlatform/chain-manager/pkg/db/ent/fiatcurrencytype"

	"entgo.io/ent/dialect/sql"
	fiatcurrencycrud "github.com/NpoolPlatform/chain-manager/pkg/crud/fiat/currency"
	"github.com/NpoolPlatform/chain-middleware/pkg/coin/currency"
	fiatcurrencypb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiat/currency"
	"github.com/google/uuid"

	coinpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
)

var DefaultCoinTypes = []string{"usdterc20", "tusdttrc20"}

func GetFiatCurrency(ctx context.Context, fiatTypeID string) (*npool.FiatCurrency, error) {
	fiatCurrencies := []*npool.FiatCurrency{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Debug().
			FiatCurrency.
			Query().
			Where(
				entfiatcurrency.FiatCurrencyTypeID(uuid.MustParse(fiatTypeID)),
			).
			Order(ent.Desc(entfiatcurrency.FieldCreatedAt))
		return join(stm).
			Scan(_ctx, &fiatCurrencies)
	})
	if err != nil {
		return nil, err
	}
	if len(fiatCurrencies) == 0 {
		logger.Sugar().Errorw("fiatCurrencies is empty")
		return nil, nil
	}

	fiatCurrencies = expand(fiatCurrencies)

	coinInfos, _, err := coin.GetCoins(ctx, &coinpb.Conds{
		Names: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: DefaultCoinTypes,
		},
	}, 0, 1)
	if err != nil {
		return nil, err
	}

	if len(coinInfos) == 0 {
		return nil, fmt.Errorf("coin info not found")
	}
	coinCurrency, err := currency.GetCoinCurrency(ctx, coinInfos[0].ID)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	fiatCurrency := fiatCurrencies[0]

	return &npool.FiatCurrency{
		ID:                 fiatCurrency.ID,
		FiatCurrencyTypeID: fiatCurrency.FiatCurrencyTypeID,
		FeedTypeStr:        fiatCurrency.FeedTypeStr,
		FeedType:           fiatCurrency.FeedType,
		FiatCurrencyName:   fiatCurrency.FiatCurrencyName,
		FiatCurrencyLogo:   fiatCurrency.FiatCurrencyLogo,
		MarketValueHigh:    fiatCurrency.MarketValueHigh,
		MarketValueLow:     fiatCurrency.MarketValueLow,
		CreatedAt:          fiatCurrency.CreatedAt,
		UpdatedAt:          fiatCurrency.UpdatedAt,
		CoinTypeID:         coinCurrency.CoinTypeID,
		CoinName:           coinCurrency.CoinName,
		CoinLogo:           coinCurrency.CoinLogo,
		CoinUnit:           coinCurrency.CoinUnit,
		CoinENV:            coinCurrency.CoinENV,
	}, nil
}

func GetFiatCurrencies(ctx context.Context, conds *npool.Conds) ([]*npool.FiatCurrency, error) {
	fiatCurrencies := []*npool.FiatCurrency{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := fiatcurrencycrud.SetQueryConds(&fiatcurrencypb.Conds{
			ID:                  conds.ID,
			FiatCurrencyTypeID:  conds.FiatCurrencyTypeID,
			FiatCurrencyTypeIDs: conds.FiatCurrencyTypeIDs,
			StartAt:             conds.StartAt,
			EndAt:               conds.EndAt,
		}, cli)
		if err != nil {
			return err
		}
		stm.Order(ent.Desc(entfiatcurrency.FieldCreatedAt))
		return join(stm).
			Scan(_ctx, &fiatCurrencies)
	})
	if err != nil {
		return nil, err
	}
	if len(fiatCurrencies) == 0 {
		logger.Sugar().Errorw("fiatCurrencies is empty")
		return nil, nil
	}

	fiatCurrencies = expand(fiatCurrencies)

	return fiatCurrencies, nil
}

func GetCoinFiatCurrencies(ctx context.Context, coinTypeIDs, fiatTypeIDs []string) ([]*npool.FiatCurrency, error) {
	fiatCurrencies := []*npool.FiatCurrency{}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, id := range fiatTypeIDs {
			var linfos []*npool.FiatCurrency
			stm, err := fiatcurrencycrud.SetQueryConds(&fiatcurrencypb.Conds{
				FiatCurrencyTypeID: &commonpb.StringVal{
					Op:    cruder.EQ,
					Value: id,
				},
			}, cli)
			if err != nil {
				return err
			}
			stm.
				Order(ent.Desc(entfiatcurrency.FieldCreatedAt)).
				Limit(1)
			if err := join(stm).Scan(_ctx, &linfos); err != nil {
				return err
			}
			fiatCurrencies = append(fiatCurrencies, linfos...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(fiatCurrencies) == 0 {
		logger.Sugar().Errorw("fiatCurrencies is empty")
		return nil, nil
	}

	fiatCurrencies = expand(fiatCurrencies)

	coinCurrencies, err := currency.GetCurrencies(ctx, &currencymwpb.Conds{
		CoinTypeIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: coinTypeIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(coinCurrencies) == 0 {
		logger.Sugar().Errorw("coinCurrencies is empty")
		return nil, nil
	}
	infos := []*npool.FiatCurrency{}

	for _, coinCurrency := range coinCurrencies {
		for _, fiatCurrency := range fiatCurrencies {
			marketValueHigh, err := decimal.NewFromString(fiatCurrency.MarketValueHigh)
			if err != nil {
				return nil, err
			}
			marketValueHigh1, err := decimal.NewFromString(coinCurrency.MarketValueHigh)
			if err != nil {
				return nil, err
			}
			marketValueLow, err := decimal.NewFromString(fiatCurrency.MarketValueLow)
			if err != nil {
				return nil, err
			}
			marketValueLow1, err := decimal.NewFromString(coinCurrency.MarketValueLow)
			if err != nil {
				return nil, err
			}
			infos = append(infos, &npool.FiatCurrency{
				ID:                 fiatCurrency.ID,
				FiatCurrencyTypeID: fiatCurrency.FiatCurrencyTypeID,
				FeedTypeStr:        fiatCurrency.FeedTypeStr,
				FeedType:           fiatCurrency.FeedType,
				FiatCurrencyName:   fiatCurrency.FiatCurrencyName,
				FiatCurrencyLogo:   fiatCurrency.FiatCurrencyLogo,
				MarketValueHigh:    marketValueHigh.Mul(marketValueHigh1).String(),
				MarketValueLow:     marketValueLow.Mul(marketValueLow1).String(),
				CreatedAt:          fiatCurrency.CreatedAt,
				UpdatedAt:          fiatCurrency.UpdatedAt,
				CoinTypeID:         coinCurrency.CoinTypeID,
				CoinName:           coinCurrency.CoinName,
				CoinLogo:           coinCurrency.CoinLogo,
				CoinUnit:           coinCurrency.CoinUnit,
				CoinENV:            coinCurrency.CoinENV,
			})
		}
	}

	return infos, nil
}

func GetHistories(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.FiatCurrency, uint32, error) {
	var infos []*npool.FiatCurrency
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := fiatcurrencycrud.SetQueryConds(&fiatcurrencypb.Conds{
			ID:                  conds.ID,
			FiatCurrencyTypeID:  conds.FiatCurrencyTypeID,
			FiatCurrencyTypeIDs: conds.FiatCurrencyTypeIDs,
			StartAt:             conds.StartAt,
			EndAt:               conds.EndAt,
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
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entfiatcurrencytype.Table)
			s.Select(
				s.C(entfiatcurrency.FieldID),
				s.C(entfiatcurrency.FieldFiatCurrencyTypeID),
				s.C(entfiatcurrency.FieldFeedType),
				s.C(entfiatcurrency.FieldMarketValueHigh),
				s.C(entfiatcurrency.FieldMarketValueLow),
				s.C(entfiatcurrency.FieldCreatedAt),
				s.C(entfiatcurrency.FieldUpdatedAt),
			)
			s.
				LeftJoin(t1).
				On(
					s.C(entfiatcurrency.FieldFiatCurrencyTypeID),
					t1.C(entfiatcurrencytype.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entfiatcurrencytype.FieldName), "fiat_currency_name"),
					sql.As(t1.C(entfiatcurrencytype.FieldLogo), "fiat_currency_logo"),
				)
		})
}

func expand(infos []*npool.FiatCurrency) []*npool.FiatCurrency {
	for _, info := range infos {
		info.FeedType = currencymgrpb.FeedType(currencymgrpb.FeedType_value[info.FeedTypeStr])
	}
	return infos
}
