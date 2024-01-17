package permissions

import (
	"reflect"
	"strings"
	"unicode"
)

type Permission int

const (
	Admin Permission = iota + 1
	CanCreatePublic
	CanEdit

	count // Keep this at the end.
)

func (p Permission) String() string {
	// Each Permission String must exactly match the enum
	// name to allow reflection to work.
	switch p {
	case Admin:
		return "Admin"
	case CanCreatePublic:
		return "CanCreatePublic"
	case CanEdit:
		return "CanEdit"
	}
	return ""
}

func (p Permission) FieldName() string {
	var b strings.Builder
	for i, v := range p.String() {
		if i == 0 {
			b.WriteRune(unicode.ToLower(v))
			continue
		}
		if unicode.IsLower(v) {
			b.WriteRune(v)
		} else {
			b.WriteRune('_')
			b.WriteRune(unicode.ToLower(v))
		}
	}
	return b.String()
}

func FromString(s string) Permission {
	switch s {
	case Admin.String():
		return Admin
	case CanCreatePublic.String():
		return CanCreatePublic
	case CanEdit.String():
		return CanEdit
	default:
		return 0
	}
}

type Permissions map[Permission]struct{}

func NewPermissions(permissions ...Permission) Permissions {
	ret := Permissions{}
	for _, p := range permissions {
		ret[p] = struct{}{}
	}
	return ret
}

func (p Permissions) Has(perm Permission) bool {
	_, ok := p[perm]
	return ok
}

func (p Permissions) Add(perm Permission) {
	p[perm] = struct{}{}
}

func (p Permissions) StringSlice() []string {
	ret := make([]string, 0, len(p))
	for k := range p {
		ret = append(ret, k.String())
	}
	return ret
}

func From(userPermissions any) (p Permissions) {
	ret := Permissions{}
	s := reflect.ValueOf(userPermissions)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	for i := Admin; i < count; i++ {
		v := s.FieldByName(Permission(i).String())
		if v.Bool() {
			ret[Permission(i)] = struct{}{}
		}
	}
	return ret
}

func All() Permissions {
	ret := make(Permissions, count)
	for i := int(Admin); i < int(count); i++ {
		ret[Permission(i)] = struct{}{}
	}
	return ret
}
