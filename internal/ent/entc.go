//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	templates := entgql.AllTemplates

	templates = append(templates, gen.MustParse(
		gen.NewTemplate("ksuid.tmpl").
			ParseFiles("internal/ent/schema/ksuid/template/ksuid.tmpl")),
	)

	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("gqlgen.yaml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("internal/graph/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.FeatureNames(
			"entql",
			"privacy",
			"schema/snapshot",
			"sql/upsert",
		),
		entc.Extensions(ex),
	}
	if err := entc.Generate("./internal/ent/schema", &gen.Config{Templates: templates}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
