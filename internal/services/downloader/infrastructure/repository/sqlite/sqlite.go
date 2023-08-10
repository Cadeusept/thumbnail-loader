package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const thumbnailsCacheTable = "thumbnails_cache"

type Config struct {
	DBPath string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", cfg.DBPath)

	if err != nil {
		return nil, fmt.Errorf("error conneccting to database: %w", err)
	}

	return db, nil
}
