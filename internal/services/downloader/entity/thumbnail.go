package entity

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cadeusept/thumbnail-loader/internal/models"
	"github.com/sirupsen/logrus"
)

const invalidSymbols string = "?&/<%="

// findVideoID получает из ссылки id видеоролика и сохраняет в Thumbnail структуру
// также проверяет длину id и содержание в id невалидных символов
func findVideoID(t *models.Thumbnail) error {
	equalIndex := strings.Index(t.Link, "=")
	ampIndex := strings.Index(t.Link, "&")
	slash := strings.LastIndex(t.Link, "/")
	questionIndex := strings.Index(t.Link, "?")
	var id string

	if equalIndex != -1 {

		if ampIndex != -1 {
			id = t.Link[equalIndex+1 : ampIndex]
		} else if questionIndex != -1 && strings.Contains(t.Link, "?t=") {
			id = t.Link[slash+1 : questionIndex]
		} else {
			id = t.Link[equalIndex+1:]
		}

	} else {
		id = t.Link[slash+1:]
	}

	t.VideoID = id

	if strings.ContainsAny(id, invalidSymbols) {
		return errors.New("invalid characters in video id")
	}
	if len(id) < 10 {
		return errors.New("the video id must be at least 10 characters long")
	}
	return nil
}

// getUrlResponse получает и проверяет две картинки (с плохим и с хорошим разрешением)
// если "/maxresdefault.jpg" не существует или возвращает плохой код
// пробуем получить картинку в худшем или единственном качестве "/hqdefault.jpg".
func getURLResponse(t *models.Thumbnail) (*http.Response, error) {

	// два разрешения картинки
	const (
		vi     = "https://i.ytimg.com/vi/"
		resMax = "/maxresdefault.jpg"
		resHQ  = "/hqdefault.jpg"
	)

	resp, err := http.Get(vi + t.VideoID + resMax)

	if err != nil || resp.StatusCode != 200 {
		logrus.Printf("unable to get max resolution: code %v\n", resp.StatusCode)

		// пробуем получить картинку с меньшим разрешением
		resp, err = http.Get(vi + t.VideoID + resHQ)

		if err != nil || resp.StatusCode != 200 {
			logrus.Fatalf("error getting picture: %v\n", err)
			return nil, err
		}
	}

	return resp, nil
}

func DownloadThumbnail(t *models.Thumbnail) (string, error) {
	err := findVideoID(t)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
	}

	/*
		err = createFolder(t.thumbnailsDir)
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}


		// Walk walks thumbnails dir and save file names
		err = filepath.Walk(t.thumbnailsDir, t.walkFunc)
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}

		readyFile, errCreate := createFile(t.setThumbnailName())
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}

		errWrite := writeFile(readyFile, t.getURLResponse())
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}
	*/

	picture, err := getURLResponse(t)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
		return "", err
	}

	return fmt.Sprint(picture.Body), nil
}
