// Code generated by ent, DO NOT EDIT.

package publisher

import (
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ksuid.ID) predicate.Publisher {
	return predicate.Publisher(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Publisher {
	return predicate.Publisher(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Publisher {
	return predicate.Publisher(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Publisher {
	return predicate.Publisher(sql.FieldContainsFold(FieldName, v))
}

// HasBooks applies the HasEdge predicate on the "books" edge.
func HasBooks() predicate.Publisher {
	return predicate.Publisher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BooksTable, BooksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBooksWith applies the HasEdge predicate on the "books" edge with a given conditions (other predicates).
func HasBooksWith(preds ...predicate.Book) predicate.Publisher {
	return predicate.Publisher(func(s *sql.Selector) {
		step := newBooksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Publisher) predicate.Publisher {
	return predicate.Publisher(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Publisher) predicate.Publisher {
	return predicate.Publisher(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Publisher) predicate.Publisher {
	return predicate.Publisher(sql.NotPredicates(p))
}