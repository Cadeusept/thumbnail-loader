package models

// Thumbnail имеет три поля которые получают корректный url и текущие имена файлов
/*
type Thumbnail struct {
	VideoID        string
	Link           string
	FileName       []string
	ThumbnailsDir  string
	ThumbnailsName string
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		VideoID:        "",
		Link:           "",
		FileName:       []string{},
		ThumbnailsDir:  "./thumbnails",
		ThumbnailsName: "",
	}
}
*/

type Thumbnail struct {
	Id      int    `db:"id"`
	UrlHash string `binding:"required"`
	Picture string `binding:"required"`
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		Id:      0,
		UrlHash: "",
		Picture: "",
	}
}
