package app

import (
	controller "github.com/cadeusept/thumbnail-loader/internal/services/downloader/controller/grpc"
	downloader_repository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository/sqlite"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// TODO
	cfg := sqlite.Config{
		DBName: "__test.db",
	}

	rSQLite, err := sqlite.NewSqliteDB(cfg)
	if err != nil {
		logrus.Fatalf("error creating database: %v", err)
		return
	}

	repo := downloader_repository.NewDownloaderRepo(rSQLite)

	uc := usecase.NewDownloadUseCase(repo)

	grpcServer := grpc.NewServer()

	s := controller.NewDownloadServerGRPC(grpcServer, uc)

	url := "http://127.0.0.1:9091"

	if err := s.Start(url); err != nil {
		logrus.Fatalf("error during downloader serve: %v", err)
	}
}
