//nolint:nolintlint,dupl
package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCoin(ctx context.Context, in *npool.ExistCoinRequest) (*npool.ExistCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoin",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoin",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistCoinConds(ctx context.Context, in *npool.ExistCoinCondsRequest) (*npool.ExistCoinCondsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinCondsResponse{
		Info: info,
	}, nil
}
