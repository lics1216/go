// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAddresses holds the string denoting the addresses edge name in mutations.
	EdgeAddresses = "addresses"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AddressesTable is the table that holds the addresses relation/edge.
	AddressesTable = "addresses"
	// AddressesInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressesInverseTable = "addresses"
	// AddressesColumn is the table column denoting the addresses relation/edge.
	AddressesColumn = "user_addresses"
	// CardsTable is the table that holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "user_cards"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPasswordHash,
	FieldCreatedAt,
	FieldUpdatedAt,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
