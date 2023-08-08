package sqlite

import "github.com/jmoiron/sqlx"

const thumbnailsCacheTable = "thumbnails_cache"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "__test.db")

	if err != nil {
		return nil, err
	}

	return db, nil
}
