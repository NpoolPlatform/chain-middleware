//nolint:dupl
package currencyhistory

import (
	"context"

	history1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/fiat/currency/history"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat/currency/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.Conds),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrencies",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCurrencies(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrencies",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrenciesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
