//nolint:dupl
package currencyvalue

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	value1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/value"

	valuemgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency/value"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	coinmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"

	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func ValidateCreate(ctx context.Context, in *valuemgrpb.CurrencyReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCurrency", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}

	_, err := coinmgrcli.GetCoinBase(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return err
	}

	lowValue, err := decimal.NewFromString(in.GetMarketValueLow())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrency", "MarketValueLow", in.GetMarketValueLow(), "error", err)
		return err
	}

	highValue, err := decimal.NewFromString(in.GetMarketValueHigh())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrency", "MarketValueHigh", in.GetMarketValueHigh(), "error", err)
		return err
	}

	if highValue.Cmp(lowValue) < 0 {
		return fmt.Errorf("invalid value")
	}

	return nil
}

func ValidateCreates(ctx context.Context, in []*valuemgrpb.CurrencyReq) error {
	for _, info := range in {
		if err := ValidateCreate(ctx, info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) CreateCurrency(ctx context.Context, in *npool.CreateCurrencyRequest) (*npool.CreateCurrencyResponse, error) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := value1.CreateCurrency(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrency", "error", err)
		return &npool.CreateCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCurrencies(ctx context.Context, in *npool.CreateCurrenciesRequest) (*npool.CreateCurrenciesResponse, error) {
	if err := ValidateCreates(ctx, in.GetInfos()); err != nil {
		return &npool.CreateCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := value1.CreateCurrencies(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateCurrencies", "error", err)
		return &npool.CreateCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCurrenciesResponse{
		Infos: infos,
	}, nil
}
