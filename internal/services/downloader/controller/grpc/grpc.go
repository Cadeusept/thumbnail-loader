package controller

import (
	"context"
	"fmt"
	"net"

	"github.com/cadeusept/thumbnail-loader/internal/services/downloader"
	downloaderProto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type DownloadServerGRPC struct {
	downloaderProto.UnimplementedDownloaderServiceServer

	grpcServer *grpc.Server
	downloadUC downloader.DownloadUseCaseI
}

func NewDownloadServerGRPC(srv *grpc.Server, uc downloader.DownloadUseCaseI) *DownloadServerGRPC {
	return &DownloadServerGRPC{
		grpcServer: srv,
		downloadUC: uc,
	}
}

func (s DownloadServerGRPC) Start(url string) error {
	lis, err := net.Listen("tcp", url)
	if err != nil {
		return err
	}

	downloaderProto.RegisterDownloaderServiceServer(s.grpcServer, s)

	log.Info("server successfully started")

	return s.grpcServer.Serve(lis)
}

func (s DownloadServerGRPC) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	s.grpcServer.GracefulStop()
	return nil
}

func (s DownloadServerGRPC) DownloadThumbnail(ctx context.Context, req *downloaderProto.DownloadTRequest) (*downloaderProto.DownloadTResponse, error) {
	url := req.GetLink()

	picture, err := s.downloadUC.DownloadThumbnail(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading thumbnail: %w", err)
	}

	return &downloaderProto.DownloadTResponse{
		Picture: picture,
	}, nil
}
