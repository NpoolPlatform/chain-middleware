package currencyfeed

import (
	"context"

	feedmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"
)

func DeleteCurrencyFeed(ctx context.Context, id string) (*npool.CurrencyFeed, error) {
	info, err := GetCurrencyFeed(ctx, id)
	if err != nil {
		return nil, err
	}

	_, err = feedmgrcli.DeleteCurrencyFeed(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}
