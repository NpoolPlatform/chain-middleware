package currencyvalue

import (
	"context"

	value1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/value"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RefreshCurrencies(ctx context.Context, in *npool.RefreshCurrenciesRequest) (*npool.RefreshCurrenciesResponse, error) {
	if err := value1.RefreshCurrencies(ctx); err != nil {
		return &npool.RefreshCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &npool.RefreshCurrenciesResponse{}, nil
}
