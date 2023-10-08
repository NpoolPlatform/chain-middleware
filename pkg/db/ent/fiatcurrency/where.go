// Code generated by ent, DO NOT EDIT.

package fiatcurrency

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// FiatID applies equality check predicate on the "fiat_id" field. It's identical to FiatIDEQ.
func FiatID(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiatID), v))
	})
}

// FeedType applies equality check predicate on the "feed_type" field. It's identical to FeedTypeEQ.
func FeedType(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedType), v))
	})
}

// MarketValueLow applies equality check predicate on the "market_value_low" field. It's identical to MarketValueLowEQ.
func MarketValueLow(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueHigh applies equality check predicate on the "market_value_high" field. It's identical to MarketValueHighEQ.
func MarketValueHigh(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarketValueHigh), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// FiatIDEQ applies the EQ predicate on the "fiat_id" field.
func FiatIDEQ(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFiatID), v))
	})
}

// FiatIDNEQ applies the NEQ predicate on the "fiat_id" field.
func FiatIDNEQ(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFiatID), v))
	})
}

// FiatIDIn applies the In predicate on the "fiat_id" field.
func FiatIDIn(vs ...uuid.UUID) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFiatID), v...))
	})
}

// FiatIDNotIn applies the NotIn predicate on the "fiat_id" field.
func FiatIDNotIn(vs ...uuid.UUID) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFiatID), v...))
	})
}

// FiatIDGT applies the GT predicate on the "fiat_id" field.
func FiatIDGT(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFiatID), v))
	})
}

// FiatIDGTE applies the GTE predicate on the "fiat_id" field.
func FiatIDGTE(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFiatID), v))
	})
}

// FiatIDLT applies the LT predicate on the "fiat_id" field.
func FiatIDLT(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFiatID), v))
	})
}

// FiatIDLTE applies the LTE predicate on the "fiat_id" field.
func FiatIDLTE(v uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFiatID), v))
	})
}

// FiatIDIsNil applies the IsNil predicate on the "fiat_id" field.
func FiatIDIsNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFiatID)))
	})
}

// FiatIDNotNil applies the NotNil predicate on the "fiat_id" field.
func FiatIDNotNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFiatID)))
	})
}

// FeedTypeEQ applies the EQ predicate on the "feed_type" field.
func FeedTypeEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedType), v))
	})
}

// FeedTypeNEQ applies the NEQ predicate on the "feed_type" field.
func FeedTypeNEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeedType), v))
	})
}

// FeedTypeIn applies the In predicate on the "feed_type" field.
func FeedTypeIn(vs ...string) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeedType), v...))
	})
}

// FeedTypeNotIn applies the NotIn predicate on the "feed_type" field.
func FeedTypeNotIn(vs ...string) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeedType), v...))
	})
}

// FeedTypeGT applies the GT predicate on the "feed_type" field.
func FeedTypeGT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeedType), v))
	})
}

// FeedTypeGTE applies the GTE predicate on the "feed_type" field.
func FeedTypeGTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeedType), v))
	})
}

// FeedTypeLT applies the LT predicate on the "feed_type" field.
func FeedTypeLT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeedType), v))
	})
}

// FeedTypeLTE applies the LTE predicate on the "feed_type" field.
func FeedTypeLTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeedType), v))
	})
}

// FeedTypeContains applies the Contains predicate on the "feed_type" field.
func FeedTypeContains(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFeedType), v))
	})
}

// FeedTypeHasPrefix applies the HasPrefix predicate on the "feed_type" field.
func FeedTypeHasPrefix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFeedType), v))
	})
}

// FeedTypeHasSuffix applies the HasSuffix predicate on the "feed_type" field.
func FeedTypeHasSuffix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFeedType), v))
	})
}

// FeedTypeIsNil applies the IsNil predicate on the "feed_type" field.
func FeedTypeIsNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFeedType)))
	})
}

// FeedTypeNotNil applies the NotNil predicate on the "feed_type" field.
func FeedTypeNotNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFeedType)))
	})
}

// FeedTypeEqualFold applies the EqualFold predicate on the "feed_type" field.
func FeedTypeEqualFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFeedType), v))
	})
}

// FeedTypeContainsFold applies the ContainsFold predicate on the "feed_type" field.
func FeedTypeContainsFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFeedType), v))
	})
}

// MarketValueLowEQ applies the EQ predicate on the "market_value_low" field.
func MarketValueLowEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowNEQ applies the NEQ predicate on the "market_value_low" field.
func MarketValueLowNEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowIn applies the In predicate on the "market_value_low" field.
func MarketValueLowIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMarketValueLow), v...))
	})
}

// MarketValueLowNotIn applies the NotIn predicate on the "market_value_low" field.
func MarketValueLowNotIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMarketValueLow), v...))
	})
}

// MarketValueLowGT applies the GT predicate on the "market_value_low" field.
func MarketValueLowGT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowGTE applies the GTE predicate on the "market_value_low" field.
func MarketValueLowGTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowLT applies the LT predicate on the "market_value_low" field.
func MarketValueLowLT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowLTE applies the LTE predicate on the "market_value_low" field.
func MarketValueLowLTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMarketValueLow), v))
	})
}

// MarketValueLowIsNil applies the IsNil predicate on the "market_value_low" field.
func MarketValueLowIsNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMarketValueLow)))
	})
}

// MarketValueLowNotNil applies the NotNil predicate on the "market_value_low" field.
func MarketValueLowNotNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMarketValueLow)))
	})
}

// MarketValueHighEQ applies the EQ predicate on the "market_value_high" field.
func MarketValueHighEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighNEQ applies the NEQ predicate on the "market_value_high" field.
func MarketValueHighNEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighIn applies the In predicate on the "market_value_high" field.
func MarketValueHighIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMarketValueHigh), v...))
	})
}

// MarketValueHighNotIn applies the NotIn predicate on the "market_value_high" field.
func MarketValueHighNotIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMarketValueHigh), v...))
	})
}

// MarketValueHighGT applies the GT predicate on the "market_value_high" field.
func MarketValueHighGT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighGTE applies the GTE predicate on the "market_value_high" field.
func MarketValueHighGTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighLT applies the LT predicate on the "market_value_high" field.
func MarketValueHighLT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighLTE applies the LTE predicate on the "market_value_high" field.
func MarketValueHighLTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMarketValueHigh), v))
	})
}

// MarketValueHighIsNil applies the IsNil predicate on the "market_value_high" field.
func MarketValueHighIsNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMarketValueHigh)))
	})
}

// MarketValueHighNotNil applies the NotNil predicate on the "market_value_high" field.
func MarketValueHighNotNil() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMarketValueHigh)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		p(s.Not())
	})
}
