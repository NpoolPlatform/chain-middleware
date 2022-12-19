//nolint:dupl
package fiat

import (
	"context"
	"fmt"

	fiatcurrency1 "github.com/NpoolPlatform/chain-middleware/pkg/fiat"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	fiatcurrencymgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/fiat/currency"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	"github.com/NpoolPlatform/chain-manager/pkg/client/fiat/currencytype"
	"github.com/shopspring/decimal"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func ValidateCreate(ctx context.Context, in *fiatcurrencymgrpb.FiatCurrencyReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateFiatCurrency", "ID", in.GetID(), "error", err)
			return err
		}
	}

	if _, err := uuid.Parse(in.GetFiatCurrencyTypeID()); err != nil {
		logger.Sugar().Errorw("CreateFiatCurrency", "FiatTypeID", in.GetFiatCurrencyTypeID(), "error", err)
		return err
	}

	_, err := currencytype.GetFiatCurrencyType(ctx, in.GetFiatCurrencyTypeID())
	if err != nil {
		logger.Sugar().Errorw("CreateFiatCurrency", "FiatTypeID", in.GetFiatCurrencyTypeID(), "error", err)
		return err
	}

	lowValue, err := decimal.NewFromString(in.GetMarketValueLow())
	if err != nil {
		logger.Sugar().Errorw("CreateFiatCurrency", "MarketValueLow", in.GetMarketValueLow(), "error", err)
		return err
	}

	highValue, err := decimal.NewFromString(in.GetMarketValueHigh())
	if err != nil {
		logger.Sugar().Errorw("CreateFiatCurrency", "MarketValueHigh", in.GetMarketValueHigh(), "error", err)
		return err
	}

	if highValue.Cmp(lowValue) < 0 {
		return fmt.Errorf("invalid value")
	}

	return nil
}

func ValidateCreates(ctx context.Context, in []*fiatcurrencymgrpb.FiatCurrencyReq) error {
	for _, info := range in {
		if err := ValidateCreate(ctx, info); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) CreateFiatCurrency(ctx context.Context, in *npool.CreateFiatCurrencyRequest) (*npool.CreateFiatCurrencyResponse, error) {
	if err := ValidateCreate(ctx, in.GetInfo()); err != nil {
		return &npool.CreateFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := fiatcurrency1.CreateFiatCurrency(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateFiatCurrency", "error", err)
		return &npool.CreateFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrencyResponse{}, nil
}

func (s *Server) CreateFiatCurrencies(
	ctx context.Context,
	in *npool.CreateFiatCurrenciesRequest,
) (
	*npool.CreateFiatCurrenciesResponse,
	error,
) {
	if err := ValidateCreates(ctx, in.GetInfos()); err != nil {
		return &npool.CreateFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err := fiatcurrency1.CreateFiatCurrencies(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateFiatCurrencies", "error", err)
		return &npool.CreateFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFiatCurrenciesResponse{}, nil
}
