package entity

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// создаёт папку, если она существует, не делает ничего
func CreateFolder(thumbnailsDir string) error {
	err := os.MkdirAll(thumbnailsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create thumbnails folder: %v", err)
	}
	return nil
}

// createFile создаёт и сохраняет jpg thumbnail
// по умолчанию папка "downloadedThumbnails"
func CreateFile(thumbnailsName string) (*os.File, error) {

	// create file with auto set in the name's last number
	createdFile, err := os.Create(thumbnailsName)
	if err != nil {
		logrus.Errorf("%v", err)
		return nil, fmt.Errorf("can't create file: %v", err)
	}

	return createdFile, nil
}

// writeFile write response body from valid url
// at the created jpg thumbnail file.
func WriteFile(readyFile *os.File, picture string) error {
	r := strings.NewReader(picture)
	_, err := io.Copy(readyFile, r)
	if err != nil {
		log.Println(err)
		return err
	}
	readyFile.Close()
	logrus.Print("file written")
	return nil
}

func (t *Thumbnail) WalkFunc(path string, info os.FileInfo, err error) error {
	if info.Name() != "thumbnails" {
		t.FileName = append(t.FileName, info.Name())
	}

	return nil
}

// setNameDigit получает последнее имя файла в директории
// затем ставит следкющее число в имени файла
func SetNameDigit(inputArr []string) string {

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

		digitsCounter, err := strconv.Atoi(numbers)
		if err != nil {
			logrus.Errorln("String to Int Atoi conversion error!", err)
			return ""
		}

		digitsCounter++
		return strconv.Itoa(digitsCounter)
	}
	return "0"
}

func (t *Thumbnail) SetThumbnailName() string {
	return t.ThumbnailsDir + "/thumbnail_" + SetNameDigit(t.FileName) + ".jpg"
}
