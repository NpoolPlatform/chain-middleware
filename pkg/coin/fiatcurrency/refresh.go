package fiatcurrency

import (
	"context"
	"fmt"

	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"
	coinmwpb "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"
	fiatcurrency "github.com/NpoolPlatform/chain-middleware/pkg/fiatcurrency"
)

func RefreshFiatCurrencies(ctx context.Context) error {
	offset := int32(0)
	const limit = int32(100)

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
			if coin.Presale {
				continue
			}
			names = append(names, coin.Name)
		}

		if len(names) == 0 {
			continue
		}

		prices, feedType, err := fiatcurrency.CoinUSDPrices(names)
		if err != nil {
			return err
		}

		currencies := []*fiatcurrencymgrpb.FiatCurrencyReq{}

		coinMap := map[string]*coinmwpb.Coin{}
		for _, coin := range coins {
			coinMap[coin.Name] = coin
		}

		for name, fiatcurrency := range prices {
			coin, ok := coinMap[name]
			if !ok {
				return fmt.Errorf("invalid coin: %v", name)
			}

			curr := fiatcurrency.String()

			currencies = append(currencies, &fiatcurrencymgrpb.FiatCurrencyReq{
				FiatTypeID:      &coin.ID,
				FeedType:        &feedType,
				MarketValueHigh: &curr,
				MarketValueLow:  &curr,
			})
		}

		if _, err = CreateFiatCurrencies(ctx, currencies); err != nil {
			return err
		}

		offset += limit
	}
}
