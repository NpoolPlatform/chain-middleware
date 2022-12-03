package currency

import (
	"context"

	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RefreshCurrencies(ctx context.Context, in *npool.RefreshCurrenciesRequest) (*npool.RefreshCurrenciesResponse, error) {
	if err := currency1.RefreshCurrencies(ctx); err != nil {
		return &npool.RefreshCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &npool.RefreshCurrenciesResponse{}, nil
}
