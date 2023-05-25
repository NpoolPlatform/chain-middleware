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

func (s *Server) CreateTx(ctx context.Context, in *npool.CreateTxRequest) (*npool.CreateTxResponse, error) {
	req := in.GetInfo()
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithID(req.ID),
		tx1.WithCoinTypeID(req.CoinTypeID),
		tx1.WithFromAccountID(req.FromAccountID),
		tx1.WithToAccountID(req.ToAccountID),
		tx1.WithAmount(req.Amount),
		tx1.WithFeeAmount(req.FeeAmount),
		tx1.WithChainTxID(req.ChainTxID),
		tx1.WithState(req.State),
		tx1.WithExtra(req.Extra),
		tx1.WithType(req.Type),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTx",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateTx(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTx",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateTxs(ctx context.Context, in *npool.CreateTxsRequest) (*npool.CreateTxsResponse, error) {
	handler, err := tx1.NewHandler(
		ctx,
		tx1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTxs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateTxs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTxs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxsResponse{
		Infos: infos,
	}, nil
}
