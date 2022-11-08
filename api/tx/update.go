package tran

import (
	"context"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	tx1 "github.com/NpoolPlatform/chain-middleware/pkg/tx"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) UpdateTx(ctx context.Context, in *npool.UpdateTxRequest) (*npool.UpdateTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateTx", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	switch in.GetInfo().GetState() {
	case txmgrpb.TxState_StateCreated:
	case txmgrpb.TxState_StateWait:
	case txmgrpb.TxState_StateTransferring:
	case txmgrpb.TxState_StateSuccessful:
	case txmgrpb.TxState_StateFail:
	default:
		logger.Sugar().Errorw("UpdateTx", "State", in.GetInfo().GetState(), "error", "State is invalid")
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, "State is invalid")
	}
	span = commontracer.TraceInvoker(span, "tx", "tx", "UpdateTx")

	info, err := tx1.UpdateTx(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTx", "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: info,
	}, nil
}
