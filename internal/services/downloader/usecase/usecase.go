package usecase

import "github.com/cadeusept/thumbnail-loader/internal/services/downloader"

type downloadUseCase struct {
	sqliteRepo *downloader.DownloadRepoI
}

func NewDownloadUseCase(r *downloader.DownloadRepoI) *downloadUseCase {
	return &downloadUseCase{
		sqliteRepo: r,
	}
}

func (uc *downloadUseCase) DownloadThumbnail() {

}
