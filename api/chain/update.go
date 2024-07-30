package chain

import (
	"context"

	chain1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/chain"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/chain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateChain(ctx context.Context, in *npool.UpdateChainRequest) (*npool.UpdateChainResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateChain",
			"In", in,
		)
		return &npool.UpdateChainResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := chain1.NewHandler(
		ctx,
		chain1.WithID(req.ID, false),
		chain1.WithEntID(req.EntID, false),
		chain1.WithChainType(req.ChainType, false),
		chain1.WithNativeUnit(req.NativeUnit, false),
		chain1.WithAtomicUnit(req.AtomicUnit, false),
		chain1.WithUnitExp(req.UnitExp, false),
		chain1.WithENV(req.ENV, false),
		chain1.WithChainID(req.ChainID, false),
		chain1.WithNickname(req.NickName, false),
		chain1.WithGasType(req.GasType, false),
		chain1.WithLogo(req.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateChain",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateChainResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = handler.UpdateChain(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateChain",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateChainResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateChainResponse{
		Info: nil,
	}, nil
}
