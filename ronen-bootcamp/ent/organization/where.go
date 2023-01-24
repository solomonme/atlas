// Code generated by ent, DO NOT EDIT.

package organization

import (
	"playground/ronen-bootcamp/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Revenues applies equality check predicate on the "revenues" field. It's identical to RevenuesEQ.
func Revenues(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevenues), v))
	})
}

// SecurityScore applies equality check predicate on the "security_score" field. It's identical to SecurityScoreEQ.
func SecurityScore(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityScore), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RevenuesEQ applies the EQ predicate on the "revenues" field.
func RevenuesEQ(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevenues), v))
	})
}

// RevenuesNEQ applies the NEQ predicate on the "revenues" field.
func RevenuesNEQ(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevenues), v))
	})
}

// RevenuesIn applies the In predicate on the "revenues" field.
func RevenuesIn(vs ...float64) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRevenues), v...))
	})
}

// RevenuesNotIn applies the NotIn predicate on the "revenues" field.
func RevenuesNotIn(vs ...float64) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRevenues), v...))
	})
}

// RevenuesGT applies the GT predicate on the "revenues" field.
func RevenuesGT(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevenues), v))
	})
}

// RevenuesGTE applies the GTE predicate on the "revenues" field.
func RevenuesGTE(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevenues), v))
	})
}

// RevenuesLT applies the LT predicate on the "revenues" field.
func RevenuesLT(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevenues), v))
	})
}

// RevenuesLTE applies the LTE predicate on the "revenues" field.
func RevenuesLTE(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevenues), v))
	})
}

// SecurityScoreEQ applies the EQ predicate on the "security_score" field.
func SecurityScoreEQ(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecurityScore), v))
	})
}

// SecurityScoreNEQ applies the NEQ predicate on the "security_score" field.
func SecurityScoreNEQ(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecurityScore), v))
	})
}

// SecurityScoreIn applies the In predicate on the "security_score" field.
func SecurityScoreIn(vs ...float64) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSecurityScore), v...))
	})
}

// SecurityScoreNotIn applies the NotIn predicate on the "security_score" field.
func SecurityScoreNotIn(vs ...float64) predicate.Organization {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSecurityScore), v...))
	})
}

// SecurityScoreGT applies the GT predicate on the "security_score" field.
func SecurityScoreGT(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecurityScore), v))
	})
}

// SecurityScoreGTE applies the GTE predicate on the "security_score" field.
func SecurityScoreGTE(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecurityScore), v))
	})
}

// SecurityScoreLT applies the LT predicate on the "security_score" field.
func SecurityScoreLT(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecurityScore), v))
	})
}

// SecurityScoreLTE applies the LTE predicate on the "security_score" field.
func SecurityScoreLTE(v float64) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecurityScore), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAudits applies the HasEdge predicate on the "audits" edge.
func HasAudits() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuditsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuditsTable, AuditsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuditsWith applies the HasEdge predicate on the "audits" edge with a given conditions (other predicates).
func HasAuditsWith(preds ...predicate.Audit) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuditsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuditsTable, AuditsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}
