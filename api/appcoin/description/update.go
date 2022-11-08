package description

import (
	"context"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin/description"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) UpdateCoinDescription(
	ctx context.Context,
	in *npool.UpdateCoinDescriptionRequest,
) (
	*npool.UpdateCoinDescriptionResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoinDescription")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCoinDescription", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Title", in.GetInfo().GetTitle())
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
	}
	if in.GetInfo().GetMessage() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Message", in.GetInfo().GetMessage())
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, "Message is invalid")
	}

	span = commontracer.TraceInvoker(span, "appcoin", "appcoin", "UpdateCoinDescription")

	info, err := description1.UpdateCoinDescription(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoinDescription", "error", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: info,
	}, nil
}
