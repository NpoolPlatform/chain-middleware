package description

import (
	"context"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin/description"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoinDescription(ctx context.Context, in *npool.UpdateCoinDescriptionRequest) (*npool.UpdateCoinDescriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCoinDescription",
			"In", in,
		)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(req.ID, true),
		description1.WithTitle(req.Title, false),
		description1.WithMessage(req.Message, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: info,
	}, nil
}
