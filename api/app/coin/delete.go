package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoin(ctx context.Context, in *npool.DeleteCoinRequest) (*npool.DeleteCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
		)
		return &npool.DeleteCoinResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinResponse{
		Info: info,
	}, nil
}
