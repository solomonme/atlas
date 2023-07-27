package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Organization struct {
	ent.Schema
}

func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Float("revenues").Min(0),
		field.Float("security_score").Default(50).Min(0).Max(100),
	}
}

func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("audits", Audit.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
