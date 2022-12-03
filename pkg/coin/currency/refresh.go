package currency

import (
	"context"
	"fmt"

	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"
	currency "github.com/NpoolPlatform/chain-middleware/pkg/currency"
)

func RefreshCurrencies(ctx context.Context) error {
	offset := int32(0)
	limit := int32(100)

	for {
		coins, _, err := coin1.GetCoins(ctx, &coinmwpb.Conds{}, offset, limit)
		if err != nil {
			return err
		}
		if len(coins) == 0 {
			return nil
		}

		names := []string{}
		for _, coin := range coins {
			if coin.StableUSD {
				continue
			}
			names = append(names, coin.Name)
		}

		if len(names) == 0 {
			continue
		}

		prices, feedType, err := currency.CoinUSDPrices(ctx, names)
		if err != nil {
			return err
		}

		currencies := []*currencymgrpb.CurrencyReq{}

		coinMap := map[string]*coinmwpb.Coin{}
		for _, coin := range coins {
			coinMap[coin.Name] = coin
		}

		for name, currency := range prices {
			coin, ok := coinMap[name]
			if !ok {
				return fmt.Errorf("invalid coin")
			}

			curr := currency.String()

			currencies = append(currencies, &currencymgrpb.CurrencyReq{
				CoinTypeID:      &coin.ID,
				FeedType:        &feedType,
				MarketValueHigh: &curr,
				MarketValueLow:  &curr,
			})
		}

		if _, err = CreateCurrencies(ctx, currencies); err != nil {
			return err
		}
	}

	return nil
}
