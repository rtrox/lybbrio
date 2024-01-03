// Code generated by ent, DO NOT EDIT.

package book

import (
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ksuid.ID) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// CalibreID applies equality check predicate on the "calibre_id" field. It's identical to CalibreIDEQ.
func CalibreID(v int64) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCalibreID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSort, v))
}

// PublishedDate applies equality check predicate on the "published_date" field. It's identical to PublishedDateEQ.
func PublishedDate(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublishedDate, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPath, v))
}

// Isbn applies equality check predicate on the "isbn" field. It's identical to IsbnEQ.
func Isbn(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// SeriesIndex applies equality check predicate on the "series_index" field. It's identical to SeriesIndexEQ.
func SeriesIndex(v float64) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSeriesIndex, v))
}

// CalibreIDEQ applies the EQ predicate on the "calibre_id" field.
func CalibreIDEQ(v int64) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCalibreID, v))
}

// CalibreIDNEQ applies the NEQ predicate on the "calibre_id" field.
func CalibreIDNEQ(v int64) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCalibreID, v))
}

// CalibreIDIn applies the In predicate on the "calibre_id" field.
func CalibreIDIn(vs ...int64) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCalibreID, vs...))
}

// CalibreIDNotIn applies the NotIn predicate on the "calibre_id" field.
func CalibreIDNotIn(vs ...int64) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCalibreID, vs...))
}

// CalibreIDGT applies the GT predicate on the "calibre_id" field.
func CalibreIDGT(v int64) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCalibreID, v))
}

// CalibreIDGTE applies the GTE predicate on the "calibre_id" field.
func CalibreIDGTE(v int64) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCalibreID, v))
}

// CalibreIDLT applies the LT predicate on the "calibre_id" field.
func CalibreIDLT(v int64) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCalibreID, v))
}

// CalibreIDLTE applies the LTE predicate on the "calibre_id" field.
func CalibreIDLTE(v int64) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCalibreID, v))
}

// CalibreIDIsNil applies the IsNil predicate on the "calibre_id" field.
func CalibreIDIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldCalibreID))
}

// CalibreIDNotNil applies the NotNil predicate on the "calibre_id" field.
func CalibreIDNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldCalibreID))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldSort, v))
}

// SortContains applies the Contains predicate on the "sort" field.
func SortContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldSort, v))
}

// SortHasPrefix applies the HasPrefix predicate on the "sort" field.
func SortHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldSort, v))
}

// SortHasSuffix applies the HasSuffix predicate on the "sort" field.
func SortHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldSort, v))
}

// SortEqualFold applies the EqualFold predicate on the "sort" field.
func SortEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldSort, v))
}

// SortContainsFold applies the ContainsFold predicate on the "sort" field.
func SortContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldSort, v))
}

// PublishedDateEQ applies the EQ predicate on the "published_date" field.
func PublishedDateEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublishedDate, v))
}

// PublishedDateNEQ applies the NEQ predicate on the "published_date" field.
func PublishedDateNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPublishedDate, v))
}

// PublishedDateIn applies the In predicate on the "published_date" field.
func PublishedDateIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPublishedDate, vs...))
}

// PublishedDateNotIn applies the NotIn predicate on the "published_date" field.
func PublishedDateNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPublishedDate, vs...))
}

// PublishedDateGT applies the GT predicate on the "published_date" field.
func PublishedDateGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPublishedDate, v))
}

// PublishedDateGTE applies the GTE predicate on the "published_date" field.
func PublishedDateGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPublishedDate, v))
}

// PublishedDateLT applies the LT predicate on the "published_date" field.
func PublishedDateLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPublishedDate, v))
}

// PublishedDateLTE applies the LTE predicate on the "published_date" field.
func PublishedDateLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPublishedDate, v))
}

// PublishedDateIsNil applies the IsNil predicate on the "published_date" field.
func PublishedDateIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldPublishedDate))
}

// PublishedDateNotNil applies the NotNil predicate on the "published_date" field.
func PublishedDateNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldPublishedDate))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPath, v))
}

// IsbnEQ applies the EQ predicate on the "isbn" field.
func IsbnEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldIsbn, v))
}

// IsbnNEQ applies the NEQ predicate on the "isbn" field.
func IsbnNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldIsbn, v))
}

// IsbnIn applies the In predicate on the "isbn" field.
func IsbnIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldIsbn, vs...))
}

// IsbnNotIn applies the NotIn predicate on the "isbn" field.
func IsbnNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldIsbn, vs...))
}

