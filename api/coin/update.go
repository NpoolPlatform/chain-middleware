package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCoin(ctx context.Context, in *npool.UpdateCoinRequest) (*npool.UpdateCoinResponse, error) {
	req := in.GetInfo()
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID),
		coin1.WithName(req.Name),
		coin1.WithUnit(req.Unit),
		coin1.WithLogo(req.Logo),
		coin1.WithReservedAmount(req.ReservedAmount),
		coin1.WithHomePage(req.HomePage),
		coin1.WithSpecs(req.Specs),
		// TODO: this should be get from chain type
		coin1.WithFeeCoinTypeID(req.FeeCoinTypeID),
		coin1.WithWithdrawFeeAmount(req.WithdrawFeeAmount),
		coin1.WithCollectFeeAmount(req.CollectFeeAmount),
		coin1.WithHotWalletFeeAmount(req.HotWalletFeeAmount),
		coin1.WithLowFeeAmount(req.LowFeeAmount),
		coin1.WithHotLowFeeAmount(req.HotLowFeeAmount),
		coin1.WithHotWalletFeeAmount(req.HotWalletFeeAmount),
		coin1.WithPaymentAccountCollectAmount(req.PaymentAccountCollectAmount),
		coin1.WithLeastTransferAmount(req.LeastTransferAmount),
		coin1.WithPresale(req.Presale),
		coin1.WithForPay(req.ForPay),
		coin1.WithDisabled(req.Disabled),
		// TODO: this should be in create from register coin
		coin1.WithStableUSD(req.StableUSD),
		// TODO: this should be in create from register coin
		coin1.WithNeedMemo(req.NeedMemo),
		coin1.WithRefreshCurrency(req.RefreshCurrency),
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
