package currencyvalue

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"os"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/go-resty/resty/v2"
)

const (
	coinbaseAPI = "https://api.coinbase.com/v2/exchange-rates?currency=USD"
	timeout     = 5
)

type apiData struct {
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

type apiResp struct {
	Data apiData `json:"data"`
}

func UsdFaitCurrency() (map[string]decimal.Decimal, error) {
	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	logger.Sugar().Errorw("CoinBaseUSDPrice", "URL", coinbaseAPI)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(coinbaseAPI)
	if err != nil {
		logger.Sugar().Errorw("CoinBaseUSDPrice", "error", err)
		return nil, err
	}
	r := apiResp{}
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
