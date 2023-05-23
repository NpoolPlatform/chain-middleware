// Code generated by ent, DO NOT EDIT.

package coinextra

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coinextra type in the database.
	Label = "coin_extra"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldHomePage holds the string denoting the home_page field in the database.
	FieldHomePage = "home_page"
	// FieldSpecs holds the string denoting the specs field in the database.
	FieldSpecs = "specs"
	// FieldStableUsd holds the string denoting the stable_usd field in the database.
	FieldStableUsd = "stable_usd"
	// Table holds the table name of the coinextra in the database.
	Table = "coin_extras"
)

// Columns holds all SQL columns for coinextra fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCoinTypeID,
	FieldHomePage,
	FieldSpecs,
	FieldStableUsd,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/chain-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultHomePage holds the default value on creation for the "home_page" field.
	DefaultHomePage string
	// DefaultSpecs holds the default value on creation for the "specs" field.
	DefaultSpecs string
	// DefaultStableUsd holds the default value on creation for the "stable_usd" field.
	DefaultStableUsd bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)