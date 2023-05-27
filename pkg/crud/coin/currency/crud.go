package currency

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currency"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uuid.UUID
	CoinTypeID      *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.CurrencyCreate, req *Req) *ent.CurrencyCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.MarketValueHigh != nil {
		c.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		c.SetMarketValueLow(*req.MarketValueLow)
	}
	return c
}

func UpdateSet(u *ent.CurrencyUpdateOne, req *Req) *ent.CurrencyUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	ID          *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

func SetQueryConds(q *ent.CurrencyQuery, conds *Conds) (*ent.CurrencyQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcurrency.ID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcurrency.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entcurrency.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	q.Where(entcurrency.DeletedAt(0))
	return q, nil
}
