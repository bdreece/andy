//go:generate make -C ../.. generate/sql
package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DSN        string `yaml:"dsn"`
	SchemaPath string `yaml:"schema_path"`
}

func NewDBTX(cfg *Config) (DBTX, error) {
    db, err := sql.Open("sqlite3", cfg.DSN)
    if err != nil {
        return nil, err
    }

    if err = applyConfig(context.Background(), db, cfg); err != nil {
        return nil, err
    }

    return db, nil
}
