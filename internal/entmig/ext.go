package entmig

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	atlas "ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqlclient"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// EnvDevDatabase is the env variable for selecting a dev database.
const EnvDevDatabase = "ENT_DEV_DATABASE"

// Option configures the generation of versioned migrations.
type Option func(*Extension)

// Extension extends Ent's code-generation to support the generation of
// versioned migrations. To use this extension, add it to the project's
// entc.go file and then run code-gen with an additional argument `migrate`:
//
//	go run entc.go migrate
//
// Optionally, supply a name for the versioned migration:
//
//	go run entc.go migrate name_for_migration.
type Extension struct {
	entc.DefaultExtension
	devdb string
	opts  []schema.MigrateOption
}

// New returns a new Extension.
func New(opts ...Option) *Extension {
	x := &Extension{
		devdb: os.Getenv(EnvDevDatabase),
		opts: []schema.MigrateOption{
			schema.WithFormatter(atlas.DefaultFormatter),
		},
	}
	for _, opt := range opts {
		opt(x)
	}
	if x.devdb == "" {
		panic("entmig: must set dev db, use entmig.WithDevDatabase() or set " + EnvDevDatabase)
	}
	return x
}

// WithDevDatabase configures Extension to use the dev database with
// the provided url.
func WithDevDatabase(u string) Option {
	return func(x *Extension) {
		x.devdb = u
	}
}

// WithMigrationOpts configures Extension to use the provided MigrateOptions.
func WithMigrationOpts(opts ...schema.MigrateOption) Option {
	return func(x *Extension) {
		x.opts = append(x.opts, opts...)
	}
}

// Hooks returns the code-gen hook that generates versioned migrations.
func (e *Extension) Hooks() []gen.Hook {
	return []gen.Hook{e.hook}
}

// hook returns a generator that maintains the versioned migrations directory.
func (e *Extension) hook(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(graph *gen.Graph) error {
		if err := e.Generate(graph); err != nil {
			return err
		}
		return next.Generate(graph)
	})
}

// Generate generates the versioned migrations.
func (e *Extension) Generate(graph *gen.Graph) error {
	if len(os.Args) < 2 || os.Args[1] != "migrate" {
		return nil
	}
	var name string
	if len(os.Args) > 2 {
		name = os.Args[2]
	}
	client, err := sqlclient.Open(context.Background(), e.devdb)
	if err != nil {
		return fmt.Errorf("connecting to dev db: %w", err)
	}
	defer client.Close()
	tables, err := graph.Tables()
	if err != nil {
		return err
	}
	dir, err := e.migdir(graph.Target)
	fmt.Println("DIR:", dir)
	if err != nil {
		return err
	}
	opts := append(e.opts,
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
	)
	return schema.Diff(context.Background(), e.devdb, name, tables, opts...)
}

func (e *Extension) migdir(tgt string) (*atlas.LocalDir, error) {
	md := filepath.Join(tgt, "migrations")
	if _, err := os.Stat(md); os.IsNotExist(err) {
		if err := os.MkdirAll(md, os.ModePerm); err != nil {
			return nil, err
		}
	}
	return atlas.NewLocalDir(md)
}
