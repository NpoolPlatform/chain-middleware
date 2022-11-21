package currencyvalue

import (
	"context"

	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	valuemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/currency/value"
)

func CreateCurrency(ctx context.Context, in *valuemgrpb.CurrencyReq) (*npool.Currency, error) {
	info, err := valuemgrcli.CreateCurrency(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrency(ctx, info.ID)
}
