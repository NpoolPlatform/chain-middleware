package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *npool.CreateCoinRequest) (*npool.CreateCoinResponse, error) {
	req := in.GetInfo()
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithName(req.Name, true),
		coin1.WithUnit(req.Unit, true),
		coin1.WithENV(req.ENV, true),
		coin1.WithChainType(req.ChainType, true),
		coin1.WithChainNativeUnit(req.ChainNativeUnit, true),
		coin1.WithChainAtomicUnit(req.ChainAtomicUnit, true),
		coin1.WithChainUnitExp(req.ChainUnitExp, true),
		coin1.WithGasType(req.GasType, true),
		coin1.WithChainID(req.ChainID, true),
		coin1.WithChainNickname(req.ChainNickname, true),
		coin1.WithChainNativeCoinName(req.ChainNativeCoinName, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info,
	}, nil
}
