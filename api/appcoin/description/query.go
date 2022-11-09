//nolint:nolintlint,dupl
package description

import (
	"context"

	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	constant1 "github.com/NpoolPlatform/chain-middleware/pkg/const"
	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin/description"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) GetCoinDescription(
	ctx context.Context,
	in *npool.GetCoinDescriptionRequest,
) (
	*npool.GetCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinDescription")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetCoinDescription", "ID", in.GetID(), "error", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	info, err := description1.GetCoinDescription(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetCoinDescription", "ID", in.GetID(), "error", err)
		return &npool.GetCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinDescriptions(
	ctx context.Context,
	in *npool.GetCoinDescriptionsRequest,
) (
	*npool.GetCoinDescriptionsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoinDescriptions")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	conds := in.GetConds()
	if conds == nil {
		conds = &descmgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCoinDescriptions", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCoinDescriptions", "AppID", conds.GetAppID().GetValue(), "error", err)
			return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetCoinDescriptions", "CoinTypeID", conds.GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.UsedFor != nil {
		switch descmgrpb.UsedFor(conds.GetUsedFor().GetValue()) {
		case descmgrpb.UsedFor_ProductPage:
		default:
			logger.Sugar().Errorw("GetCoinDescriptions", "UsedFor", conds.GetUsedFor().GetValue(), "error", err)
			return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultRowLimit
	}

	infos, total, err := description1.GetCoinDescriptions(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCoinDescriptions", "ID", conds, "error", err)
		return &npool.GetCoinDescriptionsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinDescriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
