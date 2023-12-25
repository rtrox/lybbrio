// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"

	"lybbrio/internal/ent"
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/viewer"

	"entgo.io/ent/entql"
	"github.com/rs/zerolog/log"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			log := log.Ctx(ctx)
			log.Error().Msg("viewer-context is missing")
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.IsAdmin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

func FilterUserRule() privacy.QueryMutationRule {
	type UserFilter interface {
		WhereUserID(entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		user, ok := view.User()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		uf, ok := f.(UserFilter)
		if !ok {
			return privacy.Denyf("filter does not implement UserFilter")
		}
		uf.WhereUserID(entql.StringEQ(user.ID.String()))
		return privacy.Skip
	})
}

func FilterUserOrPublicRule() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		type UserOrPublicFilter interface {
			WhereUserID(entql.StringP)
			WherePublic(entql.BoolP)
		}
		view := viewer.FromContext(ctx)
		user, ok := view.User()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		_, ok = f.(UserOrPublicFilter)
		if !ok {
			return privacy.Denyf("filter does not implement UserOrPublicFilter")
		}
		f.Where(
			entql.Or(
				entql.FieldEQ("user_id", user.ID.String()),
				entql.FieldEQ("public", true),
			),
		)
		return privacy.Skip
	})
}

func DenyPublicWithoutPermissionRule() privacy.MutationRule {
	type PublicableMutation interface {
		Public() (r bool, exists bool)
	}
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		pm, ok := m.(PublicableMutation)
		if !ok {
			return privacy.Denyf("mutation does not implement PublicableMutation")
		}
		public, mutated := pm.Public()
		if mutated && public {
			view := viewer.FromContext(ctx)
			perms, ok := view.Permissions()
			if !ok {
				return privacy.Denyf("missing permissions information in viewer-context")
			}
			if !perms.CanCreatePublic {
				return privacy.Denyf("user does not have permission to create public objects")
			}
		}
		return privacy.Skip
	})
}

// DenyMismatchedUserRule is a rule that returns deny decision if the viewer
// is not the same as the user on the object.
func DenyMismatchedUserRule() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		view := viewer.FromContext(ctx)
		user, ok := view.User()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		mutationUserID, ok := m.Field("user_id")
		if !ok {
			return privacy.Denyf("missing user_id field in mutation")
		} // temporary
		if user.ID != mutationUserID.(ksuid.ID) {
			return privacy.Denyf("cannot mutate objects owned by other users")
		}
		return privacy.Skip
	})
}

// FilterSelfRule is a rule that returns deny decision if the viewer
// is not the same as the user object.
func FilterSelfRule() privacy.QueryMutationRule {
	type SelfFilter interface {
		WhereID(entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		user, ok := view.User()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		sf, ok := f.(SelfFilter)
		if !ok {
			return privacy.Denyf("filter does not implement SelfFilter")
		}
		sf.WhereID(entql.StringEQ(user.ID.String()))
		return privacy.Skip
	})
}
