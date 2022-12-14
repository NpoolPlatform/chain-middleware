package currency

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	coinbase "github.com/NpoolPlatform/chain-middleware/pkg/fiatcurrency/coinbase"
	coingecko "github.com/NpoolPlatform/chain-middleware/pkg/fiatcurrency/coingecko"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	"github.com/shopspring/decimal"
)

func CoinUSDPrices(fiatCurrencyNames []string) (map[string]decimal.Decimal, currencymgrpb.FeedType, error) {
	prices := map[string]decimal.Decimal{}

	prices1, err := coingecko.UsdFiatCurrency(fiatCurrencyNames)
	if err == nil {
		for name, price := range prices1 {
			prices[name] = price
		}
		return prices, currencymgrpb.FeedType_CoinGecko, nil
	}

	logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinGecko", "error", err)

	prices1, err = coinbase.UsdFaitCurrency()
	if err != nil {
		logger.Sugar().Errorw("CoinUSDPrices", "Feed", "CoinBase", "error", err)
		return nil, currencymgrpb.FeedType_DefaultFeedType, err
	}

	for name, price := range prices1 {
		prices[name] = price
	}

	return prices, currencymgrpb.FeedType_CoinBase, nil
}
