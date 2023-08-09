package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	controller "github.com/cadeusept/thumbnail-loader/internal/services/downloader/controller/grpc"
	downloader_repository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository/sqlite"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

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

	// graceful shutdown
	c_quit := make(chan os.Signal, 1)
	signal.Notify(c_quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c_quit

	log.Printf("catched signal: %s. App shutting down...", sig.String())

	if err = s.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error during gRPC server shutdown: %v", err.Error())
	}

	if err = rSQLite.DB.Close(); err != nil {
		logrus.Errorf("error during DB connection closure: %v", err.Error())
	}
}
