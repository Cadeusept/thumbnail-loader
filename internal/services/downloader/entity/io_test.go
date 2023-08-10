package entity

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"testing"
)

const TestDir = "./../downloadedThumbnails"

func TestCreateFile(t *testing.T) {

	tmb := &Thumbnail{}

	_ = createFolder(TestDir)
	thumbnailsName := TestDir[2:] + "/thumbnail_" + setNameDigit(tmb.FileName) + ".jpg"

	mf, err := NewMuxFile(thumbnailsName)
	if err != nil {
		t.Error("Wrong created file!")
	}
	mf.file.Close()
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

	err := createFolder(thumbnailsDir)
	if err == nil {
		t.Error("Expected nil return!")
	}
}

func TestCreateFileWrong(t *testing.T) {

	// thumbnailsDir := "./test_data"
	_ = createFolder(TestDir)
	var thumbnailsName string
	osName := runtime.GOOS

	if osName == "windows" {
		thumbnailsName = TestDir[2:] + "/<"
	} else if osName == "linux" {
		thumbnailsName = "/"
	}

	mf, err := NewMuxFile(thumbnailsName)
	mf.file.Close()
	defer os.Remove("./" + thumbnailsName)

	if err == nil {
		t.Error("incorrect file name. Must be nil return!")
	}
}

func TestWriteFile(t *testing.T) {

	mf, _ := NewMuxFile(TestDir + "/thumbnail_test2.jpg")
	defer mf.file.Close()
	defer os.Remove(TestDir + "/thumbnail_test2.jpg")

	resp, err := http.Get("https://www.youtube.com/watch?v=N2wJQSBx5i4")

	if writeFile(mf, resp) != nil {
		t.Errorf("Write file failed %v\n", err)
	}
	resp.Body.Close()
}

func TestWriteFileWrong(t *testing.T) {

	resp, _ := http.Get("https://www.youtube.com/watch?v=N2wJQSBx5i4")
	err := writeFile(nil, resp)

	if err == nil {
		t.Errorf("Write file failed %v\n", err)
	}
	resp.Body.Close()
}

// после прохода тестов удаляет тестовую директорию
func removeTestDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Println(err)
	}
}
