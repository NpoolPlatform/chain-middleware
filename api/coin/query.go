//nolint:nolintlint,dupl
package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/const"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) GetCoin(
	ctx context.Context,
	in *npool.GetCoinRequest,
) (
	*npool.GetCoinResponse,
	error,
) {
	var err error

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetCoin", "ID", in.GetID(), "error", err)
		return &npool.GetCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := coin1.GetCoin(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetCoin", "ID", in.GetID(), "error", err)
		return &npool.GetCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoins(
	ctx context.Context,
	in *npool.GetCoinsRequest,
) (
	*npool.GetCoinsResponse,
	error,
) {
	var err error

	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCoins", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.ENV != nil {
		switch conds.GetENV().GetValue() {
		case "main", "test":
		default:
			logger.Sugar().Errorw("GetCoins", "ENV", conds.GetENV().GetValue(), "error", err)
			return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Name != nil && conds.GetName().GetValue() == "" {
		logger.Sugar().Errorw("GetCoins", "Name", conds.GetName().GetValue(), "error", "Name is empty")
		return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, "Name is empty")
	}
	for _, id := range conds.GetIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetCoins", "IDs", conds.GetIDs().GetValue(), "error", err)
			return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant.DefaultRowLimit
	}

	infos, total, err := coin1.GetCoins(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCoins", "Conds", conds, "error", err)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetCoinOnly(
	ctx context.Context,
	in *npool.GetCoinOnlyRequest,
) (
	*npool.GetCoinOnlyResponse,
	error,
) {
	var err error

	conds := in.GetConds()
	if conds == nil {
		conds = &npool.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCoinOnly", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetCoinOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.ENV != nil {
		switch conds.GetENV().GetValue() {
		case "main", "test":
		default:
			logger.Sugar().Errorw("GetCoinOnly", "ENV", conds.GetENV().GetValue(), "error", err)
			return &npool.GetCoinOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.Name != nil && conds.GetName().GetValue() == "" {
		logger.Sugar().Errorw("GetCoinOnly", "Name", conds.GetName().GetValue(), "error", "Name is empty")
		return &npool.GetCoinOnlyResponse{}, status.Error(codes.InvalidArgument, "Name is empty")
	}
	for _, id := range conds.GetIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetCoinOnly", "IDs", conds.GetIDs().GetValue(), "error", err)
			return &npool.GetCoinOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	info, err := coin1.GetCoinOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorw("GetCoinOnly", "Conds", conds, "error", err)
		return &npool.GetCoinOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinOnlyResponse{
		Info: info,
	}, nil
}
