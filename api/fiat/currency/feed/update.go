package currencyfeed

import (
	"context"

	currencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency/feed"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFeed(ctx context.Context, in *npool.UpdateFeedRequest) (*npool.UpdateFeedResponse, error) {
	req := in.GetInfo()
	handler, err := currencyfeed1.NewHandler(
		ctx,
		currencyfeed1.WithID(req.ID, true),
		currencyfeed1.WithFeedFiatName(req.FeedFiatName, false),
		currencyfeed1.WithDisabled(req.Disabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFeed(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFeedResponse{
		Info: info,
	}, nil
}
