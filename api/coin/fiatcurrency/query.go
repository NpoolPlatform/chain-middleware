//nolint:dupl
package fiatcurrency

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	currency1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/fiatcurrency"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/fiatcurrency"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/const"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetFiatCurrency(ctx context.Context, in *npool.GetFiatCurrencyRequest) (*npool.GetFiatCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetFiatCurrency", "ID", in.GetID(), "error", err)
		return &npool.GetFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := currency1.GetFiatCurrency(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetFiatCurrency", "ID", in.GetID(), "error", err)
		return &npool.GetFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinFiatCurrency(ctx context.Context, in *npool.GetCoinFiatCurrencyRequest) (*npool.GetCoinFiatCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("GetCoinFiatCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.GetCoinFiatCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := currency1.GetCoinFiatCurrency(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("GetCoinFiatCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.GetCoinFiatCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinFiatCurrencyResponse{
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
	if conds.CoinTypeID != nil {
		if _, err := uuid.Parse(conds.GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("ValudateConds", "CoinTypeID", conds.GetCoinTypeID().GetValue(), "error", err)
			return err
		}
	}
	for _, id := range conds.GetCoinTypeIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("ValudateConds", "CoinTypeIDs", conds.GetCoinTypeIDs().GetValue(), "error", err)
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

	infos, err := currency1.GetFiatCurrencies(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetFiatCurrencies", "error", err)
		return &npool.GetFiatCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFiatCurrenciesResponse{
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