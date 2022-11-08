package description

import (
	"context"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	coinbasemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin/description"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

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

	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoinDescription", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "AppID", in.GetInfo().GetAppID(), "error", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := coinbasemgrcli.ExistCoinBase(ctx, in.GetInfo().GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateCoinDescription", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "exist", exist)
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "CoinTypeID not exist")
	}
	switch in.GetInfo().GetUsedFor() {
	case descmgrpb.UsedFor_ProductPage:
	default:
		logger.Sugar().Errorw("CreateCoinDescription", "UsedFor", in.GetInfo().GetUsedFor())
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}
	if in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("CreateCoinDescription", "Title", in.GetInfo().GetTitle())
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
	}
	if in.GetInfo().GetMessage() == "" {
		logger.Sugar().Errorw("CreateCoinDescription", "Message", in.GetInfo().GetMessage())
		return &npool.CreateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Message is invalid")
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
