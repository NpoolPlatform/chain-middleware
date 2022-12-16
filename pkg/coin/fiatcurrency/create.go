package fiatcurrency

import (
	"context"

	fiatcurrencymgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/fiatcurrency"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"
)

func CreateFiatCurrency(ctx context.Context, in *fiatcurrencymgrpb.FiatCurrencyReq) (*npool.FiatCurrency, error) {
	_, err := fiatcurrencymgrcli.CreateFiatCurrency(ctx, in)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func CreateFiatCurrencies(ctx context.Context, in []*fiatcurrencymgrpb.FiatCurrencyReq) ([]*npool.FiatCurrency, error) {
	_, err := fiatcurrencymgrcli.CreateFiatCurrencies(ctx, in)
	if err != nil {
		return nil, err
	}

	return nil, err
}
