package db

import (
	"fmt"

	_ "lybbrio/internal/ent/runtime"

	"lybbrio/internal/config"
	"lybbrio/internal/ent"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

func Open(conf *config.DatabaseConfig) (*ent.Client, error) {
	switch conf.Driver {
	case "sqlite3":
		return OpenSQLite(conf)
	case "mysql":
		return OpenMySQL(conf)
	case "postgres":
		return OpenPostgres(conf)
	case "test":
		return OpenTest()
	default:
		// Should be unreachable, as the config validation should catch this.
		return nil, fmt.Errorf("unknown database driver: %s", conf.Driver)
	}
}

func OpenSQLite(conf *config.DatabaseConfig) (*ent.Client, error) {
	db, err := sql.Open(dialect.SQLite, conf.DSN)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.SQLite, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func OpenMySQL(conf *config.DatabaseConfig) (*ent.Client, error) {
	db, err := sql.Open(dialect.MySQL, conf.DSN)
	if err != nil {
		return nil, err
	}
	if conf.MaxIdleConns > 0 {
		db.SetMaxIdleConns(conf.MaxIdleConns)
	}
	if conf.MaxOpenConns > 0 {
		db.SetMaxOpenConns(conf.MaxOpenConns)
	}
	if conf.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	}
	drv := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func OpenPostgres(conf *config.DatabaseConfig) (*ent.Client, error) {
	db, err := sql.Open("pgx", conf.DSN)
	if err != nil {
		return nil, err
	}

	if conf.MaxIdleConns > 0 {
		db.SetMaxIdleConns(conf.MaxIdleConns)
	}
	if conf.MaxOpenConns > 0 {
		db.SetMaxOpenConns(conf.MaxOpenConns)
	}
	if conf.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func OpenTest() (*ent.Client, error) {
	return ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
}
