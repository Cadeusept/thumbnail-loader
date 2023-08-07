package usecase

import (
	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/entity"
	"github.com/sirupsen/logrus"
)

type downloadUseCase struct {
	sqliteRepo *downloader.DownloadRepoI
}

func NewDownloadUseCase(r *downloader.DownloadRepoI) *downloadUseCase {
	return &downloadUseCase{
		sqliteRepo: r,
	}
}

func (uc *downloadUseCase) DownloadThumbnail(t *models.Thumbnail) (string, error) {
	// TODO: check cache

	picture, err := entity.DownloadThumbnail(t)
	if err != nil {
		logrus.Fatalf("error downloading thumbnail: %s", err)
		return "", err
	}

	// TODO: do cache

	return picture, nil
}
