package currency

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"

	coincurrencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency/feed"
	coinfiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/fiat"
	coinfiatcurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/fiat/currency"
	fiatcurrencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency/feed"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coincurrencyfeedmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"
	coinfiatmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"
	fiatcurrencyfeedmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"
)

//nolint:funlen,gocyclo
func _refreshCoinFiats(ctx context.Context, feedType basetypes.CurrencyFeedType) error {
	offset := int32(0)
	const limit = int32(100)

	for {
		h1, err := coinfiat1.NewHandler(
			ctx,
			coinfiat1.WithConds(&coinfiatmwpb.Conds{}),
			coinfiat1.WithOffset(offset),
			coinfiat1.WithLimit(limit),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		infos, _, err := h1.GetCoinFiats(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}
		if len(infos) == 0 {
			return nil
		}

		ids := []string{}
		for _, info := range infos {
			ids = append(ids, info.CoinTypeID)
		}

		h2, err := coincurrencyfeed1.NewHandler(
			ctx,
			coincurrencyfeed1.WithConds(&coincurrencyfeedmwpb.Conds{
				FeedType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(feedType)},
				Disabled:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
				CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
			}),
			coincurrencyfeed1.WithOffset(0),
			coincurrencyfeed1.WithLimit(int32(len(ids))),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		coinFeeds, _, err := h2.GetFeeds(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		coinFeedMap := map[string]*coincurrencyfeedmwpb.Feed{}
		for _, _feed := range coinFeeds {
			coinFeedMap[_feed.CoinTypeID] = _feed
		}

		ids = []string{}
		for _, info := range infos {
			ids = append(ids, info.FiatID)
		}

		h3, err := fiatcurrencyfeed1.NewHandler(
			ctx,
			fiatcurrencyfeed1.WithConds(&fiatcurrencyfeedmwpb.Conds{
				FeedType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(feedType)},
				Disabled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
				FiatIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
			}),
			fiatcurrencyfeed1.WithOffset(0),
			fiatcurrencyfeed1.WithLimit(int32(len(ids))),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		fiatFeeds, _, err := h3.GetFeeds(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		fiatFeedMap := map[string]*fiatcurrencyfeedmwpb.Feed{}
		for _, _feed := range fiatFeeds {
			fiatFeedMap[_feed.FiatID] = _feed
		}

		coinNames := []string{}
		fiatNames := []string{}

		for _, info := range infos {
			_coinFeed, ok := coinFeedMap[info.CoinTypeID]
			if !ok {
				continue
			}
			_fiatFeed, ok := fiatFeedMap[info.FiatID]
			if !ok {
				continue
			}

			coinNames = append(coinNames, _coinFeed.FeedCoinName)
			fiatNames = append(fiatNames, _fiatFeed.FeedFiatName)
		}

		var prices map[string]map[string]decimal.Decimal
		switch feedType {
		case basetypes.CurrencyFeedType_CoinGecko:
			prices, err = coingecko.CoinGeckoPrices(coinNames, fiatNames)
		case basetypes.CurrencyFeedType_CoinBase:
			fallthrough //nolint
		default:
			return fmt.Errorf("invalid feedtype")
		}
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoinFiats",
				"Error", err,
			)
			return err
		}

		_coinFeedMap := map[string][]*coincurrencyfeedmwpb.Feed{}
		for _, _feed := range coinFeeds {
			_coinFeedMap[_feed.FeedCoinName] = append(_coinFeedMap[_feed.FeedCoinName], _feed)
		}
		_fiatFeedMap := map[string][]*fiatcurrencyfeedmwpb.Feed{}
		for _, _feed := range fiatFeeds {
			_fiatFeedMap[_feed.FeedFiatName] = append(_fiatFeedMap[_feed.FeedFiatName], _feed)
		}

		for _coinName, _val1 := range prices {
			_coinFeeds, ok := _coinFeedMap[_coinName]
			if !ok {
				continue
			}
			for _fiatName, _val2 := range _val1 {
				_fiatFeeds, ok := _fiatFeedMap[_fiatName]
				if !ok {
					continue
				}

				price := _val2.String()
				for _, _coinFeed := range _coinFeeds {
					for _, _fiatFeed := range _fiatFeeds {
						h4, err := coinfiatcurrency1.NewHandler(
							ctx,
							coinfiatcurrency1.WithCoinTypeID(&_coinFeed.CoinTypeID),
							coinfiatcurrency1.WithFiatID(&_fiatFeed.FiatID),
							coinfiatcurrency1.WithMarketValueHigh(&price),
							coinfiatcurrency1.WithMarketValueLow(&price),
						)
						if err != nil {
							logger.Sugar().Errorw(
								"_refreshCoinFiats",
								"Error", err,
							)
							return err
						}

						_, err = h4.CreateCurrency(ctx)
						if err != nil {
							logger.Sugar().Errorw(
								"_refreshCoinFiats",
								"Error", err,
							)
							return err
						}
					}
				}
			}
		}

		offset += limit
	}
}

func refreshCoinFiats(ctx context.Context) {
	if err := _refreshCoinFiats(ctx, basetypes.CurrencyFeedType_CoinGecko); err != nil {
		logger.Sugar().Errorw(
			"refreshCoinFiats",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko,
			"Error", err,
		)
	}
}
