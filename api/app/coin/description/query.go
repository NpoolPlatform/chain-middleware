package description

import (
	"context"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin/description"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoinDescription(ctx context.Context, in *npool.GetCoinDescriptionRequest) (*npool.GetCoinDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinDescriptions(ctx context.Context, in *npool.GetCoinDescriptionsRequest) (*npool.GetCoinDescriptionsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithConds(in.Conds),
		description1.WithOffset(in.GetOffset()),
		description1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoinDescriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
