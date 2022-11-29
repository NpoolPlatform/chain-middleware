package currencyvalue

import (
	"context"
	"strings"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	feedpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	valuecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/value"
	coinfeed "github.com/NpoolPlatform/chain-middleware/pkg/client/coin/currency/feed"
)

func CreateCurrency(ctx context.Context, in *valuemgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := valuecrud.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrency(ctx, info.ID.String())
}

func CreateCurrencies(ctx context.Context) error {
	offset := 0
	limit := 1000
	for {
		currencyFeeds, _, err := coinfeed.GetCurrencyFeeds(ctx, nil, int32(offset), int32(limit))
		if err != nil {
			logger.Sugar().Errorw("getCoinFeed", "offset", offset, "limit", limit, "error", err)
			return err
		}
		if len(currencyFeeds) == 0 {
			break
		}

		currencies := getLiveCurrencies(currencyFeeds)
		if err != nil {
			logger.Sugar().Errorw("processCurrencyFeeds", "offset", offset, "limit", limit, "error", err)
			return err
		}

		_, err = valuecrud.CreateBulk(ctx, currencies)
		if err != nil {
			return err
		}

		offset += limit
	}

	return nil
}

func getLiveCurrencies(feed []*feedpb.CurrencyFeed) []*valuemgrpb.CurrencyReq {
	coinMap := map[string]struct{}{}

	coinNames := []string{}
	for _, val := range feed {
		if _, ok := coinMap[val.CoinTypeID]; !ok {
			coinNames = append(coinNames, val.CoinName)
			coinMap[val.CoinTypeID] = struct{}{}
		}
	}

	isReq := false
	currencies := []*valuemgrpb.CurrencyReq{}
	coinPrice := map[string]string{}
	coinMap = map[string]struct{}{}

	for _, val := range feed {
		if _, ok := coinMap[val.CoinTypeID]; ok {
			continue
		}
		if MapCoinGecko(val.CoinName) == "" {
			continue
		}
		currencie := &valuemgrpb.CurrencyReq{
			CoinTypeID:   &val.CoinTypeID,
			FeedSourceID: &val.ID,
		}
		if PriceCoin(val.CoinName) {
			price := "1.0"
			currencie.MarketValueHigh = &price
			currencie.MarketValueLow = &price
			currencies = append(currencies, currencie)
			coinMap[val.CoinTypeID] = struct{}{}
		}

		switch val.FeedType {
		case feedmgrpb.FeedType_CoinGecko:
			if !isReq {
				price, err := CoinGeckoUSDPrices(coinNames)
				if err != nil {
					logger.Sugar().Errorw("CoinGeckoUSDPrices", "error", err)
					continue
				}
				isReq = true
				coinPrice = price
			}
			price, ok := coinPrice[MapCoinGecko(strings.ToLower(val.CoinName))]
			if ok {
				currencie.MarketValueHigh = &price
				currencie.MarketValueLow = &price
				currencies = append(currencies, currencie)
				coinMap[val.CoinTypeID] = struct{}{}
			}
		case feedmgrpb.FeedType_CoinBase:
			price, err := CoinBaseUSDPrice(val.CoinName)
			if err != nil {
				continue
			}
			currencie.MarketValueHigh = &price
			currencie.MarketValueLow = &price
			currencies = append(currencies, currencie)
			coinMap[val.CoinTypeID] = struct{}{}
		}
	}

	return currencies
}
