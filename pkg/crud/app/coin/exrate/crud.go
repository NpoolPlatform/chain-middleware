package exrate

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entexrate "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/exchangerate"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	CoinTypeID    *uuid.UUID
	MarketValue   *decimal.Decimal
	SettlePercent *uint32
	SettleTips    []string
	Setter        *uuid.UUID
	DeletedAt     *uint32
}

func CreateSet(c *ent.ExchangeRateCreate, req *Req) *ent.ExchangeRateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.MarketValue != nil {
		c.SetMarketValue(*req.MarketValue)
		settleValue := *req.MarketValue
		if req.SettlePercent != nil {
			settleValue = settleValue.Mul(decimal.NewFromInt(int64(*req.SettlePercent)))
			settleValue = settleValue.Div(decimal.NewFromInt(100)) //nolint
		}
		c.SetSettleValue(settleValue)
	}

	if req.SettlePercent != nil {
		c.SetSettlePercent(*req.SettlePercent)
	}
	if len(req.SettleTips) > 0 {
		c.SetSettleTips(req.SettleTips)
	}
	if req.Setter != nil {
		c.SetSetter(*req.Setter)
	}
	return c
}

func UpdateSet(u *ent.ExchangeRateUpdateOne, req *Req) *ent.ExchangeRateUpdateOne {
	settlePercent, _ := u.Mutation().SettlePercent()
	marketValue, _ := u.Mutation().MarketValue()

	if req.MarketValue != nil {
		u = u.SetMarketValue(*req.MarketValue)
		marketValue = *req.MarketValue
	}
	if req.SettlePercent != nil {
		u = u.SetSettlePercent(*req.SettlePercent)
		settlePercent = *req.SettlePercent
	}

	if len(req.SettleTips) > 0 {
		u = u.SetSettleTips(req.SettleTips)
	}

	settleValue, _ := u.Mutation().SettleValue() //nolint
	settleValue = marketValue.Mul(decimal.NewFromInt(int64(settlePercent)))
	settleValue = settleValue.Div(decimal.NewFromInt(100)) //nolint
	u = u.SetSettleValue(settleValue)

	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.ExchangeRateQuery, conds *Conds) (*ent.ExchangeRateQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entexrate.ID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entexrate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entexrate.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	q.Where(entexrate.DeletedAt(0))
	return q, nil
}
