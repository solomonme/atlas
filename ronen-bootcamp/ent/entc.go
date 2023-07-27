//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/go-sql-driver/mysql"

	"playground/ronen-bootcamp/internal/entmig"
)

func main() {
	err := entc.Generate("./ent/schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureVersionedMigration, gen.FeaturePrivacy},
		},
		entc.Extensions(
			entmig.New(
				entmig.WithDevDatabase("mysql://root:pass@localhost:3307/test"),
				entmig.WithMigrationOpts(
					schema.WithGlobalUniqueID(true),
				),
			),
		))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
