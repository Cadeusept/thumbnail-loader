package downloader

type DownloadUseCaseI interface {
	DownloadThumbnail(url string) (string, error)
}
