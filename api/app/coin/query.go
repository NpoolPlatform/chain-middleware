package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCoin(ctx context.Context, in *npool.GetCoinRequest) (*npool.GetCoinResponse, error) {
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(&in.ID),
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
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithConds(in.Conds),
		appcoin1.WithOffset(in.GetOffset()),
		appcoin1.WithLimit(in.GetLimit()),
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
