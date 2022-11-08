package coin

import (
	"context"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	commonpb "github.com/NpoolPlatform/message/npool"

	coinmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"
	coinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func (s *Server) CreateCoin(
	ctx context.Context,
	in *npool.CreateCoinRequest,
) (
	*npool.CreateCoinResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoin", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetInfo().GetName() == "" {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetInfo().GetName())
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, "Name is invalid")
	}
	if in.GetInfo().GetUnit() == "" {
		logger.Sugar().Errorw("CreateCoin", "Unit", in.GetInfo().GetUnit())
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, "Unit is invalid")
	}
	switch in.GetInfo().GetENV() {
	case "main":
	case "test":
	default:
		logger.Sugar().Errorw("CreateCoin", "ENV", in.GetInfo().GetENV())
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, "ENV is invalid")
	}

	info, err := coinmgrcli.GetCoinBaseOnly(ctx, &coinmgrpb.Conds{
		Name: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetName(),
		},
		ENV: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetENV(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetInfo().GetName(), "ENV", in.GetInfo().GetENV(), "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info != nil {
		info1, err := coin1.GetCoin(ctx, info.ID)
		if err != nil {
			logger.Sugar().Errorw("CreateCoin", "error", err)
			return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
		}
		return &npool.CreateCoinResponse{
			Info: info1,
		}, nil
	}

	info2, err := coin1.CreateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info2,
	}, nil
}
