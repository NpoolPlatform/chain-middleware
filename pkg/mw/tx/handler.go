package tx

import (
	"context"
	"fmt"

	txcrud "github.com/NpoolPlatform/chain-middleware/pkg/crud/tx"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID            *uuid.UUID
	CoinTypeID    *uuid.UUID
	FromAccountID *uuid.UUID
	ToAccountID   *uuid.UUID
	Amount        *decimal.Decimal
	FeeAmount     *decimal.Decimal
	ChainTxID     *string
	State         *basetypes.TxState
	Extra         *string
	Type          *basetypes.TxType
	Reqs          []*txcrud.Req
	Conds         *txcrud.Conds
	Offset        int32
	Limit         int32
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

func WithCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithFromAccountID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FromAccountID = &_id
		return nil
	}
}

func WithToAccountID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ToAccountID = &_id
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Amount = &_amount
		return nil
	}
}

func WithFeeAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.FeeAmount = &_amount
		return nil
	}
}

func WithChainTxID(txID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if txID == nil {
			return nil
		}
		if *txID == "" {
			return fmt.Errorf("invalid txid")
		}
		h.ChainTxID = txID
		return nil
	}
}

func WithState(state *basetypes.TxState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case basetypes.TxState_TxStateCreated:
		case basetypes.TxState_TxStateWait:
		case basetypes.TxState_TxStateTransferring:
		case basetypes.TxState_TxStateSuccessful:
		case basetypes.TxState_TxStateFail:
		default:
			return fmt.Errorf("invalid txstate")
		}
		h.State = state
		return nil
	}
}

func WithExtra(extra *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Extra = extra
		return nil
	}
}

func WithType(_type *basetypes.TxType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case basetypes.TxType_TxWithdraw:
		case basetypes.TxType_TxFeedGas:
		case basetypes.TxType_TxPaymentCollect:
		case basetypes.TxType_TxBenefit:
		case basetypes.TxType_TxLimitation:
		case basetypes.TxType_TxPlatformBenefit:
		case basetypes.TxType_TxUserBenefit:
		default:
			return fmt.Errorf("invalid txtype")
		}
		h.Type = _type
		return nil
	}
}

func WithReqs(reqs []*npool.TxReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*txcrud.Req{}
		for _, req := range reqs {
			_req := &txcrud.Req{
				Extra: req.Extra,
			}
			if req.ID != nil {
				id, err := uuid.Parse(req.GetID())
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.CoinTypeID != nil {
				id, err := uuid.Parse(req.GetCoinTypeID())
				if err != nil {
					return err
				}
				_req.CoinTypeID = &id
			}
			if req.FromAccountID != nil {
				id, err := uuid.Parse(req.GetFromAccountID())
				if err != nil {
					return err
				}
				_req.FromAccountID = &id
			}
			if req.ToAccountID != nil {
				id, err := uuid.Parse(req.GetToAccountID())
				if err != nil {
					return err
				}
				_req.ToAccountID = &id
			}
			if req.Amount != nil {
				amount, err := decimal.NewFromString(req.GetAmount())
				if err != nil {
					return err
				}
				_req.Amount = &amount
			}
			if req.FeeAmount != nil {
				amount, err := decimal.NewFromString(req.GetFeeAmount())
				if err != nil {
					return err
				}
				_req.FeeAmount = &amount
			}
			if req.State != nil {
				switch req.GetState() {
				case basetypes.TxState_TxStateCreated:
				case basetypes.TxState_TxStateWait:
				case basetypes.TxState_TxStateTransferring:
				case basetypes.TxState_TxStateSuccessful:
				case basetypes.TxState_TxStateFail:
				default:
					return fmt.Errorf("invalid txstate")
				}
				_req.State = req.State
			}
			if req.Type != nil {
				switch req.GetType() {
				case basetypes.TxType_TxWithdraw:
				case basetypes.TxType_TxFeedGas:
				case basetypes.TxType_TxPaymentCollect:
				case basetypes.TxType_TxBenefit:
				case basetypes.TxType_TxLimitation:
				case basetypes.TxType_TxPlatformBenefit:
				case basetypes.TxType_TxUserBenefit:
				default:
					return fmt.Errorf("invalid txtype")
				}
				_req.Type = req.Type
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &txcrud.Conds{}
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
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{
				Op:  conds.GetAccountID().GetOp(),
				Val: id,
			}
		}
		if conds.AccountIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAccountIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AccountIDs = &cruder.Cond{
				Op:  conds.GetAccountIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.State != nil {
			h.Conds.State = &cruder.Cond{
				Op:  conds.GetState().GetOp(),
				Val: basetypes.TxState(conds.GetState().GetValue()),
			}
		}
		if conds.Type != nil {
			h.Conds.Type = &cruder.Cond{
				Op:  conds.GetType().GetOp(),
				Val: basetypes.TxType(conds.GetType().GetValue()),
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
