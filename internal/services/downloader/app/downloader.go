package app

import (
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader"
	controller "github.com/cadeusept/thumbnail-loader/internal/services/downloader/controller/grpc"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// TODO
	var repo *downloader.DownloadRepoI

	uc := usecase.NewDownloadUseCase(repo)

	grpcServer := grpc.NewServer()

	s := controller.NewDownloadServerGRPC(grpcServer, uc)

	url := "http://127.0.0.1:9091"

	if err := s.Start(url); err != nil {
		logrus.Fatalf("error during downloader serve: %v", err)
	}
}
