package coin

import (
	"context"
	"fmt"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	key := fmt.Sprintf(
		"%v:%v:%v",
		basetypes.Prefix_PrefixCreateCoin,
		req.GetName(),
		req.GetENV(),
	)
	if err := redis2.TryLock(key, 0); err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

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
