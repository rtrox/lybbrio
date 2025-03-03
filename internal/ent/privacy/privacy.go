// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"lybbrio/internal/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AuthorQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AuthorQueryRuleFunc func(context.Context, *ent.AuthorQuery) error

// EvalQuery return f(ctx, q).
func (f AuthorQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AuthorQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AuthorQuery", q)
}

// The AuthorMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AuthorMutationRuleFunc func(context.Context, *ent.AuthorMutation) error

// EvalMutation calls f(ctx, m).
func (f AuthorMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AuthorMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AuthorMutation", m)
}

// The BookQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BookQueryRuleFunc func(context.Context, *ent.BookQuery) error

// EvalQuery return f(ctx, q).
func (f BookQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BookQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BookQuery", q)
}

// The BookMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BookMutationRuleFunc func(context.Context, *ent.BookMutation) error

// EvalMutation calls f(ctx, m).
func (f BookMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BookMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BookMutation", m)
}

// The BookCoverQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BookCoverQueryRuleFunc func(context.Context, *ent.BookCoverQuery) error

// EvalQuery return f(ctx, q).
func (f BookCoverQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BookCoverQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BookCoverQuery", q)
}

// The BookCoverMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BookCoverMutationRuleFunc func(context.Context, *ent.BookCoverMutation) error

// EvalMutation calls f(ctx, m).
func (f BookCoverMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BookCoverMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BookCoverMutation", m)
}

// The BookFileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BookFileQueryRuleFunc func(context.Context, *ent.BookFileQuery) error

// EvalQuery return f(ctx, q).
func (f BookFileQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BookFileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BookFileQuery", q)
}

// The BookFileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BookFileMutationRuleFunc func(context.Context, *ent.BookFileMutation) error

// EvalMutation calls f(ctx, m).
func (f BookFileMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BookFileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BookFileMutation", m)
}

// The IdentifierQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type IdentifierQueryRuleFunc func(context.Context, *ent.IdentifierQuery) error

// EvalQuery return f(ctx, q).
func (f IdentifierQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.IdentifierQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.IdentifierQuery", q)
}

// The IdentifierMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type IdentifierMutationRuleFunc func(context.Context, *ent.IdentifierMutation) error

// EvalMutation calls f(ctx, m).
func (f IdentifierMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.IdentifierMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.IdentifierMutation", m)
}

// The LanguageQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LanguageQueryRuleFunc func(context.Context, *ent.LanguageQuery) error

// EvalQuery return f(ctx, q).
func (f LanguageQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LanguageQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LanguageQuery", q)
}

// The LanguageMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LanguageMutationRuleFunc func(context.Context, *ent.LanguageMutation) error

// EvalMutation calls f(ctx, m).
func (f LanguageMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LanguageMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LanguageMutation", m)
}

// The PublisherQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PublisherQueryRuleFunc func(context.Context, *ent.PublisherQuery) error

// EvalQuery return f(ctx, q).
func (f PublisherQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PublisherQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PublisherQuery", q)
}

// The PublisherMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PublisherMutationRuleFunc func(context.Context, *ent.PublisherMutation) error

// EvalMutation calls f(ctx, m).
func (f PublisherMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PublisherMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PublisherMutation", m)
}

// The SeriesQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SeriesQueryRuleFunc func(context.Context, *ent.SeriesQuery) error

// EvalQuery return f(ctx, q).
func (f SeriesQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SeriesQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SeriesQuery", q)
}

// The SeriesMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SeriesMutationRuleFunc func(context.Context, *ent.SeriesMutation) error

// EvalMutation calls f(ctx, m).
func (f SeriesMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SeriesMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SeriesMutation", m)
}

// The ShelfQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ShelfQueryRuleFunc func(context.Context, *ent.ShelfQuery) error

// EvalQuery return f(ctx, q).
func (f ShelfQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ShelfQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ShelfQuery", q)
}

// The ShelfMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ShelfMutationRuleFunc func(context.Context, *ent.ShelfMutation) error

// EvalMutation calls f(ctx, m).
func (f ShelfMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ShelfMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ShelfMutation", m)
}

// The TagQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TagQueryRuleFunc func(context.Context, *ent.TagQuery) error

// EvalQuery return f(ctx, q).
func (f TagQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TagQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TagQuery", q)
}

// The TagMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TagMutationRuleFunc func(context.Context, *ent.TagMutation) error

// EvalMutation calls f(ctx, m).
func (f TagMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TagMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TagMutation", m)
}

// The TaskQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TaskQueryRuleFunc func(context.Context, *ent.TaskQuery) error

// EvalQuery return f(ctx, q).
func (f TaskQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TaskQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TaskQuery", q)
}

// The TaskMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TaskMutationRuleFunc func(context.Context, *ent.TaskMutation) error

// EvalMutation calls f(ctx, m).
func (f TaskMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TaskMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TaskMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}

// The UserPermissionsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserPermissionsQueryRuleFunc func(context.Context, *ent.UserPermissionsQuery) error

// EvalQuery return f(ctx, q).
func (f UserPermissionsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserPermissionsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserPermissionsQuery", q)
}

// The UserPermissionsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserPermissionsMutationRuleFunc func(context.Context, *ent.UserPermissionsMutation) error

// EvalMutation calls f(ctx, m).
func (f UserPermissionsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserPermissionsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserPermissionsMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AuthorQuery:
		return q.Filter(), nil
	case *ent.BookQuery:
		return q.Filter(), nil
	case *ent.BookCoverQuery:
		return q.Filter(), nil
	case *ent.BookFileQuery:
		return q.Filter(), nil
	case *ent.IdentifierQuery:
		return q.Filter(), nil
	case *ent.LanguageQuery:
		return q.Filter(), nil
	case *ent.PublisherQuery:
		return q.Filter(), nil
	case *ent.SeriesQuery:
		return q.Filter(), nil
	case *ent.ShelfQuery:
		return q.Filter(), nil
	case *ent.TagQuery:
		return q.Filter(), nil
	case *ent.TaskQuery:
		return q.Filter(), nil
	case *ent.UserQuery:
		return q.Filter(), nil
	case *ent.UserPermissionsQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AuthorMutation:
		return m.Filter(), nil
	case *ent.BookMutation:
		return m.Filter(), nil
	case *ent.BookCoverMutation:
		return m.Filter(), nil
	case *ent.BookFileMutation:
		return m.Filter(), nil
	case *ent.IdentifierMutation:
		return m.Filter(), nil
	case *ent.LanguageMutation:
		return m.Filter(), nil
	case *ent.PublisherMutation:
		return m.Filter(), nil
	case *ent.SeriesMutation:
		return m.Filter(), nil
	case *ent.ShelfMutation:
		return m.Filter(), nil
	case *ent.TagMutation:
		return m.Filter(), nil
	case *ent.TaskMutation:
		return m.Filter(), nil
	case *ent.UserMutation:
		return m.Filter(), nil
	case *ent.UserPermissionsMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
