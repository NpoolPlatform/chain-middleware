package currencyvalue

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/shopspring/decimal"

	"github.com/go-resty/resty/v2"
)

const (
	coinGeckoAPI = "https://api.coingecko.com/api/v3"
	timeout      = 5
)

func coinNameMap(coinName string) (string, bool) {
	coinMap := map[string]string{
		"fil":          "filecoin",
		"filecoin":     "filecoin",
		"tfilecoin":    "filecoin",
		"btc":          "bitcoin",
		"bitcoin":      "bitcoin",
		"tbitcoin":     "bitcoin",
		"tethereum":    "ethereum",
		"eth":          "ethereum",
		"ethereum":     "ethereum",
		"tusdt":        "tether",
		"usdt":         "tether",
		"tusdterc20":   "tether",
		"usdterc20":    "tether",
		"tusdttrc20":   "tether",
		"usdttrc20":    "tether",
		"sol":          "solana",
		"solana":       "solana",
		"tsolana":      "solana",
		"tbinancecoin": "binancecoin",
		"binancecoin":  "binancecoin",
		"tbinanceusd":  "binance-usd",
		"binanceusd":   "binance-usd",
		"ttron":        "tron",
		"tron":         "tron",
		"tusdcerc20":   "tusdcerc20",
		"usdcerc20":    "usdcerc20",
	}
	if coin, ok := coinMap[coinName]; ok {
		return coin, true
	}
	return coinName, false
}

func CoinGeckoUSDPrices(coinNames []string) (map[string]decimal.Decimal, error) {
	coins := ""
	coinMap := map[string]string{}

	for _, val := range coinNames {
		coin, ok := coinNameMap(strings.ToLower(val))
		if !ok {
			logger.Sugar().Errorw("CoinGeckoUSDPrices", "Coin", val)
			continue
		}
		if coins != "" {
			coins += ","
		}
		coins += coin
		coinMap[coin] = val
	}

	if coins == "" {
		return nil, fmt.Errorf("invalid coins")
	}

	logger.Sugar().Errorw("CoinGeckoUSDPrices", "Coins", coins)

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	url := fmt.Sprintf("%v%v?ids=%v&vs_currencies=usd", coinGeckoAPI, "/simple/price", coins)
	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		return nil, err
	}
	respMap := map[string]map[string]float64{}
	err = json.Unmarshal(resp.Body(), &respMap)
	if err != nil {
		logger.Sugar().Errorw("CoinGeckoUSDPrices", "error", err)
		return nil, err
	}

	infoMap := map[string]decimal.Decimal{}
	for key, val := range respMap {
		coin, ok := coinMap[key]
		if !ok {
			return nil, fmt.Errorf("invalid coin")
		}
		price := decimal.NewFromInt(0)
		for _, v := range val {
			price = decimal.NewFromFloat(v)
		}
		infoMap[coin] = price
	}

	return infoMap, nil
}
