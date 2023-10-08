//nolint:nolintlint,dupl
package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoin(ctx context.Context, in *npool.GetCoinRequest) (*npool.GetCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoins(ctx context.Context, in *npool.GetCoinsRequest) (*npool.GetCoinsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
		coin1.WithOffset(in.GetOffset()),
		coin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCoinOnly(ctx context.Context, in *npool.GetCoinOnlyRequest) (*npool.GetCoinOnlyResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetCoinOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCoinOnly",
			"In", in,
			"Error", err,
		)
		return &npool.GetCoinOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinOnlyResponse{
		Info: info,
	}, nil
}
