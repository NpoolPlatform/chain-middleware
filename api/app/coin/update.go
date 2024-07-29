package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(req.ID, true),
		appcoin1.WithName(req.Name, false),
		appcoin1.WithDisplayNames(req.DisplayNames, false),
		appcoin1.WithLogo(req.Logo, false),
		appcoin1.WithForPay(req.ForPay, false),
		appcoin1.WithProductPage(req.ProductPage, false),
		appcoin1.WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount, false),
		appcoin1.WithDailyRewardAmount(req.DailyRewardAmount, false),
		appcoin1.WithDisabled(req.Disabled, false),
		appcoin1.WithDisplay(req.Display, false),
		appcoin1.WithDisplayIndex(req.DisplayIndex, false),
		appcoin1.WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw, false),
		appcoin1.WithMarketValue(req.MarketValue, false),
		appcoin1.WithSettlePercent(req.SettlePercent, false),
		appcoin1.WithSettleTips(req.SettleTips, false),
		appcoin1.WithSetter(req.Setter, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinResponse{
		Info: info,
	}, nil
}
