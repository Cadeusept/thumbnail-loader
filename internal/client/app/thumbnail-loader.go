package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	clientGRPC "github.com/cadeusept/thumbnail-loader/internal/client/infrastructure/downloaderGRPC"
	downloader_proto "github.com/cadeusept/thumbnail-loader/internal/services/downloader/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var asyncFlag bool
var rootCmd = cobra.Command{
	Use:     "thumbnail-loader",
	Version: "v1.0.0",
	Short:   "It's simple thumbnail loader",
	Long:    "You can load thumbnails one by one or altogether asynchroniously",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://127.0.0.1:9091"
		conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("error connecting to gRPC server: %v", err.Error())
			return
		}

		log.Infof("gRPC server successfully connected")

		// c := downloader_proto.NewDownloaderServiceClient(conn)
		clientGRPC := clientGRPC.NewDownloadClientGRPC(downloader_proto.NewDownloaderServiceClient(conn))

		if asyncFlag {
			// Send async
			go clientGRPC.DownloadThumbnailsAsync(context.Background(), args)
		} else {
			// Send sync
			go clientGRPC.DownloadThumbnailsSync(context.Background(), args)
		}

		// graceful shutdown
		c_quit := make(chan os.Signal, 1)
		signal.Notify(c_quit, syscall.SIGTERM, syscall.SIGINT)
		sig := <-c_quit

		log.Printf("catched signal: %s. App shutting down...", sig.String())
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&asyncFlag, "async", "a", false, "Downloads thumbnails async")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err.Error())
	}
}
