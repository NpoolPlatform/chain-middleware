package chain

import (
	"context"

	chain1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/chain"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateChain(ctx context.Context, in *npool.CreateChainRequest) (*npool.CreateChainResponse, error) {
	req := in.GetInfo()
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithEntID(req.EntID, false),
		chain1.WithChainType(req.ChainType, true),
		chain1.WithNativeUnit(req.NativeUnit, true),
		chain1.WithAtomicUnit(req.AtomicUnit, true),
		chain1.WithUnitExp(req.UnitExp, true),
		chain1.WithENV(req.ENV, true),
		chain1.WithChainID(req.ChainID, false),
		chain1.WithNickname(req.NickName, false),
		chain1.WithGasType(req.GasType, true),
		chain1.WithLogo(req.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateChain",
			"In", in,
			"Error", err,
		)
		return &npool.CreateChainResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.CreateChain(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateChain",
			"In", in,
			"Error", err,
		)
		return &npool.CreateChainResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateChainResponse{
		Info: nil,
	}, nil
}
