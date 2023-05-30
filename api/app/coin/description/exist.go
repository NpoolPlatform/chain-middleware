package description

import (
	"context"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin/description"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint:lll
func (s *Server) ExistCoinDescriptionConds(ctx context.Context, in *npool.ExistCoinDescriptionCondsRequest) (*npool.ExistCoinDescriptionCondsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinDescriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinDescriptionCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistCoinDescriptionConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCoinDescriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCoinDescriptionCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinDescriptionCondsResponse{
		Info: info,
	}, nil
}
