package models

// Thumbnail имеет три поля которые получают корректный url и текущие имена файлов
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
