package setting

import (
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"
	entsetting "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/setting"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                          *uuid.UUID
	CoinTypeID                  *uuid.UUID
	FeeCoinTypeID               *uuid.UUID
	WithdrawFeeByStableUSD      *bool
	WithdrawFeeAmount           *decimal.Decimal
	CollectFeeAmount            *decimal.Decimal
	HotWalletFeeAmount          *decimal.Decimal
	LowFeeAmount                *decimal.Decimal
	HotLowFeeAmount             *decimal.Decimal
	HotWalletAccountAmount      *decimal.Decimal
	PaymentAccountCollectAmount *decimal.Decimal
	LeastTransferAmount         *decimal.Decimal
	NeedMemo                    *bool
}

func CreateSet(c *ent.SettingCreate, req *Req) *ent.SettingCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FeeCoinTypeID != nil {
		c.SetFeeCoinTypeID(*req.FeeCoinTypeID)
	}
	if req.WithdrawFeeByStableUSD != nil {
		c.SetWithdrawFeeByStableUsd(*req.WithdrawFeeByStableUSD)
	}
	if req.WithdrawFeeAmount != nil {
		c.SetWithdrawFeeAmount(*req.WithdrawFeeAmount)
	}
	if req.CollectFeeAmount != nil {
		c.SetCollectFeeAmount(*req.CollectFeeAmount)
	}
	if req.HotWalletFeeAmount != nil {
		c.SetHotWalletFeeAmount(*req.HotWalletFeeAmount)
	}
	if req.LowFeeAmount != nil {
		c.SetLowFeeAmount(*req.LowFeeAmount)
	}
	if req.HotLowFeeAmount != nil {
		c.SetHotLowFeeAmount(*req.HotLowFeeAmount)
	}
	if req.HotWalletAccountAmount != nil {
		c.SetHotWalletAccountAmount(*req.HotWalletAccountAmount)
	}
	if req.PaymentAccountCollectAmount != nil {
		c.SetPaymentAccountCollectAmount(*req.PaymentAccountCollectAmount)
	}
	if req.LeastTransferAmount != nil {
		c.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.NeedMemo != nil {
		c.SetNeedMemo(*req.NeedMemo)
	}
	return c
}

func UpdateSet(u *ent.SettingUpdateOne, req *Req) *ent.SettingUpdateOne {
	if req.FeeCoinTypeID != nil {
		u.SetFeeCoinTypeID(*req.FeeCoinTypeID)
	}
	if req.WithdrawFeeByStableUSD != nil {
		u.SetWithdrawFeeByStableUsd(*req.WithdrawFeeByStableUSD)
	}
	if req.WithdrawFeeAmount != nil {
		u.SetWithdrawFeeAmount(*req.WithdrawFeeAmount)
	}
	if req.CollectFeeAmount != nil {
		u.SetCollectFeeAmount(*req.CollectFeeAmount)
	}
	if req.HotWalletFeeAmount != nil {
		u.SetHotWalletFeeAmount(*req.HotWalletFeeAmount)
	}
	if req.LowFeeAmount != nil {
		u.SetLowFeeAmount(*req.LowFeeAmount)
	}
	if req.HotLowFeeAmount != nil {
		u.SetHotLowFeeAmount(*req.HotLowFeeAmount)
	}
	if req.HotWalletAccountAmount != nil {
		u.SetHotWalletAccountAmount(*req.HotWalletAccountAmount)
	}
	if req.PaymentAccountCollectAmount != nil {
		u.SetPaymentAccountCollectAmount(*req.PaymentAccountCollectAmount)
	}
	if req.LeastTransferAmount != nil {
		u.SetLeastTransferAmount(*req.LeastTransferAmount)
	}
	if req.NeedMemo != nil {
		u.SetNeedMemo(*req.NeedMemo)
	}

	return u
}

type Conds struct {
	ID            *cruder.Cond
	CoinTypeID    *cruder.Cond
	FeeCoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.SettingQuery, conds *Conds) (*ent.SettingQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entsetting.ID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entsetting.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	if conds.FeeCoinTypeID != nil {
		id, ok := conds.FeeCoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid feecointypeid")
		}
		switch conds.FeeCoinTypeID.Op {
		case cruder.EQ:
			q.Where(entsetting.FeeCoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid setting field")
		}
	}
	q.Where(entsetting.DeletedAt(0))
	return q, nil
}
