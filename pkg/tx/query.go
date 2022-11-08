package tx

import (
	"context"
	"fmt"

	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	enttran "github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"

	"github.com/google/uuid"
)

func GetTx(ctx context.Context, id string) (*npool.Tx, error) {
	var infos []*npool.Tx

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return cli.
			Tran.
			Query().
			Where(
				enttran.ID(uuid.MustParse(id)),
			).
			Select(
				enttran.FieldID,
				enttran.FieldCoinTypeID,
				enttran.FieldFromAccountID,
				enttran.FieldToAccountID,
				enttran.FieldAmount,
				enttran.FieldFeeAmount,
				enttran.FieldState,
				enttran.FieldChainTxID,
				enttran.FieldType,
				enttran.FieldExtra,
				enttran.FieldCreatedAt,
				enttran.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(enttran.FieldCoinTypeID),
						t1.C(entcoinbase.FieldID),
					).
					AppendSelect(
						sql.As(t1.C(entcoinbase.FieldName), "coin_name"),
						sql.As(t1.C(entcoinbase.FieldLogo), "coin_logo"),
						sql.As(t1.C(entcoinbase.FieldUnit), "coin_unit"),
						sql.As(t1.C(entcoinbase.FieldEnv), "coin_env"),
					)
			}).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}

	infos = expand(infos)

	return infos[0], nil
}

func expand(infos []*npool.Tx) []*npool.Tx {
	for _, info := range infos {
		info.State = txmgrpb.TxState(txmgrpb.TxState_value[info.StateStr])
		info.Type = txmgrpb.TxType(txmgrpb.TxType_value[info.TypeStr])
	}
	return infos
}
