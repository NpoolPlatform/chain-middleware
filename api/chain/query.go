package chain

import (
	"context"

	chain1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/chain"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetChain(ctx context.Context, in *npool.GetChainRequest) (*npool.GetChainResponse, error) {
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChain",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetChain(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChain",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChainResponse{
		Info: info,
	}, nil
}

func (s *Server) GetChains(ctx context.Context, in *npool.GetChainsRequest) (*npool.GetChainsResponse, error) {
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithConds(in.Conds),
		chain1.WithOffset(in.GetOffset()),
		chain1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChains",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetChains(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetChains",
			"In", in,
			"Error", err,
		)
		return &npool.GetChainsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChainsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
