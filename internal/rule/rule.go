// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"

	"lybbrio/internal/ent"
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/middleware"
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

func AllowIfSuperRead() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		superRead := middleware.SuperReadFromCtx(ctx)
		if view.IsAdmin() && superRead {
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
		uid, ok := view.UserID()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		uf, ok := f.(UserFilter)
		if !ok {
			return privacy.Denyf("filter does not implement UserFilter")
		}
		uf.WhereUserID(entql.StringEQ(uid.String()))
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
		uid, ok := view.UserID()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		_, ok = f.(UserOrPublicFilter)
		if !ok {
			return privacy.Denyf("filter does not implement UserOrPublicFilter")
		}
		f.Where(
			entql.Or(
				entql.FieldEQ("user_id", uid.String()),
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
			if !view.Has(permissions.CanCreatePublic) {
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
		type UserableMutation interface {
			UserID() (ksuid.ID, bool)
			OldUserID(context.Context) (ksuid.ID, error)
		}
		um, ok := m.(UserableMutation)
		if !ok {
			return privacy.Denyf("mutation does not implement UserableMutation")
		}
		view := viewer.FromContext(ctx)
		viewerID, ok := view.UserID()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}

		var uid ksuid.ID
		switch m.Op() {
		case ent.OpCreate:
			var ok bool
			uid, ok = m.(UserableMutation).UserID()
			if ok && uid != viewerID {
				return privacy.Denyf("cannot create objects owned by other users")
			}
		case ent.OpUpdateOne:
			var err error
			uid, err = um.OldUserID(ctx)
			if err != nil {
				return err
			}
			if uid != viewerID {
				return privacy.Denyf("cannot mutate objects owned by other users")
			}
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
		uid, ok := view.UserID()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		sf, ok := f.(SelfFilter)
		if !ok {
			return privacy.Denyf("filter does not implement SelfFilter")
		}
		sf.WhereID(entql.StringEQ(uid.String()))
		return privacy.Skip
	})
}

func FilterUserOrSystemRule() privacy.QueryMutationRule {
	type CreatorOrSystemFilter interface {
		WhereUserID(entql.StringP)
		WhereIsSystemTask(entql.BoolP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		uid, ok := view.UserID()
		if !ok {
			return privacy.Denyf("missing user information in viewer-context")
		}
		_, ok = f.(CreatorOrSystemFilter)
		if !ok {
			return privacy.Denyf("filter does not implement CreatorOrSystemFilter")
		}
		f.Where(
			entql.Or(
				entql.FieldEQ("user_id", uid.String()),
				entql.FieldEQ("is_system_task", true),
			),
		)
		return privacy.Skip
	})
}

func AllowTasksOfType(types ...any) privacy.MutationRule {
	return privacy.TaskMutationRuleFunc(func(ctx context.Context, m *ent.TaskMutation) error {

		var typ task_enums.TaskType
		switch m.Op() {
		case ent.OpCreate:
			taskType, ok := m.GetType()
			if !ok {
				return privacy.Denyf("missing type field in mutation")
			}
			typ = taskType
		case ent.OpUpdateOne:
			var err error
			// Type is immutable, checking OldType is enough.
			typ, err = m.OldType(ctx)
			if err != nil {
				return err
			}
		}
		for _, t := range types {
			if typ == t {
				return privacy.Allow
			}
		}
		return privacy.Denyf("cannot mutate tasks of type %v", typ)
	})
}

func DenySystemTaskForNonAdmin() privacy.MutationRule {
	return privacy.TaskMutationRuleFunc(func(ctx context.Context, m *ent.TaskMutation) error {
		view := viewer.FromContext(ctx)
		if view.IsAdmin() {
			return privacy.Skip
		}
		if m.Op() == ent.OpUpdateOne {
			system, err := m.OldIsSystemTask(ctx)
			if err != nil {
				return err
			}
			if system {
				return privacy.Denyf("cannot update system tasks")
			}
		}
		isSystemTask, ok := m.IsSystemTask()
		if !ok {
			return privacy.Skip
		} else if isSystemTask {
			return privacy.Denyf("cannot create system tasks")
		}
		return privacy.Skip
	})
}
