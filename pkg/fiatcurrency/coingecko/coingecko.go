package currencyvalue

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/shopspring/decimal"

	"github.com/go-resty/resty/v2"
)

const (
	coinGeckoAPI = "https://api.coingecko.com/api/v3"
	timeout      = 5
)

func CoinGeckoFiatCurrency(FiatCurrencyName []string) (map[string]decimal.Decimal, error) {
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

	infoMap := map[string]decimal.Decimal{}
	for _, val := range respMap {
		price := decimal.NewFromInt(0)
		for _, v := range val {
			price = decimal.NewFromFloat(v)
		}
		infoMap[coin] = price
	}

	return infoMap, nil
}
