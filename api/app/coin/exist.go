package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistCoinConds(ctx context.Context, in *npool.ExistCoinCondsRequest) (*npool.ExistCoinCondsResponse, error) {
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithConds(in.Conds),
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
