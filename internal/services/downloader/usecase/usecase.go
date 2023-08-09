package usecase

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/entity"
	downloader_repository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
	"github.com/sirupsen/logrus"
)

const salt = "jzkcj2d324if04r0kc"

type downloadUseCase struct {
	sqliteRepo downloader_repository.DownloadRepoI
}

func NewDownloadUseCase(r downloader_repository.DownloadRepoI) *downloadUseCase {
	return &downloadUseCase{
		sqliteRepo: r,
	}
}

func (uc *downloadUseCase) DownloadThumbnail(url string) (string, error) {
	var err error
	t := entity.NewThumbnail()
	t.Link = url
	t.VideoID, err = entity.FindVideoID(url)
	t.IdHash = generateIdHash(t.VideoID)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
		return "", err
	}

	// get cache
	mt := &models.Thumbnail{
		UrlHash: t.IdHash,
		Picture: "",
	}
	mt.Picture, err = uc.sqliteRepo.GetThumbnail(mt.UrlHash)
	if err != nil {
		logrus.Printf("not found in cache: %v", err.Error())
	} else {
		return mt.Picture, nil
	}

	picturePath, err := entity.DownloadThumbnail(t, url)
	if err != nil {
		logrus.Fatalf("error downloading thumbnail: %s", err)
		return "", err
	}

	// create cache
	mt.Picture = picturePath
	_, err = uc.sqliteRepo.CacheThumbnail(*mt)
	if err != nil {
		logrus.Errorf("error creating cache: %v", err.Error())
	}

	picture, err := os.ReadFile(picturePath)
	if err != nil {
		logrus.Fatalf("error reading file %s: %v", picturePath, err.Error())
		return "", nil
	}

	return string(picture), nil
}

func generateIdHash(id string) string {
	hash := sha1.New()
	hash.Write([]byte(id))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
