package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

// FiatCurrencyFeed holds the schema definition for the FiatCurrencyFeed entity.
type FiatCurrencyFeed struct {
	ent.Schema
}

func (FiatCurrencyFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FiatCurrencyFeed.
func (FiatCurrencyFeed) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("fiat_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("feed_type").
			Optional().
			Default(basetypes.CurrencyFeedType_DefaultFeedType.String()),
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
