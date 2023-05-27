package currency

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/currency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/currency/coingecko"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

func CoinUSDPrices(ctx context.Context, coinNames []string) (map[string]decimal.Decimal, basetypes.CurrencyFeedType, error) {
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
		return prices, basetypes.CurrencyFeedType_CoinGecko, nil
	}

	logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinGecko", "error", err)

	prices1, err = coinbase.CoinBaseUSDPrices(names)
	if err != nil {
		logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinBase", "error", err)
		return nil, basetypes.CurrencyFeedType_DefaultFeedType, err
	}

	for name, price := range prices1 {
		prices[name] = price
	}

	return prices, basetypes.CurrencyFeedType_CoinBase, nil
}

func CoinCurrencyPrice(ctx context.Context, coinName, currency string) (decimal.Decimal, basetypes.CurrencyFeedType, error) {
	return decimal.Decimal{}, basetypes.CurrencyFeedType_DefaultFeedType, nil
}

func USDPrices(fiatCurrencyNames []string) (map[string]decimal.Decimal, basetypes.CurrencyFeedType, error) {
	prices := map[string]decimal.Decimal{}

	prices1, err := coingecko.UsdFiatCurrency(fiatCurrencyNames)
	if err == nil {
		for name, price := range prices1 {
			prices[strings.ToUpper(name)] = price
		}
		return prices, basetypes.CurrencyFeedType_CoinGecko, nil
	}

	logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinGecko", "error", err)

	prices1, err = coinbase.UsdFaitCurrency()
	if err != nil {
		logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinBase", "error", err)
		return nil, basetypes.CurrencyFeedType_DefaultFeedType, err
	}

	for name, price := range prices1 {
		prices[name] = price
	}

	return prices, basetypes.CurrencyFeedType_CoinBase, nil
}
