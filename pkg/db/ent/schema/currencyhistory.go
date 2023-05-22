//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

// CurrencyHistory holds the schema definition for the CurrencyHistory entity.
type CurrencyHistory struct {
	ent.Schema
}

func (CurrencyHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CurrencyHistory.
func (CurrencyHistory) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(npool.FeedType_DefaultFeedType.String()),
		field.
			Other("market_value_high", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("market_value_low", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the CurrencyHistory.
func (CurrencyHistory) Edges() []ent.Edge {
	return nil
}

func (CurrencyHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "id"),
		index.Fields("coin_type_id"),
	}
}
