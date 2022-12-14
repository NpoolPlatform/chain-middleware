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

func UsdFiatCurrency(FiatCurrencyName []string) (map[string]decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	url := fmt.Sprintf("%v%v?ids=usd&vs_currencies=%v", coinGeckoAPI, "/simple/price", FiatCurrencyName)
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

	respMap1 := map[string]decimal.Decimal{}

	if _, ok := respMap["usd"]; ok {
		for key, val := range respMap["usd"] {
			c := decimal.NewFromFloat(val)
			respMap1[strings.ToLower(key)] = c
		}
	}
	return respMap1, nil
}
