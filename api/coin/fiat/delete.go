package coinfiat

import (
	"context"

	coinfiat1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin/fiat"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiat"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCoinFiat(ctx context.Context, in *npool.DeleteCoinFiatRequest) (*npool.DeleteCoinFiatResponse, error) {
	req := in.GetInfo()
	handler, err := coinfiat1.NewHandler(
		ctx,
		coinfiat1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinFiatResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCoinFiat(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCoinFiat",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCoinFiatResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinFiatResponse{
		Info: info,
	}, nil
}
