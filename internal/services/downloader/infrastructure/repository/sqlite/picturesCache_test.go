package sqlite

import (
	"testing"

	"github.com/cadeusept/thumbnail-loader/internal/models"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestDelivery_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewThumbnailCacheSqlite(db)

	tests := []struct {
		name    string
		repo    *ThumbnailCacheSqlite
		thmb    models.Thumbnail
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			repo: repo,
			thmb: models.Thumbnail{
				UrlHash: "url_hash",
				Picture: "picture",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO thumbnails_cache").WithArgs("url_hash", "picture").WillReturnRows(rows)
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "EmptyField_UrlHash",
			repo: repo,
			thmb: models.Thumbnail{
				UrlHash: "",
				Picture: "picture",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO thumbnails_cache").WithArgs("url_hash", "picture").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "EmptyField_Picture",
			repo: repo,
			thmb: models.Thumbnail{
				UrlHash: "url_hash",
				Picture: "",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO thumbnails_cache").WithArgs("url_hash", "picture").WillReturnRows(rows)
			},
			wantErr: true,
		},
		{
			name: "EmptyFields",
			repo: repo,
			thmb: models.Thumbnail{
				UrlHash: "",
				Picture: "",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO thumbnails_cache").WithArgs("url_hash", "picture").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.repo.Create(tt.thmb)
			if (err != nil) != tt.wantErr {
				t.Errorf("Got error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf("Got = %v, want %v", got, tt.want)
			}
		})
	}
}
