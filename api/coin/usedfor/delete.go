package coinusedfor

import (
	"context"

	coinusedfor1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/usedfor"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoinUsedFor(ctx context.Context, in *npool.DeleteCoinUsedForRequest) (*npool.DeleteCoinUsedForResponse, error) {
	req := in.GetInfo()
	handler, err := coinusedfor1.NewHandler(
		ctx,
		coinusedfor1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinUsedForResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoinUsedFor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinUsedForResponse{
		Info: info,
	}, nil
}
