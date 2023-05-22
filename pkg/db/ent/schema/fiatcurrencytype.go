package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// FiatCurrencyType holds the schema definition for the FiatCurrencyType entity.
type FiatCurrencyType struct {
	ent.Schema
}

func (FiatCurrencyType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the FiatCurrencyType.
func (FiatCurrencyType) Fields() []ent.Field {
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

// Edges of the FiatCurrencyType.
func (FiatCurrencyType) Edges() []ent.Edge {
	return nil
}
