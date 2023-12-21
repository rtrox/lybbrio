// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"

	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/viewer"

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

// func FilterTenantRule() privacy.QueryMutationRule {
// 	type UserFilter interface {
// 		WhereUserID(entql.StringP)
// 	}
// 	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
// 		view := viewer.FromContext(ctx)
// 		uid, ok := view.User()
// 		if !ok {
// 			return privacy.Denyf("missing user information in viewer-context")
// 		}
// 		uf, ok := f.(UserFilter)
// 		if !ok {
// 			return privacy.Denyf("filter does not implement UserFilter")
// 		}
// 		uf.WhereUserID(entql.StringEQ(uid))
// 		if f, ok := f.(UserFilter); ok {
// 			f.WhereUserID(view.UserID())
// 		}
// 		return privacy.Skip
// 	})
// }
