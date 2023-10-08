// Code generated by ent, DO NOT EDIT.

package coindescription

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the coindescription type in the database.
	Label = "coin_description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldUsedFor holds the string denoting the used_for field in the database.
	FieldUsedFor = "used_for"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// Table holds the table name of the coindescription in the database.
	Table = "coin_descriptions"
)

// Columns holds all SQL columns for coindescription fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldCoinTypeID,
	FieldUsedFor,
	FieldTitle,
	FieldMessage,
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
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultUsedFor holds the default value on creation for the "used_for" field.
	DefaultUsedFor string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
