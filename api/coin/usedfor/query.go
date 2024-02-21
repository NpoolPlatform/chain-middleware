package coinusedfor

import (
	"context"

	coinusedfor1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/usedfor"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoinUsedFor(ctx context.Context, in *npool.GetCoinUsedForRequest) (*npool.GetCoinUsedForResponse, error) {
	handler, err := coinusedfor1.NewHandler(
		ctx,
		coinusedfor1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCoinUsedFor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinUsedForResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinUsedFors(ctx context.Context, in *npool.GetCoinUsedForsRequest) (*npool.GetCoinUsedForsResponse, error) {
	handler, err := coinusedfor1.NewHandler(
		ctx,
		coinusedfor1.WithConds(in.Conds),
		coinusedfor1.WithOffset(in.GetOffset()),
		coinusedfor1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFors",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoinUsedFors(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinUsedFors",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinUsedForsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinUsedForsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
