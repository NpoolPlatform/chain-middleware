package currencyfeed

import (
	"context"

	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	feedmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/currency/feed"
)

func CreateCurrencyFeed(ctx context.Context, in *feedmgrpb.CurrencyFeedReq) (*npool.CurrencyFeed, error) {
	info, err := feedmgrcli.CreateCurrencyFeed(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCurrencyFeed(ctx, info.ID)
}
