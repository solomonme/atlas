package schema

import (
	"playground/ronen-bootcamp/ent/privacy"
	"playground/ronen-bootcamp/rule"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Fields of the Audit.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("identity"),
		field.Time("timestamp").Default(time.Now),
		field.Float("balance"),
		field.String("description").NotEmpty(),
		field.Int("org_id"),
	}
}

// Edges of the Audit.
// M to 1 from Audit to users
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("audits").
			Unique(),
		edge.From("organizations", Organization.Type).
			Ref("audits").
			Field("org_id").
			Required().
			Unique(),
	}
}

// Policy defines the privacy policy of the Audit.
func (Audit) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowIfAdmin(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			// Allow any viewer to read anything.
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
