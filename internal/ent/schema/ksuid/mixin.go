package ksuid

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

func MixinWithPrefix(prefix string) *Mixin {
	return &Mixin{
		prefix: prefix,
	}
}

type Mixin struct {
	mixin.Schema
	prefix string
}

func (m Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ID("")).
			DefaultFunc(func() ID { return MustNew(m.prefix) }),
	}
}

type Annotation struct {
	Prefix string
}

func (a Annotation) Name() string {
	return "KSUID"
}

func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: m.prefix},
	}
}
