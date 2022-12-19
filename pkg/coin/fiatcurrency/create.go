package fiatcurrency

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	fiatcurrencymgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/fiatcurrency"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	npoolpb "github.com/NpoolPlatform/message/npool"
)

func CreateFiatCurrency(ctx context.Context, in *fiatcurrencymgrpb.FiatCurrencyReq) (*npool.FiatCurrency, error) {
	_, err := fiatcurrencymgrcli.CreateFiatCurrency(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetFiatCurrency(ctx, in.GetFiatCurrencyTypeID())
}

func CreateFiatCurrencies(ctx context.Context, in []*fiatcurrencymgrpb.FiatCurrencyReq) ([]*npool.FiatCurrency, error) {
	infos, err := fiatcurrencymgrcli.CreateFiatCurrencies(ctx, in)
	if err != nil {
		return nil, err
	}
	ids := []string{}
	for _, val := range infos {
		ids = append(ids, val.FiatCurrencyTypeID)
	}
	return GetFiatCurrencies(ctx, &npool.Conds{
		FiatCurrencyTypeIDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	})
}
