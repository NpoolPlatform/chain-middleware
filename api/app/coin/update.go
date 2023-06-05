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
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(req.ID),
		appcoin1.WithName(req.Name),
		appcoin1.WithDisplayNames(req.DisplayNames),
		appcoin1.WithLogo(req.Logo),
		appcoin1.WithForPay(req.ForPay),
		appcoin1.WithProductPage(req.ProductPage),
		appcoin1.WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount),
		appcoin1.WithDailyRewardAmount(req.DailyRewardAmount),
		appcoin1.WithDisabled(req.Disabled),
		appcoin1.WithDisplay(req.Display),
		appcoin1.WithDisplayIndex(req.DisplayIndex),
		appcoin1.WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw),
		appcoin1.WithMarketValue(req.MarketValue),
		appcoin1.WithSettlePercent(req.SettlePercent),
		appcoin1.WithSettleTips(req.SettleTips),
		appcoin1.WithSetter(req.Setter),
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
