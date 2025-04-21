package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           Recorderis API
// @version         0.1.0
// @description     API for the Recorderis memories application
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  your-email@domain.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:4000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.
func main() {
	// Configure structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Composer
	userAdapter, authAdapter, memoryAdapter, tokenMgr, locationAdapter, err := Compose()
	if err != nil {
		slog.Error("Failed to compose dependencies", "error", err)
		os.Exit(1)
	}

	/* Create router */
	router := CreateRouter(userAdapter, authAdapter, memoryAdapter, tokenMgr, locationAdapter)

	// Configure server
	server := &http.Server{
		Addr:         ":4000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		slog.Info("Server starting", "port", "4000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	<-quit
	slog.Info("Server shutting down...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server stopped gracefully")
}
