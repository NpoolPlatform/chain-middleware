package appcoin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/message/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin"

	commontracer "github.com/NpoolPlatform/chain-middleware/pkg/tracer"

	coinbasemgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin"

	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) { //nolint
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
	info, err := coinbasemgrcli.GetCoinBase(ctx, in.GetInfo().GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !info.ForPay && in.GetInfo().GetForPay() {
		logger.Sugar().Errorw("UpdateCoin", "ForPay", in.GetInfo().GetForPay(), "CoinForPay", info.ForPay)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, "permission denied")
	}
	if in.GetInfo().WithdrawAutoReviewAmount != nil {
		amount, err := decimal.NewFromString(in.GetInfo().GetWithdrawAutoReviewAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "WithdrawAutoReviewAmount", in.GetInfo().GetWithdrawAutoReviewAmount(), "error", err)
			return &npool.UpdateCoinResponse{}, status.Error(
				codes.InvalidArgument,
				fmt.Sprintf("WithdrawAutoReviewAmount is invalid: %v", err),
			)
		}
	}
	if in.GetInfo().MarketValue != nil {
		amount, err := decimal.NewFromString(in.GetInfo().GetMarketValue())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("UpdateCoin", "MarketValue", in.GetInfo().GetMarketValue(), "error", err)
			return &npool.UpdateCoinResponse{}, status.Error(
				codes.InvalidArgument,
				fmt.Sprintf("MarketValue is invalid: %v", err),
			)
		}
	}
	if in.GetInfo().SettlePercent != nil && (in.GetInfo().GetSettlePercent() > 100 || in.GetInfo().GetSettlePercent() <= 0) {
		logger.Sugar().Errorw("UpdateCoin", "SettlePercent", in.GetInfo().GetSettlePercent(), "error", "SettlePercent is invalid")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, "SettlePercent is invalid")
	}
	if _, err := uuid.Parse(in.GetInfo().GetSetter()); err != nil {
		logger.Sugar().Errorw("UpdateCoin", "Setter", in.GetInfo().GetSetter(), "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Name != nil && in.GetInfo().GetName() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Name", in.GetInfo().GetName(), "error", "Name is invalid")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, "Name is invalid")
	}
	if in.GetInfo().Logo != nil && in.GetInfo().GetLogo() == "" {
		logger.Sugar().Errorw("UpdateCoin", "Logo", in.GetInfo().GetLogo(), "error", "Logo is invalid")
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, "Logo is invalid")
	}

	span = commontracer.TraceInvoker(span, "appcoin", "appcoin", "UpdateCoin")

	info1, err := appcoin1.UpdateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoin", "error", err)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info1,
	}, nil
}
