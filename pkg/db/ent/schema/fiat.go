package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Fiat holds the schema definition for the Fiat entity.
type Fiat struct {
	ent.Schema
}

func (Fiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Fiat.
func (Fiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
	}
}

// Edges of the Fiat.
func (Fiat) Edges() []ent.Edge {
	return nil
}
