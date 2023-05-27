package currency

import (
	"context"

	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCurrency(ctx context.Context, in *npool.GetCurrencyRequest) (*npool.GetCurrencyResponse, error) {
	handler, err := currency1.NewHandler(
		ctx,
		currency1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrency",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCurrency(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCurrency",
			"In", in,
			"Error", err,
		)
		return &npool.GetCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	handler, err := currency1.NewHandler(
		ctx,
		currency1.WithConds(in.Conds),
		currency1.WithOffset(in.GetOffset()),
		currency1.WithLimit(in.GetLimit()),
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
