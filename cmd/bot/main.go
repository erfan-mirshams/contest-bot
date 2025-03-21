package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"contest-bot/internal/config"
	"contest-bot/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load config")
	}

	log := logger.SetupLogger(cfg.Log.Level)

	// Create HTTP server
	mux := http.NewServeMux()

	// Echo handler
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		log.Info("Received echo request", "method", r.Method, "path", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("Failed to read request body", "error", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Start server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Error("Server forced to shutdown", "error", err)
		}
	}()

	// Start the server
	log.Info("Server listening", "port", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

	log.Info("Server stopped")
}
