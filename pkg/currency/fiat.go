package currency

import (
	"context"
	"fmt"

	fiatcurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency"
	fiatcurrencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency/feed"
	fiatcurrencymwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"
	fiatcurrencyfeedmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/currency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/shopspring/decimal"
)

func _refreshFiats(ctx context.Context, feedType basetypes.CurrencyFeedType) error {
	offset := int32(0)
	limit := int32(100)

	for {
		h1, err := fiatcurrencyfeed1.NewHandler(
			ctx,
			fiatcurrencyfeed1.WithConds(&fiatcurrencyfeedmwpb.Conds{
				FeedType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(feedType)},
				Disabled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
			}),
			fiatcurrencyfeed1.WithOffset(offset),
			fiatcurrencyfeed1.WithLimit(limit),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshFiats",
				"Error", err,
			)
			return err
		}

		feeds, _, err := h1.GetFeeds(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshFiats",
				"Error", err,
			)
			return err
		}
		if len(feeds) == 0 {
			return nil
		}

		feedMap := map[string]*fiatcurrencyfeedmwpb.Feed{}
		fiatNames := []string{}

		for _, _feed := range feeds {
			feedMap[_feed.FiatID] = _feed
			fiatNames = append(fiatNames, _feed.FeedFiatName)
		}
		if len(fiatNames) == 0 {
			return fmt.Errorf("invalid feeds")
		}

		prices := map[string]decimal.Decimal{}

		switch feedType {
		case basetypes.CurrencyFeedType_CoinGecko:
			prices, err = coingecko.CoinGeckoUSDPrices(fiatNames)
		case basetypes.CurrencyFeedType_CoinBase:
			prices, err = coinbase.CoinBaseUSDPrices(fiatNames)
		default:
			return fmt.Errorf("invalid feedtype")
		}
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshFiats",
				"Error", err,
			)
			return err
		}

		_feedMap := map[string]*fiatcurrencyfeedmwpb.Feed{}
		fiatRefreshed := map[string]bool{}
		for _, _feed := range feeds {
			_feedMap[_feed.FeedFiatName] = _feed
		}

		reqs := []*fiatcurrencymwpb.CurrencyReq{}
		for _feedFiatName, _price := range prices {
			_feed, ok := _feedMap[_feedFiatName]
			if !ok {
				continue
			}
			_priceStr := _price.String()
			reqs = append(reqs, &fiatcurrencymwpb.CurrencyReq{
				FiatID:          &_feed.FiatID,
				FeedType:        &feedType,
				MarketValueHigh: &_priceStr,
				MarketValueLow:  &_priceStr,
			})
			fiatRefreshed[_feed.FiatID] = true
		}

		for _, _feed := range feeds {
			refreshed, ok := fiatRefreshed[_feed.FiatID]
			if !ok {
				logger.Sugar().Warnw(
					"_refreshFiats",
					"FiatID", _feed.FiatID,
					"Refreshed", refreshed,
					"FeedFiatName", _feed.FeedFiatName,
				)
			}
		}

		h2, err := fiatcurrency1.NewHandler(
			ctx,
			fiatcurrency1.WithReqs(reqs),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshFiats",
				"Error", err,
			)
			return err
		}

		_, err = h2.CreateCurrencies(ctx)
		if err != nil {
			logger.Sugar().Errorw(
				"_refreshFiats",
				"Error", err,
			)
			return err
		}
	}

	return nil
}

func refreshFiats(ctx context.Context) {
	if err := _refreshFiats(ctx, basetypes.CurrencyFeedType_CoinGecko); err != nil {
		logger.Sugar().Warnw(
			"refreshFiats",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko.String(),
			"Error", err,
		)
	}

	if err := _refreshFiats(ctx, basetypes.CurrencyFeedType_CoinBase); err != nil {
		logger.Sugar().Warnw(
			"refreshFiats",
			"FeedType", basetypes.CurrencyFeedType_CoinGecko.String(),
			"Error", err,
		)
	}
}
