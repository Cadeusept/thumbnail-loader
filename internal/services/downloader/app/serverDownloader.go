package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	controller "github.com/cadeusept/thumbnail-loader/internal/services/downloader/controller/grpc"
	downloaderRepository "github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/infrastructure/repository/sqlite"
	"github.com/cadeusept/thumbnail-loader/internal/services/downloader/usecase"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	cfg := sqlite.Config{
		DBPath: "./../infrastructure/repository/sqlite/__test.db",
	}

	rSQLite, err := sqlite.NewSqliteDB(cfg)
	if err != nil {
		log.Fatalf("error creating database: %v", err)
	}

	log.Info("successfully connected to database")

	// закрытие базы данных
	defer func() {
		if err = rSQLite.DB.Close(); err != nil {
			log.Errorf("error during DB connection closure: %v", err.Error())
		} else {
			log.Info("successfully disconnected database")
		}
	}()

	repo := downloaderRepository.NewDownloaderRepo(rSQLite)

	uc := usecase.NewDownloadUseCase(repo)

	grpcServer := grpc.NewServer()

	s := controller.NewDownloadServerGRPC(grpcServer, uc)

	url := ":9091"

	if err := s.Start(url); err != nil {
		log.Fatalf("error during downloader service serve: %v", err)
	}

	defer func() {
		if err = s.Shutdown(context.Background()); err != nil {
			log.Errorf("error during gRPC server shutdown: %v", err.Error())
		}
	}()

	// graceful shutdown
	c_quit := make(chan os.Signal, 1)
	signal.Notify(c_quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c_quit

	log.Printf("signal: %s", sig.String())
}
