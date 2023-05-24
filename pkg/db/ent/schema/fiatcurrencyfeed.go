//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

// FiatCurrencyFeed holds the schema definition for the FiatCurrencyFeed entity.
type FiatCurrencyFeed struct {
	ent.Schema
}

func (FiatCurrencyFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the FiatCurrencyFeed.
func (FiatCurrencyFeed) Fields() []ent.Field {
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
			String("feed_fiat_name").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the FiatCurrencyFeed.
func (FiatCurrencyFeed) Edges() []ent.Edge {
	return nil
}

func (FiatCurrencyFeed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fiat_id", "id"),
		index.Fields("fiat_id"),
	}
}
