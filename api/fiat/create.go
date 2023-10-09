package fiat

import (
	"context"

	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateFiat(ctx context.Context, in *npool.CreateFiatRequest) (*npool.CreateFiatResponse, error) {
	req := in.GetInfo()
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithID(req.ID, false),
		fiat1.WithEntID(req.EntID, false),
		fiat1.WithName(req.Name, true),
		fiat1.WithLogo(req.Logo, true),
		fiat1.WithUnit(req.Unit, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatResponse{
		Info: info,
	}, nil
}
