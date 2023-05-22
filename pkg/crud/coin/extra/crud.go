package coinextra

import (
	"fmt"

	"github.com/NpoolPlatform/chain-manager/pkg/db/ent"
	entcoinextra "github.com/NpoolPlatform/chain-manager/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	CoinTypeID *uuid.UUID
	HomePage   *string
	Specs      *string
	StableUSD  *bool
}

func CreateSet(c *ent.CoinExtraCreate, req *Req) *ent.CoinExtraCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.HomePage != nil {
		c.SetHomePage(*req.HomePage)
	}
	if req.Specs != nil {
		c.SetSpecs(*req.Specs)
	}
	if req.StableUSD != nil {
		c.SetStableUsd(*req.StableUSD)
	}
	return c
}

func UpdateSet(u *ent.CoinExtraUpdateOne, req *Req) *ent.CoinExtraUpdateOne {
	if req.HomePage != nil {
		u = u.SetHomePage(*req.HomePage)
	}
	if req.Specs != nil {
		u = u.SetSpecs(*req.Specs)
	}
	if req.StableUSD != nil {
		u = u.SetStableUsd(*req.StableUSD)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	CoinTypeID *cruder.Cond
	StableUSD  *cruder.Cond
}

func SetQueryConds(q *ent.CoinExtraQuery, conds *Conds) (*ent.CoinExtraQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.ID(id),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.CoinTypeID(id),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	if conds.StableUSD != nil {
		stable, ok := conds.StableUSD.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid stableusd")
		}
		switch conds.StableUSD.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.StableUsd(stable),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	return q, nil
}
