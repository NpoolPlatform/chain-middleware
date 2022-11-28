package currencyvalue

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	npoolpb "github.com/NpoolPlatform/message/npool"

	valuecrud "github.com/NpoolPlatform/chain-manager/pkg/crud/coin/currency/value"
)

func CreateCurrency(ctx context.Context, in *valuemgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := valuecrud.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrency(ctx, info.ID.String())
}

func CreateCurrencies(ctx context.Context, in []*valuemgrpb.CurrencyReq) ([]*npool.Currency, error) {
	infos, err := valuecrud.CreateBulk(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, val := range infos {
		ids = append(ids, val.CoinTypeID.String())
	}

	currencies, err := GetCurrencies(ctx, &npool.Conds{
		CoinTypeIDs: &npoolpb.StringSliceVal{
			Op:    cruder.IN,
			Value: ids,
		},
	})
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
