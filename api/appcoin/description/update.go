package description

import (
	"context"
	"fmt"

	descmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	description1 "github.com/NpoolPlatform/chain-middleware/pkg/appcoin/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/google/uuid"
)

func ValidateUpdate(in *descmgrpb.CoinDescriptionReq) error {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCoinDescription", "ID", in.GetID(), "error", err)
		return err
	}
	if in.GetTitle() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetMessage() == "" {
		logger.Sugar().Errorw("UpdateCoinDescription", "Message", in.GetMessage())
		return fmt.Errorf("message is invalid")
	}
	return nil
}

func (s *Server) UpdateCoinDescription(
	ctx context.Context,
	in *npool.UpdateCoinDescriptionRequest,
) (
	*npool.UpdateCoinDescriptionResponse,
	error,
) {
	var err error

	if err := ValidateUpdate(in.GetInfo()); err != nil {
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := description1.UpdateCoinDescription(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateCoinDescription", "error", err)
		return &npool.UpdateCoinDescriptionResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinDescriptionResponse{
		Info: info,
	}, nil
}
