package tx

import (
	"context"

	txmgrcli "github.com/NpoolPlatform/chain-manager/pkg/client/tx"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"
)

func CreateTx(ctx context.Context, in *txmgrpb.TxReq) (*npool.Tx, error) {
	info, err := txmgrcli.CreateTx(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetTx(ctx, info.ID)
}

func CreateTxs(ctx context.Context, in []*txmgrpb.TxReq) ([]*npool.Tx, error) {
	infos, err := txmgrcli.CreateTxs(ctx, in)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	for _, info := range infos {
		ids = append(ids, info.ID)
	}

	return GetManyTxs(ctx, ids)
}
