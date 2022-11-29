package currencyvalue

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

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

func CoinBaseUSDPrice(coinName string) (string, error) {
	coin := MapCoinBase(strings.ToLower(coinName))

	socksProxy := os.Getenv("ENV_CURRENCY_REQUEST_PROXY")

	url := strings.ReplaceAll(coinbaseAPI, "COIN", coin)

	cli := resty.New()
	cli = cli.SetTimeout(timeout * time.Second)
	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	resp, err := cli.R().Get(url)
	if err != nil {
		return "", fmt.Errorf("fail get currency %v: %v", coin, err)
	}
	r := apiResp{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return "", fmt.Errorf("fail parse currency %v: %v", coin, err)
	}

	if coin != r.Data.Base {
		return "", fmt.Errorf("invalid get coin currency from %v: %v", url, string(resp.Body()))
	}

	return r.Data.Amount, nil
}
