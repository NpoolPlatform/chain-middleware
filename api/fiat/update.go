//nolint:dupl
package fiat

import (
	"context"

	fiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateFiat(ctx context.Context, in *npool.UpdateFiatRequest) (*npool.UpdateFiatResponse, error) {
	req := in.GetInfo()
	handler, err := fiat1.NewHandler(
		ctx,
		fiat1.WithID(req.ID),
		fiat1.WithName(req.Name),
		fiat1.WithLogo(req.Logo),
		fiat1.WithUnit(req.Unit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFiat",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFiatResponse{
		Info: info,
	}, nil
}
