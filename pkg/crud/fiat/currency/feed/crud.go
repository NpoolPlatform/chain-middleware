package currencyfeed

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entcurrencyfeed "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrencyfeed"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID           *uuid.UUID
	FiatID       *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedFiatName *string
	Disabled     *bool
}

func CreateSet(c *ent.FiatCurrencyFeedCreate, req *Req) *ent.FiatCurrencyFeedCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.FeedFiatName != nil {
		c.SetFeedFiatName(*req.FeedFiatName)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	return c
}

func UpdateSet(u *ent.FiatCurrencyFeedUpdateOne, req *Req) *ent.FiatCurrencyFeedUpdateOne {
	if req.FeedFiatName != nil {
		u.SetFeedFiatName(*req.FeedFiatName)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	FiatID   *cruder.Cond
	FiatIDs  *cruder.Cond
	Disabled *cruder.Cond
	FeedType *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.FiatCurrencyFeedQuery, conds *Conds) (*ent.FiatCurrencyFeedQuery, error) {
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
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.FiatIDs != nil {
		ids, ok := conds.FiatIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatids")
		}
		switch conds.FiatIDs.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.FiatIDIn(ids...))
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
	if conds.FeedType != nil {
		feedType, ok := conds.FeedType.Val.(basetypes.CurrencyFeedType)
		if !ok {
			return nil, fmt.Errorf("invalid feedtype")
		}
		switch conds.FeedType.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.FeedType(feedType.String()))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	q.Where(entcurrencyfeed.DeletedAt(0))
	return q, nil
}
