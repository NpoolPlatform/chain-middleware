package fiat

import (
	"context"

	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFiat(ctx context.Context, in *npool.GetFiatRequest) (*npool.GetFiatResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiat",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiat",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFiats(ctx context.Context, in *npool.GetFiatsRequest) (*npool.GetFiatsResponse, error) {
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithConds(in.Conds),
		fiat1.WithOffset(in.GetOffset()),
		fiat1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFiats(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFiats",
			"In", in,
			"Error", err,
		)
		return &npool.GetFiatsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
