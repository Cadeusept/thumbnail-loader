package entity

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// создаёт папку, если она существует, не делает ничего
func createFolder(thumbnailsDir string) error {
	err := os.MkdirAll(thumbnailsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create thumbnails folder: %v", err)
	}
	return nil
}

// createFile создаёт и сохраняет jpg thumbnail
// по умолчанию папка "downloadedThumbnails"
func createFile(thumbnailsName string) (*os.File, error) {

	// создаёт файл с автоматически выставленным номером в имени
	createdFile, err := os.Create(thumbnailsName)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %v", err)
	}

	return createdFile, nil
}

// writeFile пишет тело запроса
// в созданный файл
func writeFile(readyFile *os.File, resp *http.Response) error {
	defer resp.Body.Close()
	defer readyFile.Close()

	_, err := io.Copy(readyFile, resp.Body)
	if err != nil {
		// log.Println(err)
		return fmt.Errorf("failed to read file: %w", err)
	}
	logrus.Print("file written")
	return nil
}

func (t *Thumbnail) walkFunc(path string, info os.FileInfo, err error) error {
	if info.Name() != "thumbnails" {
		t.FileName = append(t.FileName, info.Name())
	}

	return nil
}

// setNameDigit получает последнее имя файла в директории
// затем ставит следкющее число в имени файла
func setNameDigit(inputArr []string) string {
	var err error
	var digitsCounter int

	if len(inputArr) > 0 {

		// sort and get last thumbnail filename
		sort.Slice(inputArr, func(i, j int) bool {

			startA := strings.Index(inputArr[i], "_") + 1
			endA := strings.Index(inputArr[i], ".")

			startB := strings.Index(inputArr[j], "_") + 1
			endB := strings.Index(inputArr[j], ".")

			if startA == -1 || endA == -1 || startB == -1 || endB == -1 {
				return false
			}

			numberA := inputArr[i][startA:endA]
			numberB := inputArr[j][startB:endB]

			numA, _ := strconv.Atoi(numberA)
			numB, _ := strconv.Atoi(numberB)

			return numA < numB
		})

		lastFile := inputArr[len(inputArr)-1]

		var numbers string
		for i := 0; i < len(lastFile); i++ {

			elem := lastFile[i]
			if '0' <= elem && elem <= '9' {
				numbers += strconv.Itoa(int(elem) - '0')
			}
		}

		if numbers == "" {
			digitsCounter = 0
		} else {
			digitsCounter, err = strconv.Atoi(numbers)
			if err != nil {
				logrus.Errorln("String to Int Atoi conversion error!", err)
				return ""
			}
		}

		digitsCounter++
		return strconv.Itoa(digitsCounter)
	}
	return "0"
}

func (t *Thumbnail) setThumbnailName() string {
	return t.ThumbnailsDir + "/thumbnail_" + setNameDigit(t.FileName) + ".jpg"
}
