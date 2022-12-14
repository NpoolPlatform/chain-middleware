package fiat

import (
	"context"

	fiatcurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/fiat"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) RefreshFiatCurrencies(
	ctx context.Context,
	in *npool.RefreshFiatCurrenciesRequest,
) (
	*npool.RefreshFiatCurrenciesResponse,
	error,
) {
	if err := fiatcurrency1.RefreshFiatCurrencies(ctx); err != nil {
		return &npool.RefreshFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &npool.RefreshFiatCurrenciesResponse{}, nil
}
