//nolint:nolintlint,dupl
package coin

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	constant1 "github.com/NpoolPlatform/chain-middleware/pkg/const"
	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetCoin", "ID", in.GetID(), "error", err)
		return &npool.GetCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetCoins")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
		logger.Sugar().Errorw("GetCoins", "ID", in.GetConds().GetID().GetValue(), "error", err)
		return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetConds().ENV != nil {
		switch in.GetConds().GetENV().GetValue() {
		case "main", "test":
		default:
			logger.Sugar().Errorw("GetCoins", "ENV", in.GetConds().GetENV().GetValue(), "error", err)
			return &npool.GetCoinsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultRowLimit
	}

	infos, total, err := coin1.GetCoins(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCoins", "ID", in.GetConds(), "error", err)
		return &npool.GetCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
