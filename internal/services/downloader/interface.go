package downloader

import (
	"github.com/cadeusept/thumbnail-loader/internal/models"
)

type DownloadUseCaseI interface {
	DownloadThumbnail(url string) (string, error)
}

type DownloadRepoI interface {
	CacheThumbnail(t models.Thumbnail) (int, error)
	GetThumbnail(t *models.Thumbnail) (string, error)
}
