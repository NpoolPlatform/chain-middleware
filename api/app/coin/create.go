package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/app/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoin(ctx context.Context, in *npool.CreateCoinRequest) (*npool.CreateCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := appcoin1.NewHandler(
		ctx,
		appcoin1.WithID(req.ID, false),
		appcoin1.WithEntID(req.EntID, false),
		appcoin1.WithAppID(req.AppID, true),
		appcoin1.WithCoinTypeID(req.CoinTypeID, true),
		appcoin1.WithName(req.Name, true),
		appcoin1.WithDisplayNames(req.DisplayNames, true),
		appcoin1.WithLogo(req.Logo, true),
		appcoin1.WithForPay(req.ForPay, true),
		appcoin1.WithProductPage(req.ProductPage, true),
		appcoin1.WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount, true),
		appcoin1.WithDailyRewardAmount(req.DailyRewardAmount, true),
		appcoin1.WithDisplay(req.Display, true),
		appcoin1.WithDisplayIndex(req.DisplayIndex, true),
		appcoin1.WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw, true),
		appcoin1.WithMarketValue(req.MarketValue, true),
		appcoin1.WithSettlePercent(req.SettlePercent, true),
		appcoin1.WithSettleTips(req.SettleTips, true),
		appcoin1.WithSetter(req.Setter, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info,
	}, nil
}
