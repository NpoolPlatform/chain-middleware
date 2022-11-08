package tran

import (
	"context"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	coinbasemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	tx1 "github.com/NpoolPlatform/chain-middleware/pkg/tx"

	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) CreateTx(ctx context.Context, in *npool.CreateTxRequest) (*npool.CreateTxResponse, error) { //nolint
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			logger.Sugar().Errorw("CreateTx", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	exist, err := coinbasemgrcli.ExistCoinBase(ctx, in.GetInfo().GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateTx", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "exist", exist)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, "CoinTypeID not exist")
	}
	if _, err := uuid.Parse(in.GetInfo().GetFromAccountID()); err != nil {
		logger.Sugar().Errorw("CreateTx", "FromAccountID", in.GetInfo().GetFromAccountID(), "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetToAccountID()); err != nil {
		logger.Sugar().Errorw("CreateTx", "ToAccountID", in.GetInfo().GetToAccountID(), "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	amount, err := decimal.NewFromString(in.GetInfo().GetAmount())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "Amount", in.GetInfo().GetAmount(), "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	feeAmount, err := decimal.NewFromString(in.GetInfo().GetFeeAmount())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "FeeAmount", in.GetInfo().GetFeeAmount(), "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if amount.Cmp(feeAmount) < 0 || feeAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("CreateTx",
			"FeeAmount", in.GetInfo().GetFeeAmount(),
			"Amount", in.GetInfo().GetAmount(),
			"error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, "Amount is invalid")
	}
	switch in.GetInfo().GetState() {
	case txmgrpb.TxState_StateCreated:
	case txmgrpb.TxState_StateWait:
	case txmgrpb.TxState_StateTransferring:
	case txmgrpb.TxState_StateSuccessful:
	case txmgrpb.TxState_StateFail:
	default:
		logger.Sugar().Errorw("CreateTx", "State", in.GetInfo().GetState(), "error", "State is invalid")
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, "State is invalid")
	}
	switch in.GetInfo().GetType() {
	case txmgrpb.TxType_TxWithdraw:
	case txmgrpb.TxType_TxFeedGas:
	case txmgrpb.TxType_TxPaymentCollect:
	case txmgrpb.TxType_TxBenefit:
	default:
		logger.Sugar().Errorw("CreateTx", "Type", in.GetInfo().GetType(), "error", "Type is invalid")
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, "Type is invalid")
	}

	span = commontracer.TraceInvoker(span, "tx", "tx", "CreateTx")

	info, err := tx1.CreateTx(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: info,
	}, nil
}
