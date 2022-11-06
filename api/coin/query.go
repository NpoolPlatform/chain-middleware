//nolint:nolintlint,dupl
package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
)

func (s *Server) GetCoin(
	ctx context.Context,
	in *npool.GetCoinRequest,
) (
	*npool.GetCoinResponse,
	error,
) {
	return nil, nil
}

func (s *Server) GetCoins(
	ctx context.Context,
	in *npool.GetCoinsRequest,
) (
	*npool.GetCoinsResponse,
	error,
) {
	return nil, nil
}
