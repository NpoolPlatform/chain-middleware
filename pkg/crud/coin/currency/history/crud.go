package currencyhistory

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrencyhis "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyhistory"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uuid.UUID
	CoinTypeID      *uuid.UUID
	FeedType        *basetypes.CoinCurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.CurrencyHistoryCreate, req *Req) *ent.CurrencyHistoryCreate {
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

func UpdateSet(u *ent.CurrencyHistoryUpdateOne, req *Req) *ent.CurrencyHistoryUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	ID         *cruder.Cond
	CoinTypeID *cruder.Cond
	StartAt    *cruder.Cond
	EndAt      *cruder.Cond
}

func SetQueryConds(q *ent.CurrencyHistoryQuery, conds *Conds) (*ent.CurrencyHistoryQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcurrencyhis.ID(id))
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
			q.Where(entcurrencyhis.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LTE:
			q.Where(entcurrencyhis.CreatedAtLTE(at))
		case cruder.GTE:
			q.Where(entcurrencyhis.CreatedAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.GTE:
			q.Where(entcurrencyhis.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entcurrencyhis.CreatedAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	q.Where(entcurrencyhis.DeletedAt(0))
	return q, nil
}
