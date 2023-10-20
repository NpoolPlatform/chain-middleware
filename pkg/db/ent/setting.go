// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/setting"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Setting is the model entity for the Setting schema.
type Setting struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// FeeCoinTypeID holds the value of the "fee_coin_type_id" field.
	FeeCoinTypeID uuid.UUID `json:"fee_coin_type_id,omitempty"`
	// WithdrawFeeByStableUsd holds the value of the "withdraw_fee_by_stable_usd" field.
	WithdrawFeeByStableUsd bool `json:"withdraw_fee_by_stable_usd,omitempty"`
	// WithdrawFeeAmount holds the value of the "withdraw_fee_amount" field.
	WithdrawFeeAmount decimal.Decimal `json:"withdraw_fee_amount,omitempty"`
	// CollectFeeAmount holds the value of the "collect_fee_amount" field.
	CollectFeeAmount decimal.Decimal `json:"collect_fee_amount,omitempty"`
	// HotWalletFeeAmount holds the value of the "hot_wallet_fee_amount" field.
	HotWalletFeeAmount decimal.Decimal `json:"hot_wallet_fee_amount,omitempty"`
	// LowFeeAmount holds the value of the "low_fee_amount" field.
	LowFeeAmount decimal.Decimal `json:"low_fee_amount,omitempty"`
	// HotLowFeeAmount holds the value of the "hot_low_fee_amount" field.
	HotLowFeeAmount decimal.Decimal `json:"hot_low_fee_amount,omitempty"`
	// HotWalletAccountAmount holds the value of the "hot_wallet_account_amount" field.
	HotWalletAccountAmount decimal.Decimal `json:"hot_wallet_account_amount,omitempty"`
	// PaymentAccountCollectAmount holds the value of the "payment_account_collect_amount" field.
	PaymentAccountCollectAmount decimal.Decimal `json:"payment_account_collect_amount,omitempty"`
	// LeastTransferAmount holds the value of the "least_transfer_amount" field.
	LeastTransferAmount decimal.Decimal `json:"least_transfer_amount,omitempty"`
	// NeedMemo holds the value of the "need_memo" field.
	NeedMemo bool `json:"need_memo,omitempty"`
	// RefreshCurrency holds the value of the "refresh_currency" field.
	RefreshCurrency bool `json:"refresh_currency,omitempty"`
	// CheckNewAddressBalance holds the value of the "check_new_address_balance" field.
	CheckNewAddressBalance bool `json:"check_new_address_balance,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Setting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case setting.FieldWithdrawFeeAmount, setting.FieldCollectFeeAmount, setting.FieldHotWalletFeeAmount, setting.FieldLowFeeAmount, setting.FieldHotLowFeeAmount, setting.FieldHotWalletAccountAmount, setting.FieldPaymentAccountCollectAmount, setting.FieldLeastTransferAmount:
			values[i] = new(decimal.Decimal)
		case setting.FieldWithdrawFeeByStableUsd, setting.FieldNeedMemo, setting.FieldRefreshCurrency, setting.FieldCheckNewAddressBalance:
			values[i] = new(sql.NullBool)
		case setting.FieldID, setting.FieldCreatedAt, setting.FieldUpdatedAt, setting.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case setting.FieldEntID, setting.FieldCoinTypeID, setting.FieldFeeCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Setting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Setting fields.
func (s *Setting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case setting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint32(value.Int64)
		case setting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = uint32(value.Int64)
			}
		case setting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = uint32(value.Int64)
			}
		case setting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				s.DeletedAt = uint32(value.Int64)
			}
		case setting.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				s.EntID = *value
			}
		case setting.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				s.CoinTypeID = *value
			}
		case setting.FieldFeeCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fee_coin_type_id", values[i])
			} else if value != nil {
				s.FeeCoinTypeID = *value
			}
		case setting.FieldWithdrawFeeByStableUsd:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_fee_by_stable_usd", values[i])
			} else if value.Valid {
				s.WithdrawFeeByStableUsd = value.Bool
			}
		case setting.FieldWithdrawFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_fee_amount", values[i])
			} else if value != nil {
				s.WithdrawFeeAmount = *value
			}
		case setting.FieldCollectFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field collect_fee_amount", values[i])
			} else if value != nil {
				s.CollectFeeAmount = *value
			}
		case setting.FieldHotWalletFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field hot_wallet_fee_amount", values[i])
			} else if value != nil {
				s.HotWalletFeeAmount = *value
			}
		case setting.FieldLowFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field low_fee_amount", values[i])
			} else if value != nil {
				s.LowFeeAmount = *value
			}
		case setting.FieldHotLowFeeAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field hot_low_fee_amount", values[i])
			} else if value != nil {
				s.HotLowFeeAmount = *value
			}
		case setting.FieldHotWalletAccountAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field hot_wallet_account_amount", values[i])
			} else if value != nil {
				s.HotWalletAccountAmount = *value
			}
		case setting.FieldPaymentAccountCollectAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_account_collect_amount", values[i])
			} else if value != nil {
				s.PaymentAccountCollectAmount = *value
			}
		case setting.FieldLeastTransferAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field least_transfer_amount", values[i])
			} else if value != nil {
				s.LeastTransferAmount = *value
			}
		case setting.FieldNeedMemo:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field need_memo", values[i])
			} else if value.Valid {
				s.NeedMemo = value.Bool
			}
		case setting.FieldRefreshCurrency:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_currency", values[i])
			} else if value.Valid {
				s.RefreshCurrency = value.Bool
			}
		case setting.FieldCheckNewAddressBalance:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field check_new_address_balance", values[i])
			} else if value.Valid {
				s.CheckNewAddressBalance = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Setting.
// Note that you need to call Setting.Unwrap() before calling this method if this Setting
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Setting) Update() *SettingUpdateOne {
	return (&SettingClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Setting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Setting) Unwrap() *Setting {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Setting is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Setting) String() string {
	var builder strings.Builder
	builder.WriteString("Setting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", s.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", s.EntID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("fee_coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", s.FeeCoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("withdraw_fee_by_stable_usd=")
	builder.WriteString(fmt.Sprintf("%v", s.WithdrawFeeByStableUsd))
	builder.WriteString(", ")
	builder.WriteString("withdraw_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.WithdrawFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("collect_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.CollectFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("hot_wallet_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.HotWalletFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("low_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.LowFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("hot_low_fee_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.HotLowFeeAmount))
	builder.WriteString(", ")
	builder.WriteString("hot_wallet_account_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.HotWalletAccountAmount))
	builder.WriteString(", ")
	builder.WriteString("payment_account_collect_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.PaymentAccountCollectAmount))
	builder.WriteString(", ")
	builder.WriteString("least_transfer_amount=")
	builder.WriteString(fmt.Sprintf("%v", s.LeastTransferAmount))
	builder.WriteString(", ")
	builder.WriteString("need_memo=")
	builder.WriteString(fmt.Sprintf("%v", s.NeedMemo))
	builder.WriteString(", ")
	builder.WriteString("refresh_currency=")
	builder.WriteString(fmt.Sprintf("%v", s.RefreshCurrency))
	builder.WriteString(", ")
	builder.WriteString("check_new_address_balance=")
	builder.WriteString(fmt.Sprintf("%v", s.CheckNewAddressBalance))
	builder.WriteByte(')')
	return builder.String()
}

// Settings is a parsable slice of Setting.
type Settings []*Setting

func (s Settings) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
