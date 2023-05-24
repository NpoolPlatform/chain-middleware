//nolint:dupl
package currencyfeed

import (
	"context"

	currencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/currency/feed"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFeed(ctx context.Context, in *npool.CreateFeedRequest) (*npool.CreateFeedResponse, error) {
	req := in.GetInfo()
	handler, err := currencyfeed1.NewHandler(
		ctx,
		currencyfeed1.WithID(req.ID),
		currencyfeed1.WithCoinTypeID(req.CoinTypeID),
		currencyfeed1.WithFeedType(req.FeedType),
		currencyfeed1.WithFeedCoinName(req.FeedCoinName),
		currencyfeed1.WithDisabled(req.Disabled),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFeed(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeed",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeedResponse{
		Info: info,
	}, nil
}