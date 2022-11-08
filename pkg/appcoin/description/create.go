package description

import (
	"context"

	descmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/appcoin/description"
	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"
)

func CreateCoinDescription(ctx context.Context, in *descmgrpb.CoinDescriptionReq) (*npool.CoinDescription, error) {
	info, err := descmgrcli.CreateCoinDescription(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetCoinDescription(ctx, info.ID)
}
