package description

import (
	"fmt"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	entcoindescription "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	CoinTypeID *uuid.UUID
	UsedFor    *basetypes.UsedFor
	Title      *string
	Message    *string
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinDescriptionCreate, req *Req) *ent.CoinDescriptionCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.CoinDescriptionUpdateOne, req *Req) *ent.CoinDescriptionUpdateOne {
	if req.Title != nil {
		u = u.SetTitle(*req.Title)
	}
	if req.Message != nil {
		u = u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
	UsedFor    *cruder.Cond
}

func SetQueryConds(q *ent.CoinDescriptionQuery, conds *Conds) (*ent.CoinDescriptionQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.ID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.AppID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoindescription.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(basetypes.UsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entcoindescription.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid entcoindescription field")
		}
	}
	q.Where(entcoindescription.DeletedAt(0))
	return q, nil
}
