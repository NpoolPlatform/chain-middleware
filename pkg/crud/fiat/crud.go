package currencytype

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entfiat "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiat"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID   *uuid.UUID
	Name *string
	Logo *string
}

func CreateSet(c *ent.FiatCreate, req *Req) *ent.FiatCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	return c
}

func UpdateSet(u *ent.FiatUpdateOne, req *Req) *ent.FiatUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	return u
}

type Conds struct {
	ID   *cruder.Cond
	Name *cruder.Cond
}

func SetQueryConds(q *ent.FiatQuery, conds *Conds) (*ent.FiatQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfiat.ID(id))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entfiat.Name(name))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	q.Where(entfiat.DeletedAt(0))
	return q, nil
}
