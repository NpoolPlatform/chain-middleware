package coin

import (
	"context"
	"fmt"

	coincrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/coin"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                          *uuid.UUID
	Name                        *string
	Logo                        *string
	Presale                     *bool
	Unit                        *string
	ENV                         *string
	ReservedAmount              *decimal.Decimal
	ForPay                      *bool
	HomePage                    *string
	Specs                       *string
	FeeCoinTypeID               *uuid.UUID
	WithdrawFeeByStableUSD      *bool
	WithdrawFeeAmount           *decimal.Decimal
	CollectFeeAmount            *decimal.Decimal
	HotWalletFeeAmount          *decimal.Decimal
	LowFeeAmount                *decimal.Decimal
	HotLowFeeAmount             *decimal.Decimal
	HotWalletAccountAmount      *decimal.Decimal
	PaymentAccountCollectAmount *decimal.Decimal
	Disabled                    *bool
	StableUSD                   *bool
	LeastTransferAmount         *decimal.Decimal
	NeedMemo                    *bool
	RefreshCurrency             *bool
	Conds                       *coincrud.Conds
	Offset                      int32
	Limit                       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid coinname")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Logo = logo
		return nil
	}
}

func WithPresale(presale *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Presale = presale
		return nil
	}
}

func WithUnit(unit *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid coinunit")
		}
		h.Unit = unit
		return nil
	}
}

func WithENV(env *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if env == nil {
			return nil
		}
		switch *env {
		case "main":
		case "test":
		case "local":
		default:
			return fmt.Errorf("invalid coinenv")
		}
		h.ENV = env
		return nil
	}
}

func WithReservedAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.ReservedAmount = &_amount
		return nil
	}
}

func WithForPay(forPay *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ForPay = forPay
		return nil
	}
}

func WithHomePage(homePage *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.HomePage = homePage
		return nil
	}
}

func WithSpecs(specs *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Specs = specs
		return nil
	}
}

func WithFeeCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FeeCoinTypeID = &_id
		return nil
	}
}

func WithWithdrawFeeByStableUSD(stable *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.WithdrawFeeByStableUSD = stable
		return nil
	}
}

func WithWithdrawFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.WithdrawFeeAmount = &_amount
		return nil
	}
}

func WithCollectFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.CollectFeeAmount = &_amount
		return nil
	}
}

func WithHotWalletFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotWalletFeeAmount = &_amount
		return nil
	}
}

func WithLowFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.LowFeeAmount = &_amount
		return nil
	}
}

func WithHotLowFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotLowFeeAmount = &_amount
		return nil
	}
}

func WithHotWalletAccountAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.HotWalletAccountAmount = &_amount
		return nil
	}
}

func WithPaymentAccountCollectAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.PaymentAccountCollectAmount = &_amount
		return nil
	}
}

func WithDisabled(disabled *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
		return nil
	}
}

func WithStableUSD(stable *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StableUSD = stable
		return nil
	}
}

func WithLeastTransferAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.LeastTransferAmount = &_amount
		return nil
	}
}

func WithNeedMemo(needMemo *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.NeedMemo = needMemo
		return nil
	}
}

func WithRefreshCurrency(refresh *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RefreshCurrency = refresh
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &coincrud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.Presale != nil {
			h.Conds.Presale = &cruder.Cond{
				Op:  conds.GetPresale().GetOp(),
				Val: conds.GetPresale().GetValue(),
			}
		}
		if conds.ENV != nil {
			h.Conds.ENV = &cruder.Cond{
				Op:  conds.GetENV().GetOp(),
				Val: conds.GetENV().GetValue(),
			}
		}
		if conds.ForPay != nil {
			h.Conds.ForPay = &cruder.Cond{
				Op:  conds.GetForPay().GetOp(),
				Val: conds.GetForPay().GetValue(),
			}
		}
		if conds.IDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{
				Op:  conds.GetIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{
				Op:  conds.GetDisabled().GetOp(),
				Val: conds.GetDisabled().GetValue(),
			}
		}
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.Names != nil {
			h.Conds.Names = &cruder.Cond{
				Op:  conds.GetNames().GetOp(),
				Val: conds.GetNames().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Limit = limit
		return nil
	}
}