package tx

import (
	"context"

	txmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/tx"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"
)

func UpdateTx(ctx context.Context, in *txmgrpb.TxReq) (*npool.Tx, error) {
	info, err := txmgrcli.UpdateTx(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetTx(ctx, info.ID)
}
