package fiatcurrency

import (
	"context"

	fiatcurrencymgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/fiatcurrency"
	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/fiatcurrency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func CreateFiatCurrency(ctx context.Context, in *fiatcurrencymgrpb.FiatCurrencyReq) (*npool.FiatCurrency, error) {
	info, err := fiatcurrencymgrcli.CreateFiatCurrency(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetFiatCurrency(ctx, info.ID)
}

func CreateFiatCurrencies(ctx context.Context, in []*fiatcurrencymgrpb.FiatCurrencyReq) ([]*npool.FiatCurrency, error) {
	infos, err := fiatcurrencymgrcli.CreateFiatCurrencies(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.FiatTypeID)
	}

	return GetFiatCurrencies(ctx, &npool.Conds{
		FiatCurrencyTypeIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	})
}
