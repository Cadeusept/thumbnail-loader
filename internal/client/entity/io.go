package entity

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// CreateFolder создаёт папку, если она существует, или не делает ничего
func CreateFolder(thumbnailsDir string) error {
	err := os.MkdirAll(thumbnailsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create thumbnails folder: %w", err)
	}
	return nil
}

// CreateFile создаёт и сохраняет jpg thumbnail
// по умолчанию папка "downloadedThumbnails"
func CreateFile(thumbnailsName string) (*os.File, error) {

	// создаёт файл с автоматически выставленным именем
	createdFile, err := os.Create(thumbnailsName)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %w", err)
	}

	return createdFile, nil
}

// WriteFile пишет картинку из строки в созданный файл
func WriteFile(readyFile *os.File, picture []byte) error {
	defer readyFile.Close()

	r := bytes.NewReader(picture)
	_, err := io.Copy(readyFile, r)
	if err != nil {
		return fmt.Errorf("can't copy into file: %w", err)
	}

	logrus.Print("file written")
	return nil
}

// WalcFunc это фуннкция обхода директории
func (t *Thumbnail) WalkFunc(path string, info os.FileInfo, err error) error {
	if info.Name() != "thumbnails" {
		t.FileName = append(t.FileName, info.Name())
	}

	return nil
}

// SetNameDigit получает последнее имя файла в директории
// затем ставит следующее число в имени файла
func SetNameDigit(inputArr []string) string {
	var err error
	var digitsCounter int

	if len(inputArr) > 0 {

		// сортирует и получает последний номер файла
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
				logrus.Errorln("String to Int Atoi conversion error! ", err)
				return ""
			}
		}

		digitsCounter++
		return strconv.Itoa(digitsCounter)
	}
	return "0"
}

// SetThumbnailName формирует и возвращает имя thumbnail'а
func (t *Thumbnail) SetThumbnailName() string {
	return t.ThumbnailsDir + "/thumbnail_" + SetNameDigit(t.FileName) + ".jpg"
}
