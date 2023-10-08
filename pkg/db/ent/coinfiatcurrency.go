// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinfiatcurrency"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinFiatCurrency is the model entity for the CoinFiatCurrency schema.
type CoinFiatCurrency struct {
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
	// FiatID holds the value of the "fiat_id" field.
	FiatID uuid.UUID `json:"fiat_id,omitempty"`
	// FeedType holds the value of the "feed_type" field.
	FeedType string `json:"feed_type,omitempty"`
	// MarketValueLow holds the value of the "market_value_low" field.
	MarketValueLow decimal.Decimal `json:"market_value_low,omitempty"`
	// MarketValueHigh holds the value of the "market_value_high" field.
	MarketValueHigh decimal.Decimal `json:"market_value_high,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinFiatCurrency) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinfiatcurrency.FieldMarketValueLow, coinfiatcurrency.FieldMarketValueHigh:
			values[i] = new(decimal.Decimal)
		case coinfiatcurrency.FieldID, coinfiatcurrency.FieldCreatedAt, coinfiatcurrency.FieldUpdatedAt, coinfiatcurrency.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coinfiatcurrency.FieldFeedType:
			values[i] = new(sql.NullString)
		case coinfiatcurrency.FieldEntID, coinfiatcurrency.FieldCoinTypeID, coinfiatcurrency.FieldFiatID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinFiatCurrency", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinFiatCurrency fields.
func (cfc *CoinFiatCurrency) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinfiatcurrency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cfc.ID = uint32(value.Int64)
		case coinfiatcurrency.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cfc.CreatedAt = uint32(value.Int64)
			}
		case coinfiatcurrency.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cfc.UpdatedAt = uint32(value.Int64)
			}
		case coinfiatcurrency.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cfc.DeletedAt = uint32(value.Int64)
			}
		case coinfiatcurrency.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				cfc.EntID = *value
			}
		case coinfiatcurrency.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cfc.CoinTypeID = *value
			}
		case coinfiatcurrency.FieldFiatID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fiat_id", values[i])
			} else if value != nil {
				cfc.FiatID = *value
			}
		case coinfiatcurrency.FieldFeedType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_type", values[i])
			} else if value.Valid {
				cfc.FeedType = value.String
			}
		case coinfiatcurrency.FieldMarketValueLow:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_low", values[i])
			} else if value != nil {
				cfc.MarketValueLow = *value
			}
		case coinfiatcurrency.FieldMarketValueHigh:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field market_value_high", values[i])
			} else if value != nil {
				cfc.MarketValueHigh = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinFiatCurrency.
// Note that you need to call CoinFiatCurrency.Unwrap() before calling this method if this CoinFiatCurrency
// was returned from a transaction, and the transaction was committed or rolled back.
func (cfc *CoinFiatCurrency) Update() *CoinFiatCurrencyUpdateOne {
	return (&CoinFiatCurrencyClient{config: cfc.config}).UpdateOne(cfc)
}

// Unwrap unwraps the CoinFiatCurrency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cfc *CoinFiatCurrency) Unwrap() *CoinFiatCurrency {
	_tx, ok := cfc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinFiatCurrency is not a transactional entity")
	}
	cfc.config.driver = _tx.drv
	return cfc
}

// String implements the fmt.Stringer.
func (cfc *CoinFiatCurrency) String() string {
	var builder strings.Builder
	builder.WriteString("CoinFiatCurrency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cfc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cfc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cfc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cfc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", cfc.EntID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cfc.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("fiat_id=")
	builder.WriteString(fmt.Sprintf("%v", cfc.FiatID))
	builder.WriteString(", ")
	builder.WriteString("feed_type=")
	builder.WriteString(cfc.FeedType)
	builder.WriteString(", ")
	builder.WriteString("market_value_low=")
	builder.WriteString(fmt.Sprintf("%v", cfc.MarketValueLow))
	builder.WriteString(", ")
	builder.WriteString("market_value_high=")
	builder.WriteString(fmt.Sprintf("%v", cfc.MarketValueHigh))
	builder.WriteByte(')')
	return builder.String()
}

// CoinFiatCurrencies is a parsable slice of CoinFiatCurrency.
type CoinFiatCurrencies []*CoinFiatCurrency

func (cfc CoinFiatCurrencies) config(cfg config) {
	for _i := range cfc {
		cfc[_i].config = cfg
	}
}
