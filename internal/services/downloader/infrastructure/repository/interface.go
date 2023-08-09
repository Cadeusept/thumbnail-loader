package downloader_repository

import "github.com/cadeusept/thumbnail-loader/internal/models"

type DownloadRepoI interface {
	CacheThumbnail(t models.Thumbnail) (int, error)
	GetThumbnail(urlHash string) (string, error)
}
