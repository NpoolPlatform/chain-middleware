package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	npool "github.com/NpoolPlatform/message/npool/chain/mgr/v1/appcoin/description"
	"github.com/google/uuid"
)

// CoinDescription holds the schema definition for the CoinDescription entity.
type CoinDescription struct {
	ent.Schema
}

func (CoinDescription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CoinDescription.
func (CoinDescription) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("used_for").
			Optional().
			Default(npool.UsedFor_DefaultUsedFor.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("message").
			Optional().
			Default(""),
	}
}

// Edges of the CoinDescription.
func (CoinDescription) Edges() []ent.Edge {
	return nil
}
