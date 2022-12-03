package currency

import (
	"context"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/currency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	"github.com/shopspring/decimal"
)

func stableUSD(coinName string) bool {
	priceMap := map[string]string{
		"tusdt":       "tether",
		"usdt":        "tether",
		"tusdterc20":  "tether",
		"usdterc20":   "tether",
		"tusdttrc20":  "tether",
		"usdttrc20":   "tether",
		"tbinanceusd": "binance-usd",
		"binanceusd":  "binance-usd",
		"tusdcerc20":  "usdc",
		"usdcerc20":   "usdc",
	}
	_, ok := priceMap[coinName]
	return ok
}

func CoinUSDPrices(ctx context.Context, coinNames []string) (map[string]decimal.Decimal, currencymgrpb.FeedType, error) {
	prices := map[string]decimal.Decimal{}
	names := []string{}

	for _, name := range coinNames {
		if stableUSD(name) {
			prices[name] = decimal.NewFromInt(1)
			continue
		}
		names = append(names, name)
	}

	prices1, err := coingecko.CoinGeckoUSDPrices(names)
	if err == nil {
		for name, price := range prices1 {
			prices[name] = price
		}
		return prices, currencymgrpb.FeedType_CoinGecko, nil
	}

	prices1, err = coinbase.CoinBaseUSDPrices(names)
	if err != nil {
		return nil, currencymgrpb.FeedType_DefaultFeedType, err
	}

	for name, price := range prices1 {
		prices[name] = price
	}

	return prices, currencymgrpb.FeedType_CoinBase, nil
}

func CoinCurrencyPrice(ctx context.Context, coinName, currency string) (decimal.Decimal, currencymgrpb.FeedType, error) {
	return decimal.Decimal{}, currencymgrpb.FeedType_DefaultFeedType, nil
}
