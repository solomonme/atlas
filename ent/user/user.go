// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// EdgeAudits holds the string denoting the audits edge name in mutations.
	EdgeAudits = "audits"
	// EdgeOrganizations holds the string denoting the organizations edge name in mutations.
	EdgeOrganizations = "organizations"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AuditsTable is the table that holds the audits relation/edge.
	AuditsTable = "audits"
	// AuditsInverseTable is the table name for the Audit entity.
	// It exists in this package in order to avoid circular dependency with the "audit" package.
	AuditsInverseTable = "audits"
	// AuditsColumn is the table column denoting the audits relation/edge.
	AuditsColumn = "user_audits"
	// OrganizationsTable is the table that holds the organizations relation/edge.
	OrganizationsTable = "users"
	// OrganizationsInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationsInverseTable = "organizations"
	// OrganizationsColumn is the table column denoting the organizations relation/edge.
	OrganizationsColumn = "org_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldBalance,
	FieldOrgID,
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
//	import _ "playground/ronen-bootcamp/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
)
