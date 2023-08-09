package entity

import (
	"log"
	"os"
	"runtime"
	"testing"
)

const TestDir = "./../../../downloadedThumbnails"

func TestCreateFile(t *testing.T) {

	tmb := &Thumbnail{}

	_ = CreateFolder(TestDir)
	thumbnailsName := TestDir[2:] + "/thumbnail_" + SetNameDigit(tmb.FileName) + ".jpg"

	file, err := CreateFile(thumbnailsName)
	if err != nil {
		t.Error("Wrong created file!")
	}
	file.Close()
	os.Remove("./" + thumbnailsName)
}

func TestCreateWrongDirectory(t *testing.T) {

	var thumbnailsDir string

	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsDir = "/<"
	} else if osName == "linux" {
		thumbnailsDir = ""
	}

	err := CreateFolder(thumbnailsDir)
	if err == nil {
		t.Error("Expected nil return!")
	}
}

func TestCreateFileWrong(t *testing.T) {

	// thumbnailsDir := "./test_data"
	_ = CreateFolder(TestDir)
	var thumbnailsName string
	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsName = TestDir[2:] + "/<"
	} else if osName == "linux" {
		thumbnailsName = "/"
	}

	file, err := CreateFile(thumbnailsName)
	file.Close()
	defer os.Remove("./" + thumbnailsName)

	if err == nil {
		t.Error("incorrect file name. Must be nil return!")
	}
}

func TestWriteFile(t *testing.T) {

	file, _ := os.Create(TestDir + "/thumbnail_test2.jpg")
	defer file.Close()
	defer os.Remove(TestDir + "/thumbnail_test2.jpg")

	str := "test"

	if err := WriteFile(file, str); err != nil {
		t.Errorf("Write file failed %v\n", err)
	}
	removeTestDir(TestDir)
}

// после прохода тестов удаляет тестовую директорию
func removeTestDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Println(err)
	}
}
