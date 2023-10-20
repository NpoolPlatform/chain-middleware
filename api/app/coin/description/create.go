package description

import (
	"context"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin/description"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinDescription(ctx context.Context, in *npool.CreateCoinDescriptionRequest) (*npool.CreateCoinDescriptionResponse, error) {
	req := in.GetInfo()
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(req.ID, false),
		description1.WithEntID(req.EntID, false),
		description1.WithAppID(req.AppID, true),
		description1.WithCoinTypeID(req.CoinTypeID, true),
		description1.WithUsedFor(req.UsedFor, true),
		description1.WithTitle(req.Title, true),
		description1.WithMessage(req.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: info,
	}, nil
}
