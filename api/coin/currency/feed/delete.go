package currencyfeed

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	feed1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/feed"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) DeleteCurrencyFeed(ctx context.Context, in *npool.DeleteCurrencyFeedRequest) (*npool.DeleteCurrencyFeedResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteCurrencyFeed", "ID", in.GetID(), "error", err)
		return &npool.DeleteCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := feed1.DeleteCurrencyFeed(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteCurrencyFeed", "error", err)
		return &npool.DeleteCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCurrencyFeedResponse{
		Info: info,
	}, nil
}
