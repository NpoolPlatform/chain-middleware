package coin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

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

func ValidateUpdate(ctx context.Context, in *npool.CoinReq) error { //nolint
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCoin", "ID", in.GetID(), "error", err)
		return err
	}
	if in.Name != nil {
		logger.Sugar().Errorw("UpdateCoin", "Name", in.GetName(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.Unit != nil {
		logger.Sugar().Errorw("UpdateCoin", "Unit", in.GetUnit(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.ENV != nil {
		logger.Sugar().Errorw("UpdateCoin", "ENV", in.GetENV(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.Logo != nil && in.GetLogo() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Logo", in.GetLogo(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.ReservedAmount != nil {
		if _, err := decimal.NewFromString(in.GetReservedAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "ReservedAmount", in.GetReservedAmount(), "error", err)
			return err
		}
	}
	if in.HomePage != nil && in.GetHomePage() == "" {
		logger.Sugar().Errorw("UpdateCoin", "HomePage", in.GetHomePage(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.Specs != nil && in.GetSpecs() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Specs", in.GetSpecs(), "error", "permission denied")
		return fmt.Errorf("permission denied")
	}
	if in.FeeCoinTypeID != nil {
		exist, err := basemgrcli.ExistCoinBase(ctx, in.GetFeeCoinTypeID())
		if err != nil {
			logger.Sugar().Errorw("UpdateCoin", "FeeCoinTypeID", in.GetFeeCoinTypeID(), "error", err)
			return err
		}
		if !exist {
			logger.Sugar().Errorw("UpdateCoin", "FeeCoinTypeID", in.GetFeeCoinTypeID(), "error", "not exist")
			return fmt.Errorf("FeeCoinTypeID not exist")
		}
	}
	if in.WithdrawFeeAmount != nil {
		withdrawFeeAmount, err := decimal.NewFromString(in.GetWithdrawFeeAmount())
		if err != nil || withdrawFeeAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateCoin", "WithdrawFeeAmount", in.GetWithdrawFeeAmount(), "error", err)
			return fmt.Errorf("WithdrawFeeAmount is invalid: %v", err)
		}
	}
	if in.CollectFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetCollectFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "CollectFeeAmount", in.GetCollectFeeAmount(), "error", "permission denied")
			return fmt.Errorf("CollectFeeAmount is invalid: %v", err)
		}
	}
	if in.HotWalletFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetHotWalletFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "HotWalletFeeAmount", in.GetHotWalletFeeAmount(), "error", "permission denied")
			return fmt.Errorf("HotWalletFeeAmount is invalid: %v", err)
		}
	}
	if in.LowFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetLowFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "LowFeeAmount", in.GetLowFeeAmount(), "error", "permission denied")
			return fmt.Errorf("LowFeeAmount is invalid: %v", err)
		}
	}
	if in.HotLowFeeAmount != nil {
		if _, err := decimal.NewFromString(in.GetHotLowFeeAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "HotLowFeeAmount", in.GetHotLowFeeAmount(), "error", "permission denied")
			return fmt.Errorf("HotLowFeeAmount is invalid: %v", err)
		}
	}
	if in.HotWalletAccountAmount != nil {
		if _, err := decimal.NewFromString(in.GetHotWalletAccountAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin", "HotWalletAccountAmount", in.GetHotWalletAccountAmount(), "error", "permission denied")
			return err
		}
	}
	if in.PaymentAccountCollectAmount != nil {
		if _, err := decimal.NewFromString(in.GetPaymentAccountCollectAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin",
				"PaymentAccountCollectAmount", in.GetPaymentAccountCollectAmount(),
				"error", "permission denied")
			return err
		}
	}
	if in.LeastTransferAmount != nil {
		if _, err := decimal.NewFromString(in.GetLeastTransferAmount()); err != nil {
			logger.Sugar().Errorw("UpdateCoin",
				"LeastTransferAmount", in.GetLeastTransferAmount(),
				"error", "permission denied")
			return err
		}
	}

	return nil
}

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

	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "coin", "coin", "UpdateCoin")

	info, err := coin1.UpdateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info,
	}, nil
}
