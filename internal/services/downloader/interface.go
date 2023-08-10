package downloader

type DownloadUseCaseI interface {
	DownloadThumbnail(url string) ([]byte, error)
}
