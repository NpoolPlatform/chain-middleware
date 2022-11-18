package currencyfeed

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	feed1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/feed"

	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func ValidateUpdate(in *feedmgrpb.CurrencyFeedReq) error {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCurrencyFeed", "ID", in.GetID(), "error", err)
		return err
	}

	if in.GetFeedSource() == "" {
		logger.Sugar().Errorw("UpdateCurrencyFeed", "FeedSource", in.GetFeedSource(), "error", "FeedSource is invalid")
		return fmt.Errorf("feedsource is invalid")
	}

	return nil
}

func (s *Server) UpdateCurrencyFeed(ctx context.Context, in *npool.UpdateCurrencyFeedRequest) (*npool.UpdateCurrencyFeedResponse, error) {
	if err := ValidateUpdate(in.GetInfo()); err != nil {
		return &npool.UpdateCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := feed1.UpdateCurrencyFeed(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCurrencyFeed", "error", err)
		return &npool.UpdateCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCurrencyFeedResponse{
		Info: info,
	}, nil
}
