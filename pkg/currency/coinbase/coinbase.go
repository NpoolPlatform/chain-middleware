package currencyvalue

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/go-resty/resty/v2"
)

func coinNameMap(coinName string) string {
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
		return coin
	}
	return ""
}

const (
	coinbaseAPI = "https://api.coinbase.com/v2/prices/COIN-USD/sell"
	timeout     = 5
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
	coin := coinNameMap(strings.ToLower(coinName))

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	url := strings.ReplaceAll(coinbaseAPI, "COIN", coin)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		return decimal.Decimal{}, err
	}
	r := apiResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return decimal.Decimal{}, err
	}

	if coin != r.Data.Base {
		return decimal.Decimal{}, fmt.Errorf("invalid coin currency %v: %v", url, string(resp.Body()))
	}

	amount, err := decimal.NewFromString(r.Data.Amount)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return amount, nil
}

func CoinBaseUSDPrices(coinNames []string) (map[string]decimal.Decimal, error) {
	prices := map[string]decimal.Decimal{}

	for _, name := range coinNames {
		price, err := CoinBaseUSDPrice(name)
		if err != nil {
			return nil, err
		}
		prices[name] = price
	}

	return prices, nil
}