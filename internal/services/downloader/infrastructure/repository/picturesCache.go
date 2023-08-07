package downloader_repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Thumbnail struct {
	Id      int    `db:"id"`
	UrlHash string `binding:"required"`
	Picture string `binding:"required"`
}

type ThumbnailCacheSqlite struct {
	db *sqlx.DB
}

func NewThumbnailCacheSqlite(db *sqlx.DB) *ThumbnailCacheSqlite {
	return &ThumbnailCacheSqlite{
		db: db,
	}
}

func (r *ThumbnailCacheSqlite) Create(t Thumbnail) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (url_hash, picture) values ($1, $2) RETURNING id", thumbnailsCacheTable)
	row := r.db.QueryRow(query, t.UrlHash, t.Picture)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ThumbnailCacheSqlite) GetPicture(urlHash string) (string, error) {
	var picture string
	query := fmt.Sprintf("SELECT picture FROM %s WHERE url_hash=$1", thumbnailsCacheTable)
	err := r.db.Get(picture, query, urlHash)

	return picture, err
}
