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

const (
	coinGeckoAPI = "https://api.coingecko.com/api/v3"
)

func CoinGeckoUSDPrices(coinNames []string) (map[string]string, error) {
	coins := ""
	for _, val := range coinNames {
		coin := MapCoinGecko(strings.ToLower(val))
		if coin != "" {
			coins += fmt.Sprintf("%v,", MapCoinGecko(strings.ToLower(val)))
		}
	}
	coins = coins[:len(coins)-1]

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")
	url := fmt.Sprintf("%v%v?ids=%v&vs_currencies=usd", coinGeckoAPI, "/simple/price", coins)
	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		return nil, fmt.Errorf("fail get currency %v: %v", coins, err)
	}
	respMap := map[string]map[string]float64{}
	err = json.Unmarshal(resp.Body(), &respMap)
	if err != nil {
		return nil, fmt.Errorf("fail parse currency %v: %v", coins, err)
	}

	infoMap := map[string]string{}
	for key, val := range respMap {
		price := ""
		for _, v := range val {
			price = decimal.NewFromFloat(v).String()
		}
		infoMap[key] = price
	}

	return infoMap, nil
}
