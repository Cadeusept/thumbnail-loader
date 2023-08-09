package sqlite

import "github.com/jmoiron/sqlx"

const thumbnailsCacheTable = "thumbnails_cache"

type Config struct {
	DBName string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", cfg.DBName)

	if err != nil {
		return nil, err
	}

	return db, nil
}
