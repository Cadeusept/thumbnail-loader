package sqlite

import (
	"fmt"

	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/jmoiron/sqlx"
)

type PicturesCache interface {
	Create(t models.Thumbnail) (int, error)
	GetPicture(urlHash string) (string, error)
}

type ThumbnailCacheSqlite struct {
	db *sqlx.DB
}

func NewThumbnailCacheSqlite(db *sqlx.DB) *ThumbnailCacheSqlite {
	return &ThumbnailCacheSqlite{
		db: db,
	}
}

func (r *ThumbnailCacheSqlite) Create(t models.Thumbnail) (int, error) {
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
