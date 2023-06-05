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

func (s *Server) UpdateTx(ctx context.Context, in *npool.UpdateTxRequest) (*npool.UpdateTxResponse, error) {
	req := in.GetInfo()
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(req.ID),
		tx1.WithChainTxID(req.ChainTxID),
		tx1.WithState(req.State),
		tx1.WithExtra(req.Extra),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTx",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: info,
	}, nil
}
