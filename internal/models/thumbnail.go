package models

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
