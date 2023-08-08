package downloader_repository

import (
	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository/sqlite"
	"github.com/jmoiron/sqlx"
)

type DownloaderRepo struct {
	sqlite.PicturesCache
}

func NewDownloaderRepo(db *sqlx.DB) *DownloaderRepo {
	return &DownloaderRepo{
		PicturesCache: sqlite.NewThumbnailCacheSqlite(db),
	}
}

func (r *DownloaderRepo) GetThumbnail(urlHash string) (string, error) {
	return r.PicturesCache.GetPicture(urlHash)
}

func (r *DownloaderRepo) CacheThumbnail(t models.Thumbnail) (int, error) {
	return r.PicturesCache.Create(t)
}
