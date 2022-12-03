package currency

import (
	"context"

	currencymgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/currency"
	currencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func CreateCurrency(ctx context.Context, in *currencymgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := currencymgrcli.CreateCurrency(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrency(ctx, info.ID)
}

func CreateCurrencies(ctx context.Context, in []*currencymgrpb.CurrencyReq) ([]*npool.Currency, error) {
	infos, err := currencymgrcli.CreateCurrencies(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.CoinTypeID)
	}

	return GetCurrencies(ctx, &npool.Conds{
		CoinTypeIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	})
}
