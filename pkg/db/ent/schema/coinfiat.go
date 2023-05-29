//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// CoinFiat holds the schema definition for the CoinFiat entity.
type CoinFiat struct {
	ent.Schema
}

func (CoinFiat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CoinFiat.
func (CoinFiat) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(basetypes.CurrencyFeedType_DefaultFeedType.String()),
	}
}

// Edges of the CoinFiat.
func (CoinFiat) Edges() []ent.Edge {
	return nil
}
