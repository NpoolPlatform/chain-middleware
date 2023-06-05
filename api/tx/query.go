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

func (s *Server) GetTx(ctx context.Context, in *npool.GetTxRequest) (*npool.GetTxResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTx",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTx",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTxs(ctx context.Context, in *npool.GetTxsRequest) (*npool.GetTxsResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithConds(in.Conds),
		tx1.WithOffset(in.GetOffset()),
		tx1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetTxs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTxs",
			"In", in,
			"Error", err,
		)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
