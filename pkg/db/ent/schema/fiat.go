package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// Fiat holds the schema definition for the Fiat entity.
type Fiat struct {
	ent.Schema
}

func (Fiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Fiat.
func (Fiat) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("unit").
			Optional().
			Default(""),
	}
}

// Edges of the Fiat.
func (Fiat) Edges() []ent.Edge {
	return nil
}
