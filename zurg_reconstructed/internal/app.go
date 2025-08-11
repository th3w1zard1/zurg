package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/debridmediamanager/zurg/internal/config"
	"github.com/debridmediamanager/zurg/internal/handlers"
	"github.com/debridmediamanager/zurg/internal/torrent"
	"github.com/debridmediamanager/zurg/pkg/logutil"
	"github.com/debridmediamanager/zurg/pkg/premium"
	"github.com/debridmediamanager/zurg/pkg/realdebrid"
	httpClient "github.com/debridmediamanager/zurg/pkg/http"
)

// MainApp is the main application entry point
func MainApp(ctx context.Context) error {
	// Load configuration
	cfg, err := config.LoadZurgConfig("")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize logger
	logger := logutil.NewLogger(false) // Development mode off for production
	defer logger.Sync()

	logger.Info("Starting Zurg",
		zap.String("version", "v0.9.3-final"),
		zap.String("host", cfg.GetHost()),
		zap.Int("port", cfg.GetPort()),
	)

	// Create HTTP client
	httpClient := httpClient.NewHTTPClient(
		time.Duration(cfg.GetApiTimeoutSecs())*time.Second,
		cfg.GetProxy(),
		cfg.ShouldForceIPv6(),
	)

	// Create Real-Debrid client
	rdClient := realdebrid.NewRealDebrid(cfg.GetToken(), httpClient)

	// Start premium status monitor
	go premium.MonitorPremiumStatus(ctx, rdClient, logger)

	// Create torrent manager
	torrentManager := torrent.NewTorrentManager(cfg, rdClient, logger)

	// Start background jobs
	torrentManager.StartRefreshJob(ctx)
	if cfg.EnableRepair() {
		torrentManager.StartRepairJob(ctx)
	}

	// Create handlers
	handlers := handlers.NewHandlers(cfg, torrentManager, logger)

	// Setup HTTP server
	mux := http.NewServeMux()
	handlers.AttachHandlers(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.GetHost(), cfg.GetPort()),
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server starting", zap.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	// Wait for context cancellation
	<-ctx.Done()

	logger.Info("Shutting down server...")
	
	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server shutdown failed", zap.Error(err))
		return err
	}

	logger.Info("Server stopped")
	return nil
}