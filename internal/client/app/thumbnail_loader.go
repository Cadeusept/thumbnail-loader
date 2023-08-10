package main

import (
	"context"
	"sync"

	clientGRPC "github.com/cadeusept/thumbnail-loader/internal/client/infrastructure/downloaderGRPC"
	"github.com/cadeusept/thumbnail-loader/internal/client/utils"
	downloaderProto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var asyncFlag bool
var rootCmd = cobra.Command{
	Use:     "thumbnail_loader",
	Version: "v1.0.0",
	Short:   "It's simple thumbnail loader",
	Long:    "You can load thumbnails one by one or altogether asynchroniously",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := ":9091"
		conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("error connecting to gRPC server: %v", err.Error())
			return
		}

		defer func() {
			if err := conn.Close(); err != nil {
				log.Errorf("error closing connection: %v", err.Error())
			} else {
				log.Info("gRPC connection succesfully closed")
			}
		}()

		log.Infof("gRPC server successfully connected")

		clientGRPC := clientGRPC.NewDownloadClientGRPC(downloaderProto.NewDownloaderServiceClient(conn))

		var wg sync.WaitGroup
		if asyncFlag {
			// Send async
			wg.Add(1)
			go clientGRPC.DownloadThumbnailsAsync(context.Background(), utils.RemoveDuplicateStr(args), &wg)
		} else {
			// Send sync
			wg.Add(1)
			go clientGRPC.DownloadThumbnailsSync(context.Background(), utils.RemoveDuplicateStr(args), &wg)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&asyncFlag, "async", "a", false, "Downloads thumbnails asynchroniously")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err.Error())
	}
}
