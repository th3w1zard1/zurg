package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/debridmediamanager/zurg/internal"
	"github.com/debridmediamanager/zurg/internal/version"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "zurg",
		Short: "Zurg - Real-Debrid Media Manager",
		Long:  `A media manager for Real-Debrid that provides WebDAV and HTTP access to your torrents and downloads.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Setup signal handling
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Handle graceful shutdown
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				fmt.Println("\nReceived shutdown signal, gracefully shutting down...")
				cancel()
			}()

			// Start the main application
			if err := internal.MainApp(ctx); err != nil {
				log.Fatal(err)
			}
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: func(cmd *cobra.Command, args []string) {
			internal.ShowVersion()
		},
	}

	var clearDownloadsCmd = &cobra.Command{
		Use:   "clear-downloads",
		Short: "Clear all downloads from Real-Debrid",
		Run: func(cmd *cobra.Command, args []string) {
			internal.ClearDownloads()
		},
	}

	var clearTorrentsCmd = &cobra.Command{
		Use:   "clear-torrents", 
		Short: "Clear all torrents from Real-Debrid",
		Run: func(cmd *cobra.Command, args []string) {
			internal.ClearTorrents()
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(clearDownloadsCmd)
	rootCmd.AddCommand(clearTorrentsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}