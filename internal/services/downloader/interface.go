package downloader

import (
	"github.com/cadeusept/thumbnail-loader/internal/models"
	downloader_repository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
)

type DownloadUseCaseI interface {
	DownloadThumbnail(t *models.Thumbnail) (string, error)
}

type DownloadRepoI interface {
	CacheThumbnail(t downloader_repository.Thumbnail) (int, error)
	GetThumbnail(t *models.Thumbnail) (string, error)
}
