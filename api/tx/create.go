package tran

import (
	"context"
	"fmt"

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

func ValidateCreate(ctx context.Context, in *txmgrpb.TxReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateTx", "ID", in.GetID(), "error", err)
			return err
		}
	}
	exist, err := coinbasemgrcli.ExistCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if !exist {
		logger.Sugar().Errorw("CreateTx", "CoinTypeID", in.GetCoinTypeID(), "exist", exist)
		return fmt.Errorf("cointypeid is invalid")
	}
	if _, err := uuid.Parse(in.GetFromAccountID()); err != nil {
		logger.Sugar().Errorw("CreateTx", "FromAccountID", in.GetFromAccountID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetToAccountID()); err != nil {
		logger.Sugar().Errorw("CreateTx", "ToAccountID", in.GetToAccountID(), "error", err)
		return err
	}
	amount, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "Amount", in.GetAmount(), "error", err)
		return err
	}
	feeAmount, err := decimal.NewFromString(in.GetFeeAmount())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "FeeAmount", in.GetFeeAmount(), "error", err)
		return err
	}
	if amount.Cmp(feeAmount) < 0 || feeAmount.Cmp(decimal.NewFromInt(0)) < 0 {
		logger.Sugar().Errorw("CreateTx",
			"FeeAmount", in.GetFeeAmount(),
			"Amount", in.GetAmount(),
			"error", err)
		return fmt.Errorf("amount is invalid")
	}
	switch in.GetState() {
	case txmgrpb.TxState_StateCreated:
	case txmgrpb.TxState_StateWait:
	case txmgrpb.TxState_StateTransferring:
	case txmgrpb.TxState_StateSuccessful:
	case txmgrpb.TxState_StateFail:
	default:
		logger.Sugar().Errorw("CreateTx", "State", in.GetState(), "error", "State is invalid")
		return fmt.Errorf("state is invalid")
	}
	switch in.GetType() {
	case txmgrpb.TxType_TxWithdraw:
	case txmgrpb.TxType_TxFeedGas:
	case txmgrpb.TxType_TxPaymentCollect:
	case txmgrpb.TxType_TxBenefit:
	default:
		logger.Sugar().Errorw("CreateTx", "Type", in.GetType(), "error", "Type is invalid")
		return fmt.Errorf("type is ivnalid")
	}
	return nil
}

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

	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
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
