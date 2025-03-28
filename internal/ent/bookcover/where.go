// Code generated by ent, DO NOT EDIT.

package bookcover

import (
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ksuid.ID) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldUpdateTime, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldPath, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldSize, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldHeight, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldURL, v))
}

// ContentType applies equality check predicate on the "contentType" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldContentType, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldUpdateTime, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContainsFold(FieldPath, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldSize, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldWidth, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldHeight, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContainsFold(FieldURL, v))
}

// ContentTypeEQ applies the EQ predicate on the "contentType" field.
func ContentTypeEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEQ(FieldContentType, v))
}

// ContentTypeNEQ applies the NEQ predicate on the "contentType" field.
func ContentTypeNEQ(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNEQ(FieldContentType, v))
}

// ContentTypeIn applies the In predicate on the "contentType" field.
func ContentTypeIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldIn(FieldContentType, vs...))
}

// ContentTypeNotIn applies the NotIn predicate on the "contentType" field.
func ContentTypeNotIn(vs ...string) predicate.BookCover {
	return predicate.BookCover(sql.FieldNotIn(FieldContentType, vs...))
}

// ContentTypeGT applies the GT predicate on the "contentType" field.
func ContentTypeGT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGT(FieldContentType, v))
}

// ContentTypeGTE applies the GTE predicate on the "contentType" field.
func ContentTypeGTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldGTE(FieldContentType, v))
}

// ContentTypeLT applies the LT predicate on the "contentType" field.
func ContentTypeLT(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLT(FieldContentType, v))
}

// ContentTypeLTE applies the LTE predicate on the "contentType" field.
func ContentTypeLTE(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldLTE(FieldContentType, v))
}

// ContentTypeContains applies the Contains predicate on the "contentType" field.
func ContentTypeContains(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContains(FieldContentType, v))
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "contentType" field.
func ContentTypeHasPrefix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasPrefix(FieldContentType, v))
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "contentType" field.
func ContentTypeHasSuffix(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldHasSuffix(FieldContentType, v))
}

// ContentTypeEqualFold applies the EqualFold predicate on the "contentType" field.
func ContentTypeEqualFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldEqualFold(FieldContentType, v))
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "contentType" field.
func ContentTypeContainsFold(v string) predicate.BookCover {
	return predicate.BookCover(sql.FieldContainsFold(FieldContentType, v))
}

// HasBook applies the HasEdge predicate on the "book" edge.
func HasBook() predicate.BookCover {
	return predicate.BookCover(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BookTable, BookColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBookWith applies the HasEdge predicate on the "book" edge with a given conditions (other predicates).
func HasBookWith(preds ...predicate.Book) predicate.BookCover {
	return predicate.BookCover(func(s *sql.Selector) {
		step := newBookStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BookCover) predicate.BookCover {
	return predicate.BookCover(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BookCover) predicate.BookCover {
	return predicate.BookCover(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BookCover) predicate.BookCover {
	return predicate.BookCover(sql.NotPredicates(p))
}
