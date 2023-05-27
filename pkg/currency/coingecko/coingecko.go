package coingecko

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
	coinAPI = "https://api.coingecko.com/api/v3"
	timeout = 5
)

func CoinGeckoUSDPrices(coinNames []string) (map[string]decimal.Decimal, error) {
	if len(coinNames) == 0 {
		return nil, fmt.Errorf("invalid coinnames")
	}

	coins := strings.Join(coinNames, ",")

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	url := fmt.Sprintf("%v%v?ids=%v&vs_currencies=usd", coinAPI, "/simple/price", coins)
	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinGeckoUSDPrices",
			"URL", url,
			"Proxy", socksProxy,
			"Error", err,
		)
		return nil, err
	}
	respMap := map[string]map[string]float64{}
	err = json.Unmarshal(resp.Body(), &respMap)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinGeckoUSDPrices",
			"URL", url,
			"Proxy", socksProxy,
			"Resp", string(resp.Body()),
			"Error", err,
		)
		return nil, err
	}

	infoMap := map[string]decimal.Decimal{}
	for key, val := range respMap {
		price := decimal.NewFromInt(0)
		for _, v := range val {
			price = decimal.NewFromFloat(v)
		}
		infoMap[key] = price
	}

	return infoMap, nil
}

func CoinGeckoFiatPrices(fiatNames []string) (map[string]decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	fiats := strings.Join(fiatNames, ",")
	url := fmt.Sprintf("%v%v?ids=usd&vs_currencies=%v", coinAPI, "/simple/price", fiats)
	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}
	resp, err := cli.R().Get(url)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinGeckoFiatPrices",
			"URL", url,
			"Proxy", socksProxy,
			"Error", err,
		)
		return nil, err
	}

	respMap := map[string]map[string]float64{}
	err = json.Unmarshal(resp.Body(), &respMap)
	if err != nil {
		logger.Sugar().Errorw(
			"CoinGeckoFiatPrices",
			"URL", url,
			"Proxy", socksProxy,
			"Resp", resp.Body(),
			"Error", err,
		)
		return nil, err
	}

	respMap1 := map[string]decimal.Decimal{}
	if _, ok := respMap["usd"]; ok {
		for key, val := range respMap["usd"] {
			c := decimal.NewFromFloat(val)
			respMap1[strings.ToLower(key)] = c
		}
	}
	return respMap1, nil
}
