package coinbase

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
	coinAPI = "https://api.coinbase.com/v2/prices/COIN-USD/sell"
	fiatAPI = "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	timeout = 5
)

type coinData struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type coinResp struct {
	Data coinData `json:"data"`
}

func CoinBaseUSDPrice(coinName string) (decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	url := strings.ReplaceAll(coinAPI, "COIN", coinName)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinBaseUSDPrice",
			"URL", url,
			"CoinName", coinName,
			"Proxy", socksProxy,
			"error", err,
		)
		return decimal.Decimal{}, err
	}
	r := coinResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinBaseUSDPrice",
			"URL", url,
			"CoinName", coinName,
			"Proxy", socksProxy,
			"Resp", string(resp.Body()),
			"error", err,
		)
		return decimal.Decimal{}, err
	}

	if coinName != r.Data.Base {
		logger.Sugar().Errorw(
			"CoinBaseUSDPrice",
			"URL", url,
			"CoinName", coinName,
			"Proxy", socksProxy,
			"Resp", string(resp.Body()),
			"error", err,
		)
		return decimal.Decimal{}, fmt.Errorf("invalid coin currency")
	}

	amount, err := decimal.NewFromString(r.Data.Amount)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinBaseUSDPrice",
			"URL", url,
			"CoinName", coinName,
			"Proxy", socksProxy,
			"Resp", string(resp.Body()),
			"error", err,
		)
		return decimal.Decimal{}, err
	}

	return amount, nil
}

func CoinBaseUSDPrices(coinNames []string) (map[string]decimal.Decimal, error) {
	prices := map[string]decimal.Decimal{}

	for _, name := range coinNames {
		price, err := CoinBaseUSDPrice(name)
		if err != nil {
			logger.Sugar().Errorw(
				"CoinBaseUSDPrices",
				"CoinName", name,
				"error", err,
			)
			return nil, err
		}
		prices[name] = price

		time.Sleep(500 * time.Millisecond) //nolint
	}

	if len(prices) == 0 {
		return nil, fmt.Errorf("invalid coins")
	}

	return prices, nil
}

type fiatData struct {
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

type fiatResp struct {
	Data fiatData `json:"data"`
}

func CoinBaseFiatPrices(fiatNames []string) (map[string]decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(fiatAPI)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinBaseFiatCurrency",
			"URL", fiatAPI,
			"Proxy", socksProxy,
			"error", err,
		)
		return nil, err
	}
	r := fiatResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinBaseFiatCurrency",
			"URL", fiatAPI,
			"Proxy", socksProxy,
			"Resp", string(resp.Body()),
			"error", err,
		)
		return nil, err
	}

	respMap := map[string]decimal.Decimal{}
	for k, v := range r.Data.Rates {
		c, err := decimal.NewFromString(v)
		if err != nil {
			logger.Sugar().Errorw(
				"CoinBaseFiatCurrency",
				"URL", fiatAPI,
				"Proxy", socksProxy,
				"Resp", string(resp.Body()),
				"error", err,
			)
			return nil, err
		}
		respMap[k] = c
	}
	return respMap, nil
}
