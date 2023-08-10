package usecase

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/entity"
	downloaderRepository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
	log "github.com/sirupsen/logrus"
)

const salt = "jzkcj2d324if04r0kc"

type downloadUseCase struct {
	sqliteRepo downloaderRepository.DownloadRepoI
}

func NewDownloadUseCase(r downloaderRepository.DownloadRepoI) *downloadUseCase {
	return &downloadUseCase{
		sqliteRepo: r,
	}
}

func (uc *downloadUseCase) DownloadThumbnail(url string) ([]byte, error) {
	var err error
	t := entity.NewThumbnail()
	t.Link = url
	t.VideoID, err = entity.FindVideoID(url)
	t.IdHash = generateIdHash(t.VideoID)
	if err != nil {
		return nil, fmt.Errorf("error finding video: %w", err)
	}

	// get cache
	log.Info("searching in cache")
	mt := &models.Thumbnail{
		UrlHash: t.IdHash,
		Picture: "",
	}
	mt.Picture, err = uc.sqliteRepo.GetThumbnail(mt.UrlHash)
	if err != nil {
		log.Infof("not found in cache: %v", err.Error())
	} else {
		log.Info("thumbnail found in cache")
		picture, err := os.ReadFile(mt.Picture)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}
		log.Info("thumbnail successfully sent to client")
		return picture, nil
	}

	log.Info("downloading thumbnail from youtube")
	picturePath, err := entity.DownloadThumbnail(t, url)
	if err != nil {
		return nil, fmt.Errorf("error downloading thumbnail: %w", err)
	}

	// Записываем кэш
	log.Info("saving thumbnail in cache")
	mt.Picture = picturePath
	_, err = uc.sqliteRepo.CacheThumbnail(*mt)
	if err != nil {
		log.Errorf("error creating cache: %v", err.Error())
	}

	picture, err := os.ReadFile(picturePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	log.Info("thumbnail successfully sent to client")

	return picture, nil
}

func generateIdHash(id string) string {
	hash := sha1.New()
	hash.Write([]byte(id))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
