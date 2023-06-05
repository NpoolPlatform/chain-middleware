package currencyfeed

import (
	"context"

	currencyfeed1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/fiat/currency/feed"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFeeds(ctx context.Context, in *npool.GetFeedsRequest) (*npool.GetFeedsResponse, error) {
	handler, err := currencyfeed1.NewHandler(
		ctx,
		currencyfeed1.WithConds(in.Conds),
		currencyfeed1.WithOffset(in.GetOffset()),
		currencyfeed1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeds",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeedsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFeeds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeds",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeedsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
