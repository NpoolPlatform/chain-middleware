package currencyfeed

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	feed1 "github.com/NpoolPlatform/chain-middleware/pkg/coin/currency/feed"

	constant "github.com/NpoolPlatform/chain-middleware/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency/feed"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

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

func (s *Server) GetCurrencyFeeds(ctx context.Context, in *npool.GetCurrencyFeedsRequest) (*npool.GetCurrencyFeedsResponse, error) {
	if in.Conds == nil {
		in.Conds = &npool.Conds{}
	}

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetCurrencyFeedsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	limit := int32(constant.DefaultRowLimit)
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := feed1.GetCurrencyFeeds(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCurrencyFeeds", "error", err)
		return &npool.GetCurrencyFeedsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCurrencyFeedsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
