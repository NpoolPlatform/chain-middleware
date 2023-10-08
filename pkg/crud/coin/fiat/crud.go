package coinfiat

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcoinfiat "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinfiat"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	CoinTypeID *uuid.UUID
	FiatID     *uuid.UUID
	FeedType   *basetypes.CurrencyFeedType
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinFiatCreate, req *Req) *ent.CoinFiatCreate {
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	return c
}

func UpdateSet(u *ent.CoinFiatUpdateOne, req *Req) *ent.CoinFiatUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	FiatID      *cruder.Cond
}

func SetQueryConds(q *ent.CoinFiatQuery, conds *Conds) (*ent.CoinFiatQuery, error) {
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoinfiat.CoinTypeID(id))
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
		case cruder.EQ:
			q.Where(entcoinfiat.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entcoinfiat.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	q.Where(entcoinfiat.DeletedAt(0))
	return q, nil
}
