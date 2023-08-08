package controller

import (
	"context"
	"net"

	"github.com/cadeusept/thumbnail-loader/internal/services/downloader"
	downloader_proto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type DownloadServerGRPC struct {
	downloader_proto.UnimplementedDownloaderServiceServer

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

	downloader_proto.RegisterDownloaderServiceServer(s.grpcServer, s)

	return s.grpcServer.Serve(lis)
}

func (s DownloadServerGRPC) DownloadThumbnail(ctx context.Context, req *downloader_proto.DownloadTRequest) (*downloader_proto.DownloadTResponse, error) {
	// unpack request & download thumbnail
	url := req.GetLink()

	picture, err := s.downloadUC.DownloadThumbnail(url)
	if err != nil {
		logrus.Fatalf("error downloading thumbnail: %s", err)
		return nil, err
	}

	return &downloader_proto.DownloadTResponse{
		Picture: picture,
	}, nil
}
