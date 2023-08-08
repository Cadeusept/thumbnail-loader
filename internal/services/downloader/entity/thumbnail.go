package entity

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const invalidSymbols string = "?&/<%="

// findVideoID получает из ссылки id видеоролика и возвращает его
// также проверяет длину id и содержание в id невалидных символов
func findVideoID(url string) (string, error) {
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
		logrus.Printf("unable to get max resolution: code %v\n", resp.StatusCode)

		// пробуем получить картинку с меньшим разрешением
		resp, err = http.Get(vi + videoId + resHQ)

		if err != nil || resp.StatusCode != 200 {
			logrus.Fatalf("error getting picture: %v\n", err)
			return nil, err
		}
	}

	return resp, nil
}

func DownloadThumbnail(url string) (string, error) {
	// thumbnailsDir := "../downloadedThumbnails"

	videoId, err := findVideoID(url)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
	}
	/*
		err = createFolder("../downloadedThumbnails")
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}

		// Walk walks thumbnails dir and save file names
		err = filepath.Walk(thumbnailsDir, t.walkFunc)
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}

		readyFile, errCreate := createFile(setThumbnailName())
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}

		errWrite := writeFile(readyFile, getURLResponse(videoId))
		if err != nil {
			logrus.Fatalf("error finding video: %s", err.Error())
		}
	*/
	picture, err := getURLResponse(videoId)
	if err != nil {
		logrus.Fatalf("error finding video: %s", err.Error())
		return "", err
	}

	return fmt.Sprint(picture.Body), nil
}
