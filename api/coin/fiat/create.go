package coinfiat

import (
	"context"

	coinfiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/fiat"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinFiat(ctx context.Context, in *npool.CreateCoinFiatRequest) (*npool.CreateCoinFiatResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoinFiat",
			"In", in,
		)
		return &npool.CreateCoinFiatResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := coinfiat1.NewHandler(
		ctx,
		coinfiat1.WithCoinTypeID(req.CoinTypeID, true),
		coinfiat1.WithFiatID(req.FiatID, true),
		coinfiat1.WithFeedType(req.FeedType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinFiatResponse{
		Info: info,
	}, nil
}
