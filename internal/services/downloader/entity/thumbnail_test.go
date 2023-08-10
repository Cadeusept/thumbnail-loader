package entity

import (
	"path/filepath"
	"testing"
)

func TestFindVideoID(t *testing.T) {
	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJQSBx5i4",
		"https://www.youtube.com/watch?v=65AB2pMCj4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/ZnhquCll3uQ",
		"https://www.youtube.com/embed/5eNieKeLBLE",
		"https://youtu.be/6k1oE2y7NIo?t=31",
	}

	for i, url := range urlList {
		id, err := FindVideoID(url)
		if err != nil {
			// fmt.Println("Corrupted Video URL:", tmb.VideoID)
			t.Errorf("Must be nil return! Link %s is invalid!", url)
		}

		switch i {
		case 0:
			if id != "N2wJQSBx5i4" {
				t.Errorf("%s must be N2wJQSBx5i4", id)
			}
		case 1:
			if id != "65AB2pMCj4I" {
				t.Errorf("%s must be 65AB2pMCj4I", id)
			}
		case 2:
			if id != "ZnhquCll3uQ" {
				t.Errorf("%s must be ZnhquCll3uQ", id)
			}
		case 3:
			if id != "5eNieKeLBLE" {
				t.Errorf("%s must be 5eNieKeLBLE", id)
			}
		case 4:
			if id != "6k1oE2y7NIo" {
				t.Errorf("%s must be 6k1oE2y7NIo", id)
			}
		}
	}
}

func TestFindVideoWrongID(t *testing.T) {
	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJ" + "?" + "SBx5i4",
		"https://www.youtube.com/watch?v=65AB" + "&" + "pMCj4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/Znh+" + "<" + "uC" + "%" + "l3uQ",
		"https://www.youtube.com/embed/5e" + "&" + "iKe" + "/" + "BLE",
	}

	for _, url := range urlList {
		_, err := FindVideoID(url)
		if err == nil {
			t.Errorf("Must be non nil return! Link ID %s is invalid!", url)
		}
	}
}

func TestFindVideoWrongIDLength(t *testing.T) {
	urlList := []string{
		"https://www.youtube.com/watch?v=N2wJQ",
		"https://www.youtube.com/watch?v=65AB2p4I&index=3&list=LLdt97678HxmYdM0DyZ847Uw",
		"https://youtu.be/ZnhquCl",
		"https://www.youtube.com/embed/5eNi",
	}

	for _, url := range urlList {
		_, err := FindVideoID(url)
		if err == nil {
			t.Errorf("Must be non nil return! Link ID length %s is invalid!", url)
		}
	}
}

func TestWalkFunc(t *testing.T) {

	tmb := new(Thumbnail)

	err := filepath.Walk(TestDir, tmb.walkFunc)
	if err != nil {
		t.Error("Walk Func Test Failed. Must retrun nil!")
	}
}

func TestSetNameDigit(t *testing.T) {

	namesList := []string{"thumbnail_0.jpg", "thumbnail_1.jpg", "thumbnail_2.jpg", "thumbnail_3.jpg", "thumbnail_4.jpg"}
	lastName := namesList[4]

	if setNameDigit(namesList) != string(lastName[10]+1) {
		t.Error("Must be string number return!")
	}
}

func TestSetNameDigitZeroList(t *testing.T) {

	namesList := []string{}

	if setNameDigit(namesList) != "0" {
		t.Error("Must be zero string number return!")
	}
}

func TestSetNameDigitWrongNumbers(t *testing.T) {

	namesList := []string{"<<", "<<<<<"}

	out := setNameDigit(namesList)

	if out != "1" {
		t.Errorf("Wrong number symbols! %v", out)
	}
}

func TestSetThumbnailNameVaild(t *testing.T) {

	tmb := NewThumbnail()

	result := tmb.setThumbnailName()

	if len(result) == 0 {
		t.Errorf("Empty file name! %s\n", result)
	}
}

func TestNewThumbnailVaild(t *testing.T) {

	videoID := ""
	link := ""
	thumbnailsName := ""

	result := NewThumbnail()

	if result == nil {
		t.Errorf("Should be non nil struct! %v\n", result)
	}

	if result.VideoID != videoID || result.Link != link || result.FileName == nil || result.ThumbnailsDir == "" || result.ThumbnailsName != thumbnailsName {
		t.Errorf("Should be valid struct! %v\n", result)
	}
	removeTestDir(TestDir)
}
