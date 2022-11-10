package coin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	commonpb "github.com/NpoolPlatform/message/npool"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

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

func ValidateCreate(in *npool.CoinReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoin", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if in.GetName() == "" {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetName())
		return fmt.Errorf("name is invalid")
	}
	if in.GetUnit() == "" {
		logger.Sugar().Errorw("CreateCoin", "Unit", in.GetUnit())
		return fmt.Errorf("unit is invalid")
	}
	switch in.GetENV() {
	case "main":
	case "test":
	default:
		logger.Sugar().Errorw("CreateCoin", "ENV", in.GetENV())
		return fmt.Errorf("env is invalid")
	}

	return nil
}

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

	if err := ValidateCreate(in.GetInfo()); err != nil {
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
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

	span = commontracer.TraceInvoker(span, "coin", "coin", "CreateCoin")

	info2, err := coin1.CreateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info2,
	}, nil
}
