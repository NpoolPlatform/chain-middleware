package currencyvalue

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	value1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/value"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/value"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/const"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetCurrency(ctx context.Context, in *npool.GetCurrencyRequest) (*npool.GetCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetCurrency", "ID", in.GetID(), "error", err)
		return &npool.GetCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := value1.GetCurrency(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetCurrency", "ID", in.GetID(), "error", err)
		return &npool.GetCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinCurrency(ctx context.Context, in *npool.GetCoinCurrencyRequest) (*npool.GetCoinCurrencyResponse, error) {
	if _, err := uuid.Parse(in.GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("GetCoinCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.GetCoinCurrencyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := value1.GetCoinCurrency(ctx, in.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorw("GetCoinCurrency", "CoinTypeID", in.GetCoinTypeID(), "error", err)
		return &npool.GetCoinCurrencyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinCurrencyResponse{
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

func (s *Server) GetCurrencies(ctx context.Context, in *npool.GetCurrenciesRequest) (*npool.GetCurrenciesResponse, error) {
	if in.Conds == nil {
		in.Conds = &npool.Conds{}
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCurrenciesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := int32(constant.DefaultRowLimit)
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := value1.GetCurrencies(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetCurrencies", "error", err)
		return &npool.GetCurrenciesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrenciesResponse{
		Infos: infos,
		Total: total,
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

	infos, total, err := value1.GetHistories(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetHistories", "error", err)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
