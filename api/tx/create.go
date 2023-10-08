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
		tx1.WithID(req.ID, false),
		tx1.WithEntID(req.EntID, false),
		tx1.WithCoinTypeID(req.CoinTypeID, true),
		tx1.WithFromAccountID(req.FromAccountID, true),
		tx1.WithToAccountID(req.ToAccountID, true),
		tx1.WithAmount(req.Amount, true),
		tx1.WithFeeAmount(req.FeeAmount, true),
		tx1.WithChainTxID(req.ChainTxID, true),
		tx1.WithState(req.State, true),
		tx1.WithExtra(req.Extra, true),
		tx1.WithType(req.Type, true),
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
		tx1.WithReqs(in.GetInfos(), true),
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
