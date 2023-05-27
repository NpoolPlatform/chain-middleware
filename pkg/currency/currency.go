package currency

import (
	"context"
	"fmt"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	coincurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency"
	coincurrencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency/feed"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
	coincurrencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
	coincurrencyfeedmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/currency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/shopspring/decimal"
)

func _refreshCoins(ctx context.Context, coins []*coinmwpb.Coin, feedType basetypes.CurrencyFeedType) error {
	ids := []string{}
	for _, _coin := range coins {
		ids = append(ids, _coin.ID)
	}

	h1, err := coincurrencyfeed1.NewHandler(
		ctx,
		coincurrencyfeed1.WithConds(&coincurrencyfeedmwpb.Conds{
			CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
			FeedType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(feedType)},
			Disabled:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		}),
		coincurrencyfeed1.WithOffset(0),
		coincurrencyfeed1.WithLimit(int32(len(ids))),
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

	feedMap := map[string]*coincurrencyfeedmwpb.Feed{}
	coinNames := []string{}

	for _, _feed := range feeds {
		feedMap[_feed.CoinTypeID] = _feed
		coinNames = append(coinNames, _feed.FeedCoinName)
	}
	if len(coinNames) == 0 {
		return fmt.Errorf("invalid feeds")
	}

	prices := map[string]decimal.Decimal{}

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

	coinMap := map[string]*coinmwpb.Coin{}
	_coinMap := map[string]*coinmwpb.Coin{}
	coinRefreshed := map[string]bool{}

	for _, _coin := range coins {
		coinMap[_coin.ID] = _coin
		_coinMap[_coin.Name] = _coin
	}
	for _, _feed := range feeds {
		_coinMap[_feed.FeedCoinName] = coinMap[_feed.CoinTypeID]
	}

	reqs := []*coincurrencymwpb.CurrencyReq{}
	for _coinName, _price := range prices {
		_coin, ok := _coinMap[_coinName]
		if !ok {
			continue
		}
		_priceStr := _price.String()
		reqs = append(reqs, &coincurrencymwpb.CurrencyReq{
			CoinTypeID:      &_coin.ID,
			FeedType:        &feedType,
			MarketValueHigh: &_priceStr,
			MarketValueLow:  &_priceStr,
		})
		coinRefreshed[_coin.ID] = true
	}

	for _, _coin := range coins {
		refreshed, ok := coinRefreshed[_coin.ID]
		if !ok {
			_feed := feedMap[_coin.ID]
			logger.Sugar().Warnw(
				"_refreshCoins",
				"CoinName", _coin.Name,
				"CoinTypeID", _coin.ID,
				"Refreshed", refreshed,
				"Feed", _feed,
			)
		}
	}

	h2, err := coincurrency1.NewHandler(
		ctx,
		coincurrency1.WithReqs(reqs),
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

	return nil
}

func refreshCoins(ctx context.Context) error {
	offset := int32(0)
	limit := int32(100)

	for {
		h1, err := coin1.NewHandler(
			ctx,
			coin1.WithConds(&coinmwpb.Conds{}),
			coin1.WithOffset(offset),
			coin1.WithLimit(limit),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"refreshCoins",
				"Error", err,
			)
			return err
		}

		coins, _, err := h1.GetCoins(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"refreshCoins",
				"Error", err,
			)
			return err
		}
		if len(coins) == 0 {
			return nil
		}

		offset += limit

		err = _refreshCoins(ctx, coins, basetypes.CurrencyFeedType_CoinGecko)
		if err == nil {
			continue
		}
		logger.Sugar().Warnw(
			"refreshCoins",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko.String(),
			"Error", err,
		)

		err = _refreshCoins(ctx, coins, basetypes.CurrencyFeedType_CoinBase)
		if err == nil {
			continue
		}
		logger.Sugar().Warnw(
			"refreshCoins",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko.String(),
			"Error", err,
		)
	}

	return nil
}

func refreshFiats(ctx context.Context) error {
	return nil
}
