//nolint:nolintlint,dupl
package tran

import (
	"context"

	tx1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/tx"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistTxConds(ctx context.Context, in *npool.ExistTxCondsRequest) (*npool.ExistTxCondsResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithConds(in.Conds),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTxConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.ExistTxConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTxConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxCondsResponse{
		Info: info,
	}, nil
}
