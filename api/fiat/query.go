package fiat

import (
	"context"

	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/fiat"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/const"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetFiatCurrency(ctx context.Context, in *npool.GetFiatCurrencyRequest) (*npool.GetFiatCurrencyResponse, error) {
	info, err := currency1.GetFiatCurrency(ctx, in.FiatCurrencyTypeID)
	if err != nil {
		logger.Sugar().Errorw("GetCoinFiatCurrency", "error", err)
		return &npool.GetFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyResponse{
		Info: info,
	}, nil
}

func ValidateConds(conds *npool.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("ValudateConds", "ID", conds.GetID().GetValue(), "error", err)
			return err
		}
	}
	if conds.FiatCurrencyTypeID != nil {
		if _, err := uuid.Parse(conds.GetFiatCurrencyTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("ValudateConds", "FiatCurrencyTypeID", conds.GetFiatCurrencyTypeID().GetValue(), "error", err)
			return err
		}
	}
	for _, id := range conds.GetFiatCurrencyTypeIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("ValudateConds", "FiatCurrencyTypeIDs", conds.GetFiatCurrencyTypeIDs().GetValue(), "error", err)
			return err
		}
	}
	return nil
}

func (s *Server) GetFiatCurrencies(ctx context.Context, in *npool.GetFiatCurrenciesRequest) (*npool.GetFiatCurrenciesResponse, error) {
	if in.Conds == nil {
		in.Conds = &npool.Conds{}
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := currency1.GetFiatCurrencies(ctx, in.Conds)
	if err != nil {
		logger.Sugar().Errorw("GetCoinFiatCurrency", "error", err)
		return &npool.GetFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrenciesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetCoinFiatCurrencies(
	ctx context.Context,
	in *npool.GetCoinFiatCurrenciesRequest,
) (
	*npool.GetCoinFiatCurrenciesResponse,
	error,
) {
	if len(in.GetCoinTypeIDs()) == 0 {
		logger.Sugar().Errorw("GetFiatCurrencies", "CoinTypeIDs", in.GetCoinTypeIDs())
		return &npool.GetCoinFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, "CoinTypeIDs is empty")
	}
	if len(in.GetFiatCurrencyTypeIDs()) == 0 {
		logger.Sugar().Errorw("GetFiatCurrencies", "FiatCurrencyTypeIDs", in.GetFiatCurrencyTypeIDs())
		return &npool.GetCoinFiatCurrenciesResponse{}, status.Error(codes.InvalidArgument, "CoinTypeIDs is empty")
	}

	infos, err := currency1.GetCoinFiatCurrencies(ctx, in.GetCoinTypeIDs(), in.GetFiatCurrencyTypeIDs())
	if err != nil {
		logger.Sugar().Errorw("GetCoinFiatCurrency", "error", err)
		return &npool.GetCoinFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinFiatCurrenciesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	if in.Conds == nil {
		in.Conds = &npool.Conds{}
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetHistoriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := int32(constant.DefaultRowLimit)
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := currency1.GetHistories(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetHistories", "error", err)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
