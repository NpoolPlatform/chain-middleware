package appcoin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	coinbasemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func ValidateUpdate(ctx context.Context, in *npool.CoinReq) error { //nolint
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCoin", "ID", in.GetID(), "error", err)
		return err
	}
	info, err := coinbasemgrcli.GetCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}
	if !info.ForPay && in.GetForPay() {
		logger.Sugar().Errorw("UpdateCoin", "ForPay", in.GetForPay(), "CoinForPay", info.ForPay)
		return fmt.Errorf("coin is not payable")
	}
	if info.Disabled && !in.GetDisabled() {
		logger.Sugar().Errorw("UpdateCoin", "Disabled", in.GetDisabled(), "CoinDisabled", info.Disabled)
		return fmt.Errorf("coin is not payable")
	}
	if in.WithdrawAutoReviewAmount != nil {
		amount, err := decimal.NewFromString(in.GetWithdrawAutoReviewAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "WithdrawAutoReviewAmount", in.GetWithdrawAutoReviewAmount(), "error", err)
			return fmt.Errorf("WithdrawAutoReviewAmount is invalid: %v", err)
		}
	}
	if in.DailyRewardAmount != nil {
		amount, err := decimal.NewFromString(in.GetDailyRewardAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "DailyRewardAmount", in.GetDailyRewardAmount(), "error", err)
			return fmt.Errorf("DailyRewardAmount is invalid: %v", err)
		}
	}
	if in.MarketValue != nil {
		amount, err := decimal.NewFromString(in.GetMarketValue())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "MarketValue", in.GetMarketValue(), "error", err)
			return fmt.Errorf("MarketValue is invalid: %v", err)
		}
	}
	if in.SettlePercent != nil && (in.GetSettlePercent() > 100 || in.GetSettlePercent() <= 0) {
		logger.Sugar().Errorw("UpdateCoin", "SettlePercent", in.GetSettlePercent(), "error", "SettlePercent is invalid")
		return fmt.Errorf("settlepercent is invalid")
	}
	if _, err := uuid.Parse(in.GetSetter()); err != nil {
		logger.Sugar().Errorw("UpdateCoin", "Setter", in.GetSetter(), "error", err)
		return err
	}
	if in.Name != nil && in.GetName() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Name", in.GetName(), "error", "Name is invalid")
		return fmt.Errorf("name is invalid")
	}
	if in.Logo != nil && in.GetLogo() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Logo", in.GetLogo(), "error", "Logo is invalid")
		return fmt.Errorf("logo is invalid")
	}
	if in.MaxAmountPerWithdraw != nil {
		amount, err := decimal.NewFromString(in.GetMaxAmountPerWithdraw())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "MaxAmountPerWithdraw", in.GetMaxAmountPerWithdraw(), "error", err)
			return fmt.Errorf("MaxAmountPerWithdraw is invalid: %v", err)
		}
	}

	return nil
}

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	var err error

	if err := ValidateUpdate(ctx, in.GetInfo()); err != nil {
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info1, err := appcoin1.UpdateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info1,
	}, nil
}
