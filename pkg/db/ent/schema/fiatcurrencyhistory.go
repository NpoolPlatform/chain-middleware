//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

// FiatCurrencyHistory holds the schema definition for the FiatCurrencyHistory entity.
type FiatCurrencyHistory struct {
	ent.Schema
}

func (FiatCurrencyHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the FiatCurrencyHistory.
func (FiatCurrencyHistory) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(npool.FeedType_DefaultFeedType.String()),
		field.
			Other("market_value_low", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("market_value_high", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the FiatCurrencyHistory.
func (FiatCurrencyHistory) Edges() []ent.Edge {
	return nil
}
