package currencyfeed

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	feed1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/feed"

	feedmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/feed"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	coinmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func ValidateCreate(ctx context.Context, in *feedmgrpb.CurrencyFeedReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCurrencyFeed", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateCurrencyFeed", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}

	_, err := coinmgrcli.GetCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrencyFeed", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}

	if in.GetFeedSource() == "" {
		logger.Sugar().Errorw("CreateCurrencyFeed", "FeedSource", in.GetFeedSource(), "error", "FeedSource is invalid")
		return fmt.Errorf("feedsource is invalid")
	}

	switch in.GetFeedType() {
	case feedmgrpb.FeedType_CoinBase:
	case feedmgrpb.FeedType_CoinGecko:
	default:
		logger.Sugar().Errorw("CreateCurrencyFeed", "FeedType", in.GetFeedType(), "error", "FeedType is invalid")
		return fmt.Errorf("feedtype is invalid")
	}

	return nil
}

func (s *Server) CreateCurrencyFeed(ctx context.Context, in *npool.CreateCurrencyFeedRequest) (*npool.CreateCurrencyFeedResponse, error) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateCurrencyFeedResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := feed1.CreateCurrencyFeed(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrencyFeed", "error", err)
		return &npool.CreateCurrencyFeedResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyFeedResponse{
		Info: info,
	}, nil
}
