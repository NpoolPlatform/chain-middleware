package currency

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entfiatcurrency "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uuid.UUID
	FiatID          *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.FiatCurrencyCreate, req *Req) *ent.FiatCurrencyCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
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

func UpdateSet(u *ent.FiatCurrencyUpdateOne, req *Req) *ent.FiatCurrencyUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	ID       *cruder.Cond
	FiatID   *cruder.Cond
	FiatIDs  *cruder.Cond
	FiatName *cruder.Cond
}

func SetQueryConds(q *ent.FiatCurrencyQuery, conds *Conds) (*ent.FiatCurrencyQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfiatcurrency.ID(id))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entfiatcurrency.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.FiatIDs != nil {
		ids, ok := conds.FiatIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatids")
		}
		switch conds.FiatIDs.Op {
		case cruder.IN:
			q.Where(entfiatcurrency.FiatIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	q.Where(entfiatcurrency.DeletedAt(0))
	return q, nil
}
