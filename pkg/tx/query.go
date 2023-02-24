package tx

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	txmgrpb "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	crud "github.com/NpoolPlatform/chain-manager/pkg/crud/tx"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	enttran "github.com/NpoolPlatform/chain-manager/pkg/db/ent/tran"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func GetTx(ctx context.Context, id string) (*npool.Tx, error) {
	var infos []*npool.Tx

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Tran.
			Query().
			Where(
				enttran.ID(uuid.MustParse(id)),
			)

		return join(stm).
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

func GetManyTxs(ctx context.Context, ids []string) ([]*npool.Tx, error) {
	var infos []*npool.Tx

	tids := []uuid.UUID{}
	for _, id := range ids {
		tids = append(tids, uuid.MustParse(id))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Tran.
			Query().
			Where(
				enttran.IDIn(tids...),
			)

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}

	infos = expand(infos)

	return infos, nil
}

func GetTxs(ctx context.Context, conds *txmgrpb.Conds, offset, limit int32) ([]*npool.Tx, uint32, error) {
	var infos []*npool.Tx
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)

		stm = stm.
			Order(ent.Desc(enttran.FieldUpdatedAt)).
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	infos = expand(infos)

	return infos, total, nil
}

func join(stm *ent.TranQuery) *ent.TranSelect {
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
		})
}

func expand(infos []*npool.Tx) []*npool.Tx {
	for _, info := range infos {
		info.State = txmgrpb.TxState(txmgrpb.TxState_value[info.StateStr])
		info.Type = basetypes.TxType(basetypes.TxType_value[info.TypeStr])
		info.Amount = decimal.RequireFromString(info.Amount).String()
		info.FeeAmount = decimal.RequireFromString(info.FeeAmount).String()
	}
	return infos
}
