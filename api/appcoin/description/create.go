package description

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commonpb "github.com/NpoolPlatform/message/npool"
	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	coindescmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/appcoin/description"
	coinbasemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin/description"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func ValidateCreate(ctx context.Context, in *descmgrpb.CoinDescriptionReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoinDescription", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "AppID", in.GetAppID(), "error", err)
		return err
	}
	exist, err := coinbasemgrcli.ExistCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if !exist {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetCoinTypeID(), "exist", exist)
		return fmt.Errorf("cointypeid not exist")
	}
	switch in.GetUsedFor() {
	case descmgrpb.UsedFor_ProductPage:
	default:
		logger.Sugar().Errorw("CreateCoinDescription", "UsedFor", in.GetUsedFor())
		return fmt.Errorf("usedfor is invalid")
	}
	if in.GetTitle() == "" {
		logger.Sugar().Errorw("CreateCoinDescription", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetMessage() == "" {
		logger.Sugar().Errorw("CreateCoinDescription", "Message", in.GetMessage())
		return fmt.Errorf("message is invalid")
	}
	exist, err = coindescmgrcli.ExistCoinDescriptionConds(ctx, &descmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		CoinTypeID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetCoinTypeID(),
		},
		UsedFor: &commonpb.Int32Val{
			Op:    cruder.EQ,
			Value: int32(in.GetUsedFor()),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if exist {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetCoinTypeID(), "exist", exist)
		return fmt.Errorf("coindescription exist")
	}

	return nil
}

func (s *Server) CreateCoinDescription(
	ctx context.Context,
	in *npool.CreateCoinDescriptionRequest,
) (
	*npool.CreateCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoinDescription")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "appcoin", "CreateCoinDescription")

	info, err := description1.CreateCoinDescription(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "error", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinDescriptionResponse{
		Info: info,
	}, nil
}
