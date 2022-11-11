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

func ValidateCreate(ctx context.Context, in *npool.CoinReq) (*npool.CoinReq, error) { //nolint
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoin", "ID", in.GetID(), "error", err)
			return nil, err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateCoin", "AppID", in.GetAppID(), "error", err)
		return nil, err
	}
	info, err := coinbasemgrcli.GetCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return nil, err
	}
	if !info.ForPay && in.GetForPay() {
		logger.Sugar().Errorw("CreateCoin", "ForPay", in.GetForPay(), "CoinForPay", info.ForPay)
		return nil, fmt.Errorf("cointypeid is not payable")
	}
	if in.WithdrawAutoReviewAmount != nil {
		amount, err := decimal.NewFromString(in.GetWithdrawAutoReviewAmount())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("CreateCoin", "WithdrawAutoReviewAmount", in.GetWithdrawAutoReviewAmount(), "error", err)
			return nil, fmt.Errorf("WithdrawAutoReviewAmount is invalid: %v", err)
		}
	}
	if in.MarketValue != nil {
		amount, err := decimal.NewFromString(in.GetMarketValue())
		if err != nil || amount.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Errorw("CreateCoin", "MarketValue", in.GetMarketValue(), "error", err)
			return nil, fmt.Errorf("MarketValue is invalid: %v", err)
		}
	}
	if in.SettlePercent != nil && (in.GetSettlePercent() > 100 || in.GetSettlePercent() <= 0) {
		logger.Sugar().Errorw("CreateCoin", "SettlePercent", in.GetSettlePercent(), "error", "SettlePercent is invalid")
		return nil, fmt.Errorf("settlepercent is invalid")
	}
	if in.Setter != nil {
		if _, err := uuid.Parse(in.GetSetter()); err != nil {
			logger.Sugar().Errorw("CreateCoin", "Setter", in.GetSetter(), "error", err)
			return nil, err
		}
	}

	if in.Name == nil {
		in.Name = &info.Name
	}
	if in.GetName() == "" {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetName(), "error", "Name is invalid")
		return nil, fmt.Errorf("name is invalid")
	}

	if in.Logo == nil {
		in.Logo = &info.Logo
	}

	return in, nil
}

func (s *Server) CreateCoin(ctx context.Context, in *npool.CreateCoinRequest) (*npool.CreateCoinResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateCoin")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	input, err := ValidateCreate(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "appcoin", "appcoin", "CreateCoin")

	info1, err := appcoin1.CreateCoin(ctx, input)
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info1,
	}, nil
}
