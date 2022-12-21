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

func coinNameMap(coinName string) (string, bool) {
	coinMap := map[string]string{
		"fil":          "FIL",
		"filecoin":     "FIL",
		"tfilecoin":    "FIL",
		"btc":          "BTC",
		"bitcoin":      "BTC",
		"tbitcoin":     "BTC",
		"tethereum":    "ETH",
		"eth":          "ETH",
		"ethereum":     "ETH",
		"sol":          "SOL",
		"solana":       "SOL",
		"tsolana":      "SOL",
		"tbinancecoin": "binancecoin",
		"binancecoin":  "binancecoin",
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

const (
	coinbaseAPI     = "https://api.coinbase.com/v2/prices/COIN-USD/sell"
	coinbaseAPIFiat = "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	timeout         = 5
)

type apiData struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type apiResp struct {
	Data apiData `json:"data"`
}

func CoinBaseUSDPrice(coinName string) (decimal.Decimal, error) {
	coin, ok := coinNameMap(strings.ToLower(coinName))
	if !ok {
		return decimal.Decimal{}, fmt.Errorf("not supported coin")
	}

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	url := strings.ReplaceAll(coinbaseAPI, "COIN", coin)

	logger.Sugar().Errorw("CoinBaseUSDPrice", "URL", url)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return decimal.Decimal{}, err
	}
	r := apiResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return decimal.Decimal{}, err
	}

	if coin != r.Data.Base {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", "invalid coinbase")
		return decimal.Decimal{}, fmt.Errorf("invalid coin currency %v: %v", url, string(resp.Body()))
	}

	amount, err := decimal.NewFromString(r.Data.Amount)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return decimal.Decimal{}, err
	}

	return amount, nil
}

func CoinBaseUSDPrices(coinNames []string) (map[string]decimal.Decimal, error) {
	prices := map[string]decimal.Decimal{}

	for _, name := range coinNames {
		price, err := CoinBaseUSDPrice(name)
		if err != nil {
			logger.Sugar().Errorw("CoinBaseUSDPrices", "Coin", name, "error", err)
			continue
		}
		prices[name] = price

		time.Sleep(500 * time.Millisecond) //nolint
	}

	if len(prices) == 0 {
		return nil, fmt.Errorf("invalid coins")
	}

	return prices, nil
}

type fiatAPIData struct {
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

type fiatAPIResp struct {
	Data fiatAPIData `json:"data"`
}

func UsdFaitCurrency() (map[string]decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	logger.Sugar().Errorw("CoinBaseUSDPrice", "URL", coinbaseAPIFiat)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(coinbaseAPIFiat)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return nil, err
	}
	r := fiatAPIResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return nil, err
	}

	respMap := map[string]decimal.Decimal{}
	for k, v := range r.Data.Rates {
		c, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		respMap[k] = c
	}
	return respMap, nil
}
