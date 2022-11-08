package description

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/appcoin/description"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/chain-manager/pkg/db"
	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"

	entcoinbase "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinbase"
	entcoindesc "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coindescription"

	"github.com/google/uuid"
)

func GetCoinDescription(ctx context.Context, id string) (*npool.CoinDescription, error) {
	var infos []*npool.CoinDescription

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return cli.
			CoinDescription.
			Query().
			Where(
				entcoindesc.ID(uuid.MustParse(id)),
			).
			Select(
				entcoindesc.FieldID,
				entcoindesc.FieldAppID,
				entcoindesc.FieldCoinTypeID,
				entcoindesc.FieldUsedFor,
				entcoindesc.FieldTitle,
				entcoindesc.FieldMessage,
				entcoindesc.FieldCreatedAt,
				entcoindesc.FieldUpdatedAt,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entcoinbase.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entcoindesc.FieldCoinTypeID),
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

	return infos[0], nil
}