// IsbnGT applies the GT predicate on the "isbn" field.
func IsbnGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldIsbn, v))
}

// IsbnGTE applies the GTE predicate on the "isbn" field.
func IsbnGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldIsbn, v))
}

// IsbnLT applies the LT predicate on the "isbn" field.
func IsbnLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldIsbn, v))
}

// IsbnLTE applies the LTE predicate on the "isbn" field.
func IsbnLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldIsbn, v))
}

// IsbnContains applies the Contains predicate on the "isbn" field.
func IsbnContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldIsbn, v))
}

// IsbnHasPrefix applies the HasPrefix predicate on the "isbn" field.
func IsbnHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldIsbn, v))
}

// IsbnHasSuffix applies the HasSuffix predicate on the "isbn" field.
func IsbnHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldIsbn, v))
}

// IsbnIsNil applies the IsNil predicate on the "isbn" field.
func IsbnIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldIsbn))
}

// IsbnNotNil applies the NotNil predicate on the "isbn" field.
func IsbnNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldIsbn))
}

// IsbnEqualFold applies the EqualFold predicate on the "isbn" field.
func IsbnEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldIsbn, v))
}

// IsbnContainsFold applies the ContainsFold predicate on the "isbn" field.
func IsbnContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldIsbn, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldDescription, v))
}

// SeriesIndexEQ applies the EQ predicate on the "series_index" field.
func SeriesIndexEQ(v float64) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSeriesIndex, v))
}

// SeriesIndexNEQ applies the NEQ predicate on the "series_index" field.
func SeriesIndexNEQ(v float64) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldSeriesIndex, v))
}

// SeriesIndexIn applies the In predicate on the "series_index" field.
func SeriesIndexIn(vs ...float64) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldSeriesIndex, vs...))
}

// SeriesIndexNotIn applies the NotIn predicate on the "series_index" field.
func SeriesIndexNotIn(vs ...float64) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldSeriesIndex, vs...))
}

// SeriesIndexGT applies the GT predicate on the "series_index" field.
func SeriesIndexGT(v float64) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldSeriesIndex, v))
}

// SeriesIndexGTE applies the GTE predicate on the "series_index" field.
func SeriesIndexGTE(v float64) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldSeriesIndex, v))
}

// SeriesIndexLT applies the LT predicate on the "series_index" field.
func SeriesIndexLT(v float64) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldSeriesIndex, v))
}

// SeriesIndexLTE applies the LTE predicate on the "series_index" field.
func SeriesIndexLTE(v float64) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldSeriesIndex, v))
}

// SeriesIndexIsNil applies the IsNil predicate on the "series_index" field.
func SeriesIndexIsNil() predicate.Book {
	return predicate.Book(sql.FieldIsNull(FieldSeriesIndex))
}

// SeriesIndexNotNil applies the NotNil predicate on the "series_index" field.
func SeriesIndexNotNil() predicate.Book {
	return predicate.Book(sql.FieldNotNull(FieldSeriesIndex))
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AuthorsTable, AuthorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.Author) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPublisher applies the HasEdge predicate on the "publisher" edge.
func HasPublisher() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PublisherTable, PublisherPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPublisherWith applies the HasEdge predicate on the "publisher" edge with a given conditions (other predicates).
func HasPublisherWith(preds ...predicate.Publisher) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newPublisherStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSeries applies the HasEdge predicate on the "series" edge.
func HasSeries() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SeriesTable, SeriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSeriesWith applies the HasEdge predicate on the "series" edge with a given conditions (other predicates).
func HasSeriesWith(preds ...predicate.Series) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newSeriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIdentifiers applies the HasEdge predicate on the "identifiers" edge.
func HasIdentifiers() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IdentifiersTable, IdentifiersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIdentifiersWith applies the HasEdge predicate on the "identifiers" edge with a given conditions (other predicates).
func HasIdentifiersWith(preds ...predicate.Identifier) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newIdentifiersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLanguage applies the HasEdge predicate on the "language" edge.
func HasLanguage() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LanguageTable, LanguagePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLanguageWith applies the HasEdge predicate on the "language" edge with a given conditions (other predicates).
func HasLanguageWith(preds ...predicate.Language) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newLanguageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShelf applies the HasEdge predicate on the "shelf" edge.
func HasShelf() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ShelfTable, ShelfPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShelfWith applies the HasEdge predicate on the "shelf" edge with a given conditions (other predicates).
func HasShelfWith(preds ...predicate.Shelf) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newShelfStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(sql.NotPredicates(p))
}
