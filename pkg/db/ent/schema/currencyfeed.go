package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/coin/currency"
)

// CurrencyFeed holds the schema definition for the CurrencyFeed entity.
type CurrencyFeed struct {
	ent.Schema
}

func (CurrencyFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CurrencyFeed.
func (CurrencyFeed) Fields() []ent.Field {
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
			String("feed_coin_name").
			Optional().
			Default(""),
		field.
			Bool("disabled").
			Optional().
			Default(false),
	}
}

// Edges of the CurrencyFeed.
func (CurrencyFeed) Edges() []ent.Edge {
	return nil
}

func (CurrencyFeed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coin_type_id", "id"),
		index.Fields("coin_type_id"),
	}
}
