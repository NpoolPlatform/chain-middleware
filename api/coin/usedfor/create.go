package coinusedfor

import (
	"context"

	coinusedfor1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/usedfor"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinUsedFor(ctx context.Context, in *npool.CreateCoinUsedForRequest) (*npool.CreateCoinUsedForResponse, error) {
	req := in.GetInfo()
	handler, err := coinusedfor1.NewHandler(
		ctx,
		coinusedfor1.WithCoinTypeID(req.CoinTypeID, true),
		coinusedfor1.WithUsedFor(req.UsedFor, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinUsedForResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoinUsedFor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoinUsedFor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinUsedForResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinUsedForResponse{
		Info: info,
	}, nil
}
