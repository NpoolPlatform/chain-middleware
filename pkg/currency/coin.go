package currency

import (
	"context"
	"fmt"

	coincurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency"
	coincurrencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency/feed"
	coincurrencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	coincurrencyfeedmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/currency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/shopspring/decimal"
)

//nolint:funlen,gocyclo
func _refreshCoins(ctx context.Context, feedType basetypes.CurrencyFeedType) error {
	offset := int32(0)
	const limit = int32(100)

	for {
		h1, err := coincurrencyfeed1.NewHandler(
			ctx,
			coincurrencyfeed1.WithConds(&coincurrencyfeedmwpb.Conds{
				FeedType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(feedType)},
				Disabled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
			}),
			coincurrencyfeed1.WithOffset(offset),
			coincurrencyfeed1.WithLimit(limit),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoins",
				"Error", err,
			)
			return err
		}

		feeds, _, err := h1.GetFeeds(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoins",
				"Error", err,
			)
			return err
		}
		if len(feeds) == 0 {
			return nil
		}

		feedMap := map[string]*coincurrencyfeedmwpb.Feed{}
		coinNames := []string{}

		for _, _feed := range feeds {
			if _feed.FeedCoinName == "" {
				continue
			}
			feedMap[_feed.CoinTypeID] = _feed
			coinNames = append(coinNames, _feed.FeedCoinName)
		}
		if len(coinNames) == 0 {
			return fmt.Errorf("invalid feeds")
		}

		var prices map[string]decimal.Decimal
		switch feedType {
		case basetypes.CurrencyFeedType_CoinGecko:
			prices, err = coingecko.CoinGeckoUSDPrices(coinNames)
		case basetypes.CurrencyFeedType_CoinBase:
			prices, err = coinbase.CoinBaseUSDPrices(coinNames)
		default:
			return fmt.Errorf("invalid feedtype")
		}
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoins",
				"Error", err,
			)
			return err
		}

		_feedMap := map[string][]*coincurrencyfeedmwpb.Feed{}
		coinRefreshed := map[string]bool{}
		for _, _feed := range feeds {
			_feedMap[_feed.FeedCoinName] = append(_feedMap[_feed.FeedCoinName], _feed)
		}

		reqs := []*coincurrencymwpb.CurrencyReq{}
		for _feedCoinName, _price := range prices {
			_feeds, ok := _feedMap[_feedCoinName]
			if !ok {
				continue
			}
			_priceStr := _price.String()

			for _, _feed := range _feeds {
				reqs = append(reqs, &coincurrencymwpb.CurrencyReq{
					CoinTypeID:      &_feed.CoinTypeID,
					FeedType:        &feedType,
					MarketValueHigh: &_priceStr,
					MarketValueLow:  &_priceStr,
				})
				coinRefreshed[_feed.CoinTypeID] = true
			}
		}

		for _, _feed := range feeds {
			refreshed, ok := coinRefreshed[_feed.CoinTypeID]
			if !ok {
				logger.Sugar().Warnw(
					"_refreshCoins",
					"CoinTypeID", _feed.CoinTypeID,
					"Refreshed", refreshed,
					"FeedCoinName", _feed.FeedCoinName,
				)
			}
		}

		if len(reqs) == 0 {
			continue
		}

		h2, err := coincurrency1.NewHandler(
			ctx,
			coincurrency1.WithReqs(reqs, true),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoins",
				"Error", err,
			)
			return err
		}

		_, err = h2.CreateCurrencies(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshCoins",
				"Error", err,
			)
			return err
		}

		offset += limit
	}
}

func refreshCoins(ctx context.Context) {
	if err := _refreshCoins(ctx, basetypes.CurrencyFeedType_CoinGecko); err != nil {
		logger.Sugar().Warnw(
			"refreshCoins",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko,
			"Error", err,
		)
	}

	if err := _refreshCoins(ctx, basetypes.CurrencyFeedType_CoinBase); err != nil {
		logger.Sugar().Warnw(
			"refreshCoins",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko,
			"Error", err,
		)
	}
}
