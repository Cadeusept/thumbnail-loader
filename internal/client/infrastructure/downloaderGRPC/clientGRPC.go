package clientGRPC

import (
	"context"
	"path/filepath"
	"sync"

	entity "github.com/cadeusept/thumbnail-loader/internal/client/entity"
	downloaderProto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"
	"github.com/sirupsen/logrus"
)

type DownloadClientGRPC struct {
	downloadClient downloaderProto.DownloaderServiceClient
}

func NewDownloadClientGRPC(c downloaderProto.DownloaderServiceClient) *DownloadClientGRPC {
	return &DownloadClientGRPC{
		downloadClient: c,
	}
}

func (c *DownloadClientGRPC) DownloadThumbnail(ctx context.Context, t *entity.Thumbnail, url string) error {
	resp, err := c.downloadClient.DownloadThumbnail(context.Background(),
		&downloaderProto.DownloadTRequest{
			Link: url,
		})
	if err != nil {
		return err
	} else {

		// перебирает директорию thumbnails и сохраняет имена файлов
		err = filepath.Walk(t.ThumbnailsDir, t.WalkFunc)
		if err != nil {
			return err
		}

		fileName := t.SetThumbnailName()
		readyFile, errCreate := entity.CreateFile(fileName)
		if errCreate != nil {
			return err
		}

		errWrite := entity.WriteFile(readyFile, resp.Picture)
		if errWrite != nil {
			return err
		}
	}

	return nil
}

func (c *DownloadClientGRPC) DownloadThumbnailsSync(ctx context.Context, urls []string, wg *sync.WaitGroup) {
	t := entity.NewThumbnail()
	err := entity.CreateFolder(t.ThumbnailsDir)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
	}

	for _, v := range urls {
		err := c.DownloadThumbnail(ctx, t, v)
		if err != nil {
			logrus.Errorf("error downloading thumbnail %s: %v", v, err)
		}
	}
	wg.Done()
}

func (c *DownloadClientGRPC) DownloadThumbnailsAsync(ctx context.Context, urls []string, wg *sync.WaitGroup) {
	t := entity.NewThumbnail()
	err := entity.CreateFolder(t.ThumbnailsDir)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
	}

	for _, v := range urls {
		wg.Add(1)
		go func(url string) {
			err := c.DownloadThumbnail(ctx, t, url)
			if err != nil {
				logrus.Errorf("error downloading thumbnail %s: %v", url, err)
			}
			wg.Done()
		}(v)
	}
	wg.Done()
}
