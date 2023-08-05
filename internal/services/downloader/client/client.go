package client

import downloader_proto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"

type DownloadClientGRPC struct {
	downloadClient downloader_proto.DownloaderServiceClient
}

func NewDownloadClientGRPC(c downloader_proto.DownloaderServiceClient) *DownloadClientGRPC {
	return &DownloadClientGRPC{
		downloadClient: c,
	}
}
