// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinfiat"
	"github.com/google/uuid"
)

// CoinFiat is the model entity for the CoinFiat schema.
type CoinFiat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinFiat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinfiat.FieldID, coinfiat.FieldCreatedAt, coinfiat.FieldUpdatedAt, coinfiat.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case coinfiat.FieldFeedType:
			values[i] = new(sql.NullString)
		case coinfiat.FieldEntID, coinfiat.FieldCoinTypeID, coinfiat.FieldFiatID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinFiat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinFiat fields.
func (cf *CoinFiat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinfiat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cf.ID = int(value.Int64)
		case coinfiat.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cf.CreatedAt = uint32(value.Int64)
			}
		case coinfiat.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cf.UpdatedAt = uint32(value.Int64)
			}
		case coinfiat.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cf.DeletedAt = uint32(value.Int64)
			}
		case coinfiat.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				cf.EntID = *value
			}
		case coinfiat.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cf.CoinTypeID = *value
			}
		case coinfiat.FieldFiatID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fiat_id", values[i])
			} else if value != nil {
				cf.FiatID = *value
			}
		case coinfiat.FieldFeedType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_type", values[i])
			} else if value.Valid {
				cf.FeedType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinFiat.
// Note that you need to call CoinFiat.Unwrap() before calling this method if this CoinFiat
// was returned from a transaction, and the transaction was committed or rolled back.
func (cf *CoinFiat) Update() *CoinFiatUpdateOne {
	return (&CoinFiatClient{config: cf.config}).UpdateOne(cf)
}

// Unwrap unwraps the CoinFiat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cf *CoinFiat) Unwrap() *CoinFiat {
	_tx, ok := cf.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinFiat is not a transactional entity")
	}
	cf.config.driver = _tx.drv
	return cf
}

// String implements the fmt.Stringer.
func (cf *CoinFiat) String() string {
	var builder strings.Builder
	builder.WriteString("CoinFiat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cf.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", cf.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.EntID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("fiat_id=")
	builder.WriteString(fmt.Sprintf("%v", cf.FiatID))
	builder.WriteString(", ")
	builder.WriteString("feed_type=")
	builder.WriteString(cf.FeedType)
	builder.WriteByte(')')
	return builder.String()
}

// CoinFiats is a parsable slice of CoinFiat.
type CoinFiats []*CoinFiat

func (cf CoinFiats) config(cfg config) {
	for _i := range cf {
		cf[_i].config = cfg
	}
}
