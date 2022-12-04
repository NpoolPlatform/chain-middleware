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

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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

func GetTxs(ctx context.Context, conds *txmgrpb.Conds, offset, limit int32) ([]*npool.Tx, uint32, error) {
	var infos []*npool.Tx
	var total uint32

	ids := []uuid.UUID{}
	for _, id := range conds.GetAccountIDs().GetValue() {
		ids = append(ids, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Tran.
			Query()

		if conds.ID != nil {
			stm.Where(
				enttran.ID(uuid.MustParse(conds.GetID().GetValue())),
			)
		}
		if conds.CoinTypeID != nil {
			stm.Where(
				enttran.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())),
			)
		}
		if conds.AccountID != nil {
			stm.Where(
				enttran.Or(
					enttran.FromAccountID(uuid.MustParse(conds.GetAccountID().GetValue())),
					enttran.ToAccountID(uuid.MustParse(conds.GetAccountID().GetValue())),
				),
			)
		}
		if len(ids) > 0 {
			stm.Where(
				enttran.Or(
					enttran.FromAccountIDIn(ids...),
					enttran.ToAccountIDIn(ids...),
				),
			)
		}
		if conds.State != nil {
			switch conds.GetState().GetOp() {
			case cruder.EQ:
				stm.Where(
					enttran.State(txmgrpb.TxState(conds.GetState().GetValue()).String()),
				)
			case cruder.NEQ:
				stm.Where(
					enttran.StateNEQ(txmgrpb.TxState(conds.GetState().GetValue()).String()),
				)
			}
		}
		if conds.Type != nil {
			stm.Where(
				enttran.Type(txmgrpb.TxType(conds.GetType().GetValue()).String()),
			)
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		return stm.
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
			Offset(int(offset)).
			Limit(int(limit)).
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
		return nil, 0, err
	}

	infos = expand(infos)

	return infos, total, nil
}

func expand(infos []*npool.Tx) []*npool.Tx {
	for _, info := range infos {
		info.State = txmgrpb.TxState(txmgrpb.TxState_value[info.StateStr])
		info.Type = txmgrpb.TxType(txmgrpb.TxType_value[info.TypeStr])
	}
	return infos
}
