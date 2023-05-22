package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// CoinExtra holds the schema definition for the CoinExtra entity.
type CoinExtra struct {
	ent.Schema
}

func (CoinExtra) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CoinExtra.
func (CoinExtra) Fields() []ent.Field {
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
			String("home_page").
			Optional().
			Default(""),
		field.
			String("specs").
			Optional().
			Default(""),
		field.
			Bool("stable_usd").
			Optional().
			Default(false),
	}
}

// Edges of the CoinExtra.
func (CoinExtra) Edges() []ent.Edge {
	return nil
}
