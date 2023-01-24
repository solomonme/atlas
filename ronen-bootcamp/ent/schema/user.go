package schema

import (
	"playground/ronen-bootcamp/ent/privacy"
	"playground/ronen-bootcamp/rule"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("email").
			Unique(),
		field.Float("balance"),
		field.Int("org_id"),
	}
}

// Edges of the user.
// 1 to M from User to Audits
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("audits", Audit.Type),
		edge.From("organizations", Organization.Type).
			Ref("users").
			Field("org_id").
			Required().
			Unique(),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowIfAdmin(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			privacy.AlwaysAllowRule(),
		},
	}
}
