package entity

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const invalidSymbols string = "?&/<%="

type Thumbnail struct {
	VideoID        string
	IdHash         string
	Link           string
	FileName       []string
	ThumbnailsDir  string
	ThumbnailsName string
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		VideoID:        "",
		IdHash:         "",
		Link:           "",
		FileName:       []string{},
		ThumbnailsDir:  "./../downloadedThumbnails",
		ThumbnailsName: "",
	}
}

// findVideoID получает из ссылки id видеоролика и возвращает его
// также проверяет длину id и содержание в id невалидных символов
func FindVideoID(url string) (string, error) {
	equalIndex := strings.Index(url, "=")
	ampIndex := strings.Index(url, "&")
	slash := strings.LastIndex(url, "/")
	questionIndex := strings.Index(url, "?")
	var id string

	if equalIndex != -1 {

		if ampIndex != -1 {
			id = url[equalIndex+1 : ampIndex]
		} else if questionIndex != -1 && strings.Contains(url, "?t=") {
			id = url[slash+1 : questionIndex]
		} else {
			id = url[equalIndex+1:]
		}

	} else {
		id = url[slash+1:]
	}

	if strings.ContainsAny(id, invalidSymbols) {
		return id, errors.New("invalid characters in video id")
	}
	if len(id) < 10 {
		return id, errors.New("the video id must be at least 10 characters long")
	}
	return id, nil
}

// getUrlResponse получает и проверяет две картинки (с плохим и с хорошим разрешением)
// если "/maxresdefault.jpg" не существует или возвращает плохой код
// пробуем получить картинку в худшем или единственном качестве "/hqdefault.jpg".
func getURLResponse(videoId string) (*http.Response, error) {

	// два разрешения картинки
	const (
		vi     = "https://i.ytimg.com/vi/"
		resMax = "/maxresdefault.jpg"
		resHQ  = "/hqdefault.jpg"
	)

	resp, err := http.Get(vi + videoId + resMax)

	if err != nil || resp.StatusCode != 200 {
		logrus.Errorf("unable to get max resolution: code %v\n", resp.StatusCode)

		// пробуем получить картинку с меньшим разрешением
		resp, err = http.Get(vi + videoId + resHQ)

		if err != nil || resp.StatusCode != 200 {
			return nil, fmt.Errorf("can't copy into file: %w", err)
		}
	}

	return resp, nil
}

func DownloadThumbnail(t *Thumbnail, url string) (string, error) {
	picture, err := getURLResponse(t.VideoID)
	if err != nil {
		return "", fmt.Errorf("error finding video: %w", err)
	}

	err = createFolder(t.ThumbnailsDir)
	if err != nil {
		return "", fmt.Errorf("error creating folder: %w", err)
	}

	// перебирает директорию thumbnails и сохраняет имена файлов
	err = filepath.Walk(t.ThumbnailsDir, t.walkFunc)
	if err != nil {
		return "", fmt.Errorf("error walking: %w", err)
	}

	fileName := t.setThumbnailName()
	readyFile, errCreate := NewMuxFile(fileName)
	if errCreate != nil {
		return "", fmt.Errorf("error creating file: %w", errCreate)
	}

	errWrite := writeFile(readyFile, picture)
	if errWrite != nil {
		return "", fmt.Errorf("error writing file: %w", errWrite)
	}

	return fileName, nil
}
