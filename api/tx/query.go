//nolint:nolintlint,dupl
package tran

import (
	"context"

	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	constant1 "github.com/NpoolPlatform/chain-middleware/pkg/const"
	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	tx1 "github.com/NpoolPlatform/chain-middleware/pkg/tx"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) GetTx(
	ctx context.Context,
	in *npool.GetTxRequest,
) (
	*npool.GetTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetTx", "ID", in.GetID(), "error", err)
		return &npool.GetTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	info, err := tx1.GetTx(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetTx", "ID", in.GetID(), "error", err)
		return &npool.GetTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTxs(
	ctx context.Context,
	in *npool.GetTxsRequest,
) (
	*npool.GetTxsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	conds := in.GetConds()
	if conds == nil {
		conds = &txmgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetTxs", "ID", conds.GetID().GetValue(), "error", err)
			return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetTxs", "CoinTypeID", conds.GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.AccountID != nil {
		if _, err := uuid.Parse(conds.GetAccountID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetTxs", "AccountID", conds.GetAccountID().GetValue(), "error", err)
			return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	for _, id := range conds.GetAccountIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetTxs", "AccountIDs", conds.GetAccountIDs().GetValue(), "error", err)
			return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "QueryJoin")

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultRowLimit
	}

	infos, total, err := tx1.GetTxs(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetTxs", "Conds", conds, "error", err)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
