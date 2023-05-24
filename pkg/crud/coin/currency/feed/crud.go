package currencyfeed

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrencyfeed "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyfeed"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uuid.UUID
	CoinTypeID   *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedCoinName *string
	Disabled     *bool
}

func CreateSet(c *ent.CurrencyFeedCreate, req *Req) *ent.CurrencyFeedCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.FeedCoinName != nil {
		c.SetFeedCoinName(*req.FeedCoinName)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	return c
}

func UpdateSet(u *ent.CurrencyFeedUpdateOne, req *Req) *ent.CurrencyFeedUpdateOne {
	if req.FeedCoinName != nil {
		u.SetFeedCoinName(*req.FeedCoinName)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	Disabled    *cruder.Cond
}

func SetQueryConds(q *ent.CurrencyFeedQuery, conds *Conds) (*ent.CurrencyFeedQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.ID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.Disabled(disabled))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	q.Where(entcurrencyfeed.DeletedAt(0))
	return q, nil
}
