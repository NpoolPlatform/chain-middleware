package coin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	commonpb "github.com/NpoolPlatform/message/npool"

	coinmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/coin/base"
	coinmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/base"

	coin1 "github.com/NpoolPlatform/chain-middleware/pkg/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

//nolint
func ValidateCreate(in *npool.CoinReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("CreateCoin", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if in.GetName() == "" {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetName())
		return fmt.Errorf("name is invalid")
	}
	if in.GetUnit() == "" {
		logger.Sugar().Errorw("CreateCoin", "Unit", in.GetUnit())
		return fmt.Errorf("unit is invalid")
	}
	switch in.GetENV() {
	case "main":
	case "test":
	default:
		logger.Sugar().Errorw("CreateCoin", "ENV", in.GetENV())
		return fmt.Errorf("env is invalid")
	}

	return nil
}

func (s *Server) CreateCoin(
	ctx context.Context,
	in *npool.CreateCoinRequest,
) (
	*npool.CreateCoinResponse,
	error,
) {
	var err error

	if err := ValidateCreate(in.GetInfo()); err != nil {
		return &npool.CreateCoinResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := coinmgrcli.GetCoinBaseOnly(ctx, &coinmgrpb.Conds{
		Name: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetName(),
		},
		ENV: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetENV(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "Name", in.GetInfo().GetName(), "ENV", in.GetInfo().GetENV(), "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info != nil {
		info1, err := coin1.GetCoin(ctx, info.ID)
		if err != nil {
			logger.Sugar().Errorw("CreateCoin", "error", err)
			return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
		}
		return &npool.CreateCoinResponse{
			Info: info1,
		}, nil
	}

	info2, err := coin1.CreateCoin(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateCoin", "error", err)
		return &npool.CreateCoinResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCoinResponse{
		Info: info2,
	}, nil
}
