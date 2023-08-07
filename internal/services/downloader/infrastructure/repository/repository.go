package downloader_repository

import (
	"github.com/jmoiron/sqlx"
)

type picturesCache interface {
	Create(t Thumbnail) (int, error)
	GetPicture(urlHash string) (string, error)
}

type DownloaderRepo struct {
	picturesCache
}

func NewDownloaderRepo(db *sqlx.DB) *DownloaderRepo {
	return &DownloaderRepo{
		picturesCache: NewThumbnailCacheSqlite(db),
	}
}

func (r *DownloaderRepo) GetThumbnail(urlHash string) (string, error) {
	return r.picturesCache.GetPicture(urlHash)
}

func (r *DownloaderRepo) CacheThumbnail(t Thumbnail) (int, error) {
	return r.picturesCache.Create(t)
}
