package coin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	basemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) UpdateCoin(
	ctx context.Context,
	in *npool.UpdateCoinRequest,
) (
	*npool.UpdateCoinResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCoin", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Name != nil {
		logger.Sugar().Errorw("UpdateCoin", "Name", in.GetInfo().GetName(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Unit != nil {
		logger.Sugar().Errorw("UpdateCoin", "Unit", in.GetInfo().GetUnit(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().ENV != nil {
		logger.Sugar().Errorw("UpdateCoin", "ENV", in.GetInfo().GetENV(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Logo != nil && in.GetInfo().GetLogo() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Logo", in.GetInfo().GetLogo(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().ReservedAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetReservedAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "ReservedAmount", in.GetInfo().GetReservedAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetInfo().HomePage != nil && in.GetInfo().GetHomePage() == "" {
		logger.Sugar().Errorw("UpdateCoin", "HomePage", in.GetInfo().GetHomePage(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Specs != nil && in.GetInfo().GetSpecs() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Specs", in.GetInfo().GetSpecs(), "error", "permission denied")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().FeeCoinTypeID != nil {
		exist, err := basemgrcli.ExistCoinBase(ctx, in.GetInfo().GetFeeCoinTypeID())
		if err != nil {
			logger.Sugar().Errorw("UpdateCoin", "FeeCoinTypeID", in.GetInfo().GetFeeCoinTypeID(), "error", err)
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if !exist {
			logger.Sugar().Errorw("UpdateCoin", "FeeCoinTypeID", in.GetInfo().GetFeeCoinTypeID(), "error", "not exist")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, "FeeCoinTypeID not exist")
		}
	}
	if in.GetInfo().WithdrawFeeAmount != nil {
		withdrawFeeAmount, err := decimal.NewFromString(in.GetInfo().GetWithdrawFeeAmount())
		if err != nil || withdrawFeeAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateCoin", "WithdrawFeeAmount", in.GetInfo().GetWithdrawFeeAmount(), "error", err)
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("WithdrawFeeAmount is invalid: %v", err))
		}
	}
	if in.GetInfo().CollectFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetCollectFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "CollectFeeAmount", in.GetInfo().GetCollectFeeAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("CollectFeeAmount is invalid: %v", err))
		}
	}
	if in.GetInfo().HotWalletFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetHotWalletFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "HotWalletFeeAmount", in.GetInfo().GetHotWalletFeeAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("HotWalletFeeAmount is invalid: %v", err))
		}
	}
	if in.GetInfo().LowFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetLowFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "LowFeeAmount", in.GetInfo().GetLowFeeAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("LowFeeAmount is invalid: %v", err))
		}
	}
	if in.GetInfo().HotWalletAccountAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetHotWalletAccountAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "HotWalletAccountAmount", in.GetInfo().GetHotWalletAccountAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("HotWalletAccountAmount is invalid: %v", err))
		}
	}
	if in.GetInfo().PaymentAccountCollectAmount != nil {
		if _, err := decimal.NewFromString(in.GetInfo().GetPaymentAccountCollectAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "PaymentAccountCollectAmount", in.GetInfo().GetPaymentAccountCollectAmount(), "error", "permission denied")
			return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("PaymentAccountCollectAmount is invalid: %v", err))
		}
	}

	info, err := coin1.UpdateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info,
	}, nil
}
