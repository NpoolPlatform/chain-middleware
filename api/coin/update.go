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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateCoin",
			"In", in,
		)
		return &npool.UpdateCoinResponse{}, status.Error(codes.Aborted, "invalid info")
	}
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID, true),
		coin1.WithName(req.Name, false),
		coin1.WithUnit(req.Unit, false),
		coin1.WithLogo(req.Logo, false),
		coin1.WithReservedAmount(req.ReservedAmount, false),
		coin1.WithHomePage(req.HomePage, false),
		coin1.WithSpecs(req.Specs, false),
		// TODO: this should be get from chain type
		coin1.WithFeeCoinTypeID(req.FeeCoinTypeID, false),
		coin1.WithWithdrawFeeByStableUSD(req.WithdrawFeeByStableUSD, false),
		coin1.WithWithdrawFeeAmount(req.WithdrawFeeAmount, false),
		coin1.WithCollectFeeAmount(req.CollectFeeAmount, false),
		coin1.WithHotWalletFeeAmount(req.HotWalletFeeAmount, false),
		coin1.WithLowFeeAmount(req.LowFeeAmount, false),
		coin1.WithHotLowFeeAmount(req.HotLowFeeAmount, false),
		coin1.WithHotWalletFeeAmount(req.HotWalletFeeAmount, false),
		coin1.WithHotWalletAccountAmount(req.HotWalletAccountAmount, false),
		coin1.WithPaymentAccountCollectAmount(req.PaymentAccountCollectAmount, false),
		coin1.WithLeastTransferAmount(req.LeastTransferAmount, false),
		coin1.WithPresale(req.Presale, false),
		coin1.WithForPay(req.ForPay, false),
		coin1.WithDisabled(req.Disabled, false),
		// TODO: this should be in create from register coin
		coin1.WithStableUSD(req.StableUSD, false),
		// TODO: this should be in create from register coin
		coin1.WithNeedMemo(req.NeedMemo, false),
		coin1.WithRefreshCurrency(req.RefreshCurrency, false),
		coin1.WithCheckNewAddressBalance(req.CheckNewAddressBalance, false),
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
