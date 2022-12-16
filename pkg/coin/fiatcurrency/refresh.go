package fiatcurrency

import (
	"context"

	typemgrent "github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	fiatcurrency "github.com/NpoolPlatform/chain-middleware/pkg/fiatcurrency"

	typecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/fiatcurrencytype"
)

func RefreshFiatCurrencies(ctx context.Context) error {
	offset := int32(0)
	const limit = int32(100)

	for {
		faitCurrencyTypes, _, err := typecrud.Rows(ctx, nil, int(offset), int(limit))
		if err != nil {
			return err
		}
		if len(faitCurrencyTypes) == 0 {
			return nil
		}

		names := []string{}
		for _, faitCurrencyType := range faitCurrencyTypes {
			names = append(names, faitCurrencyType.Name)
		}

		if len(names) == 0 {
			continue
		}

		prices, feedType, err := fiatcurrency.CoinUSDPrices(names)
		if err != nil {
			return err
		}

		typeMap := map[string]*typemgrent.FiatCurrencyType{}
		for _, faitCurrencyType := range faitCurrencyTypes {
			typeMap[faitCurrencyType.Name] = faitCurrencyType
		}

		currencies := []*fiatcurrencymgrpb.FiatCurrencyReq{}

		for name, fiatcurrency := range prices {
			tp, ok := typeMap[name]
			if ok {
				curr := fiatcurrency.String()
				tpID := tp.ID.String()
				currencies = append(currencies, &fiatcurrencymgrpb.FiatCurrencyReq{
					FiatCurrencyTypeID: &tpID,
					FeedType:           &feedType,
					MarketValueHigh:    &curr,
					MarketValueLow:     &curr,
				})
			}
		}

		if _, err = CreateFiatCurrencies(ctx, currencies); err != nil {
			return err
		}

		offset += limit
	}
}
