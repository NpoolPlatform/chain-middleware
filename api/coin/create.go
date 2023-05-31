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
		coin1.WithName(req.Name),
		coin1.WithUnit(req.Unit),
		coin1.WithENV(req.ENV),
		coin1.WithChainType(req.ChainType),
		coin1.WithChainNativeUnit(req.ChainNativeUnit),
		coin1.WithChainAtomicUnit(req.ChainAtomicUnit),
		coin1.WithChainUnitExp(req.ChainUnitExp),
		coin1.WithGasType(req.GasType),
		coin1.WithChainID(req.ChainID),
		coin1.WithChainNickname(req.ChainNickname),
		coin1.WithChainNativeCoinName(req.ChainNativeCoinName),
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
