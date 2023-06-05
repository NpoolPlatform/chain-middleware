package description

import (
	"context"

	descriptioncrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/app/coin/description"

	"github.com/NpoolPlatform/chain-middleware/pkg/db"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
)

func (h *Handler) ExistCoinDescriptionConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := descriptioncrud.SetQueryConds(cli.CoinDescription.Query(), h.Conds)
		if err != nil {
			return err
		}
		if exist, err = stm.Exist(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
