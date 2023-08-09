package entity

type Thumbnail struct {
	FileName       []string
	ThumbnailsDir  string
	ThumbnailsName string
}

func NewThumbnail() *Thumbnail {
	return &Thumbnail{
		FileName:       []string{},
		ThumbnailsDir:  "./../../../downloadedThumbnails",
		ThumbnailsName: "",
	}
}
