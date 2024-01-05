// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"lybbrio/internal/ent/author"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"lybbrio/internal/ent/shelf"
	"lybbrio/internal/ent/tag"
	"lybbrio/internal/ent/task"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/ent/userpermissions"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorMixin := schema.Author{}.Mixin()
	author.Policy = privacy.NewPolicies(authorMixin[0], schema.Author{})
	author.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := author.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	authorMixinFields2 := authorMixin[2].Fields()
	_ = authorMixinFields2
	authorFields := schema.Author{}.Fields()
	_ = authorFields
	// authorDescName is the schema descriptor for name field.
	authorDescName := authorFields[0].Descriptor()
	// author.NameValidator is a validator for the "name" field. It is called by the builders before save.
	author.NameValidator = authorDescName.Validators[0].(func(string) error)
	// authorDescID is the schema descriptor for id field.
	authorDescID := authorMixinFields2[0].Descriptor()
	// author.DefaultID holds the default value on creation for the id field.
	author.DefaultID = authorDescID.Default.(func() ksuid.ID)
	bookMixin := schema.Book{}.Mixin()
	book.Policy = privacy.NewPolicies(bookMixin[0], schema.Book{})
	book.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := book.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	bookMixinFields2 := bookMixin[2].Fields()
	_ = bookMixinFields2
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescTitle is the schema descriptor for title field.
	bookDescTitle := bookFields[0].Descriptor()
	// book.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	book.TitleValidator = bookDescTitle.Validators[0].(func(string) error)
	// bookDescPath is the schema descriptor for path field.
	bookDescPath := bookFields[3].Descriptor()
	// book.PathValidator is a validator for the "path" field. It is called by the builders before save.
	book.PathValidator = bookDescPath.Validators[0].(func(string) error)
	// bookDescID is the schema descriptor for id field.
	bookDescID := bookMixinFields2[0].Descriptor()
	// book.DefaultID holds the default value on creation for the id field.
	book.DefaultID = bookDescID.Default.(func() ksuid.ID)
	bookfileMixin := schema.BookFile{}.Mixin()
	bookfile.Policy = privacy.NewPolicies(bookfileMixin[0], schema.BookFile{})
	bookfile.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := bookfile.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	bookfileMixinFields1 := bookfileMixin[1].Fields()
	_ = bookfileMixinFields1
	bookfileFields := schema.BookFile{}.Fields()
	_ = bookfileFields
	// bookfileDescName is the schema descriptor for name field.
	bookfileDescName := bookfileFields[0].Descriptor()
	// bookfile.NameValidator is a validator for the "name" field. It is called by the builders before save.
	bookfile.NameValidator = bookfileDescName.Validators[0].(func(string) error)
	// bookfileDescPath is the schema descriptor for path field.
	bookfileDescPath := bookfileFields[1].Descriptor()
	// bookfile.PathValidator is a validator for the "path" field. It is called by the builders before save.
	bookfile.PathValidator = bookfileDescPath.Validators[0].(func(string) error)
	// bookfileDescSize is the schema descriptor for size field.
	bookfileDescSize := bookfileFields[2].Descriptor()
	// bookfile.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	bookfile.SizeValidator = bookfileDescSize.Validators[0].(func(int64) error)
	// bookfileDescID is the schema descriptor for id field.
	bookfileDescID := bookfileMixinFields1[0].Descriptor()
	// bookfile.DefaultID holds the default value on creation for the id field.
	bookfile.DefaultID = bookfileDescID.Default.(func() ksuid.ID)
	identifierMixin := schema.Identifier{}.Mixin()
	identifier.Policy = privacy.NewPolicies(identifierMixin[0], schema.Identifier{})
	identifier.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := identifier.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	identifierMixinFields2 := identifierMixin[2].Fields()
	_ = identifierMixinFields2
	identifierFields := schema.Identifier{}.Fields()
	_ = identifierFields
	// identifierDescType is the schema descriptor for type field.
	identifierDescType := identifierFields[0].Descriptor()
	// identifier.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	identifier.TypeValidator = identifierDescType.Validators[0].(func(string) error)
	// identifierDescValue is the schema descriptor for value field.
	identifierDescValue := identifierFields[1].Descriptor()
	// identifier.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	identifier.ValueValidator = identifierDescValue.Validators[0].(func(string) error)
	// identifierDescID is the schema descriptor for id field.
	identifierDescID := identifierMixinFields2[0].Descriptor()
	// identifier.DefaultID holds the default value on creation for the id field.
	identifier.DefaultID = identifierDescID.Default.(func() ksuid.ID)
	languageMixin := schema.Language{}.Mixin()
	language.Policy = privacy.NewPolicies(languageMixin[0], schema.Language{})
	language.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := language.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	languageMixinFields2 := languageMixin[2].Fields()
	_ = languageMixinFields2
	languageFields := schema.Language{}.Fields()
	_ = languageFields
	// languageDescCode is the schema descriptor for code field.
	languageDescCode := languageFields[0].Descriptor()
	// language.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	language.CodeValidator = languageDescCode.Validators[0].(func(string) error)
	// languageDescID is the schema descriptor for id field.
	languageDescID := languageMixinFields2[0].Descriptor()
	// language.DefaultID holds the default value on creation for the id field.
	language.DefaultID = languageDescID.Default.(func() ksuid.ID)
	publisherMixin := schema.Publisher{}.Mixin()
	publisher.Policy = privacy.NewPolicies(publisherMixin[0], schema.Publisher{})
	publisher.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := publisher.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	publisherMixinFields2 := publisherMixin[2].Fields()
	_ = publisherMixinFields2
	publisherFields := schema.Publisher{}.Fields()
	_ = publisherFields
	// publisherDescName is the schema descriptor for name field.
	publisherDescName := publisherFields[0].Descriptor()
	// publisher.NameValidator is a validator for the "name" field. It is called by the builders before save.
	publisher.NameValidator = publisherDescName.Validators[0].(func(string) error)
	// publisherDescID is the schema descriptor for id field.
	publisherDescID := publisherMixinFields2[0].Descriptor()
	// publisher.DefaultID holds the default value on creation for the id field.
	publisher.DefaultID = publisherDescID.Default.(func() ksuid.ID)
	seriesMixin := schema.Series{}.Mixin()
	series.Policy = privacy.NewPolicies(seriesMixin[0], schema.Series{})
	series.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := series.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	seriesMixinFields2 := seriesMixin[2].Fields()
	_ = seriesMixinFields2
	seriesFields := schema.Series{}.Fields()
	_ = seriesFields
	// seriesDescName is the schema descriptor for name field.
	seriesDescName := seriesFields[0].Descriptor()
	// series.NameValidator is a validator for the "name" field. It is called by the builders before save.
	series.NameValidator = seriesDescName.Validators[0].(func(string) error)
	// seriesDescSort is the schema descriptor for sort field.
	seriesDescSort := seriesFields[1].Descriptor()
	// series.SortValidator is a validator for the "sort" field. It is called by the builders before save.
	series.SortValidator = seriesDescSort.Validators[0].(func(string) error)
	// seriesDescID is the schema descriptor for id field.
	seriesDescID := seriesMixinFields2[0].Descriptor()
	// series.DefaultID holds the default value on creation for the id field.
	series.DefaultID = seriesDescID.Default.(func() ksuid.ID)
	shelfMixin := schema.Shelf{}.Mixin()
	shelf.Policy = privacy.NewPolicies(shelfMixin[0], shelfMixin[1], schema.Shelf{})
	shelf.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := shelf.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	shelfMixinFields1 := shelfMixin[1].Fields()
	_ = shelfMixinFields1
	shelfMixinFields2 := shelfMixin[2].Fields()
	_ = shelfMixinFields2
	shelfFields := schema.Shelf{}.Fields()
	_ = shelfFields
	// shelfDescPublic is the schema descriptor for public field.
	shelfDescPublic := shelfMixinFields1[0].Descriptor()
	// shelf.DefaultPublic holds the default value on creation for the public field.
	shelf.DefaultPublic = shelfDescPublic.Default.(bool)
	// shelfDescName is the schema descriptor for name field.
	shelfDescName := shelfFields[0].Descriptor()
	// shelf.NameValidator is a validator for the "name" field. It is called by the builders before save.
	shelf.NameValidator = shelfDescName.Validators[0].(func(string) error)
	// shelfDescID is the schema descriptor for id field.
	shelfDescID := shelfMixinFields2[0].Descriptor()
	// shelf.DefaultID holds the default value on creation for the id field.
	shelf.DefaultID = shelfDescID.Default.(func() ksuid.ID)
	tagMixin := schema.Tag{}.Mixin()
	tag.Policy = privacy.NewPolicies(tagMixin[0], schema.Tag{})
	tag.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := tag.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tagMixinFields2 := tagMixin[2].Fields()
	_ = tagMixinFields2
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[0].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagMixinFields2[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() ksuid.ID)
	taskMixin := schema.Task{}.Mixin()
	task.Policy = privacy.NewPolicies(taskMixin[0], schema.Task{})
	task.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := task.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	taskMixinFields1 := taskMixin[1].Fields()
	_ = taskMixinFields1
	taskMixinFields2 := taskMixin[2].Fields()
	_ = taskMixinFields2
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreateTime is the schema descriptor for create_time field.
	taskDescCreateTime := taskMixinFields1[0].Descriptor()
	// task.DefaultCreateTime holds the default value on creation for the create_time field.
	task.DefaultCreateTime = taskDescCreateTime.Default.(func() time.Time)
	// taskDescUpdateTime is the schema descriptor for update_time field.
	taskDescUpdateTime := taskMixinFields1[1].Descriptor()
	// task.DefaultUpdateTime holds the default value on creation for the update_time field.
	task.DefaultUpdateTime = taskDescUpdateTime.Default.(func() time.Time)
	// task.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	task.UpdateDefaultUpdateTime = taskDescUpdateTime.UpdateDefault.(func() time.Time)
	// taskDescProgress is the schema descriptor for progress field.
	taskDescProgress := taskFields[2].Descriptor()
	// task.DefaultProgress holds the default value on creation for the progress field.
	task.DefaultProgress = taskDescProgress.Default.(float64)
	// taskDescIsSystemTask is the schema descriptor for is_system_task field.
	taskDescIsSystemTask := taskFields[6].Descriptor()
	// task.DefaultIsSystemTask holds the default value on creation for the is_system_task field.
	task.DefaultIsSystemTask = taskDescIsSystemTask.Default.(bool)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskMixinFields2[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() ksuid.ID)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(userMixin[0], schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields1[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() ksuid.ID)
	userpermissionsMixin := schema.UserPermissions{}.Mixin()
	userpermissions.Policy = privacy.NewPolicies(userpermissionsMixin[0], schema.UserPermissions{})
	userpermissions.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := userpermissions.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userpermissionsMixinFields1 := userpermissionsMixin[1].Fields()
	_ = userpermissionsMixinFields1
	userpermissionsFields := schema.UserPermissions{}.Fields()
	_ = userpermissionsFields
	// userpermissionsDescAdmin is the schema descriptor for Admin field.
	userpermissionsDescAdmin := userpermissionsFields[1].Descriptor()
	// userpermissions.DefaultAdmin holds the default value on creation for the Admin field.
	userpermissions.DefaultAdmin = userpermissionsDescAdmin.Default.(bool)
	// userpermissionsDescCanCreatePublic is the schema descriptor for CanCreatePublic field.
	userpermissionsDescCanCreatePublic := userpermissionsFields[2].Descriptor()
	// userpermissions.DefaultCanCreatePublic holds the default value on creation for the CanCreatePublic field.
	userpermissions.DefaultCanCreatePublic = userpermissionsDescCanCreatePublic.Default.(bool)
	// userpermissionsDescCanEdit is the schema descriptor for CanEdit field.
	userpermissionsDescCanEdit := userpermissionsFields[3].Descriptor()
	// userpermissions.DefaultCanEdit holds the default value on creation for the CanEdit field.
	userpermissions.DefaultCanEdit = userpermissionsDescCanEdit.Default.(bool)
	// userpermissionsDescID is the schema descriptor for id field.
	userpermissionsDescID := userpermissionsMixinFields1[0].Descriptor()
	// userpermissions.DefaultID holds the default value on creation for the id field.
	userpermissions.DefaultID = userpermissionsDescID.Default.(func() ksuid.ID)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)
