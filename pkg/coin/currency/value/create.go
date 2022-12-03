package currencyvalue

import (
	"context"

	valuemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/currency/value"
	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func CreateCurrency(ctx context.Context, in *valuemgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := valuemgrcli.CreateCurrency(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrency(ctx, info.ID)
}

func CreateCurrencies(ctx context.Context, in []*valuemgrpb.CurrencyReq) ([]*npool.Currency, error) {
	infos, err := valuemgrcli.CreateCurrencies(ctx, in)
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
