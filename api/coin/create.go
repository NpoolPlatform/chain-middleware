//nolint:nolintlint,dupl
package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
)

func (s *Server) CreateCoin(
	ctx context.Context,
	in *npool.CreateCoinRequest,
) (
	*npool.CreateCoinResponse,
	error,
) {
	return nil, nil
}
