// Code generated by ent, DO NOT EDIT.

package userpermissions

import (
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ksuid.ID) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldUpdateTime, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldEQ(FieldUserID, vc))
}

// Admin applies equality check predicate on the "Admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldAdmin, v))
}

// CanCreatePublic applies equality check predicate on the "CanCreatePublic" field. It's identical to CanCreatePublicEQ.
func CanCreatePublic(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCanCreatePublic, v))
}

// CanEdit applies equality check predicate on the "CanEdit" field. It's identical to CanEditEQ.
func CanEdit(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCanEdit, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldLTE(FieldUpdateTime, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...ksuid.ID) predicate.UserPermissions {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.UserPermissions(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...ksuid.ID) predicate.UserPermissions {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.UserPermissions(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldLTE(FieldUserID, vc))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v ksuid.ID) predicate.UserPermissions {
	vc := string(v)
	return predicate.UserPermissions(sql.FieldContainsFold(FieldUserID, vc))
}

// AdminEQ applies the EQ predicate on the "Admin" field.
func AdminEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldAdmin, v))
}

// AdminNEQ applies the NEQ predicate on the "Admin" field.
func AdminNEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldAdmin, v))
}

// CanCreatePublicEQ applies the EQ predicate on the "CanCreatePublic" field.
func CanCreatePublicEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCanCreatePublic, v))
}

// CanCreatePublicNEQ applies the NEQ predicate on the "CanCreatePublic" field.
func CanCreatePublicNEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldCanCreatePublic, v))
}

// CanEditEQ applies the EQ predicate on the "CanEdit" field.
func CanEditEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldEQ(FieldCanEdit, v))
}

// CanEditNEQ applies the NEQ predicate on the "CanEdit" field.
func CanEditNEQ(v bool) predicate.UserPermissions {
	return predicate.UserPermissions(sql.FieldNEQ(FieldCanEdit, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserPermissions {
	return predicate.UserPermissions(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserPermissions {
	return predicate.UserPermissions(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPermissions) predicate.UserPermissions {
	return predicate.UserPermissions(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPermissions) predicate.UserPermissions {
	return predicate.UserPermissions(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPermissions) predicate.UserPermissions {
	return predicate.UserPermissions(sql.NotPredicates(p))
}
