package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"market-api/config"
	"market-api/server"
)

func main() {
	globalConfig := config.LoadConfig()

	// HTTP client timeout
	http.DefaultClient.Timeout = 10 * time.Second

	log.Printf("🚀 MarketPulse ID API v1.0 starting on %s", globalConfig.Port)

	srv := &http.Server{
		Addr:         globalConfig.Port,
		Handler:      server.NewHandler(globalConfig),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctxShut, cancelShut := context.WithTimeout(ctx, 30*time.Second)
	defer cancelShut()
	if err := srv.Shutdown(ctxShut); err != nil {
		log.Printf("Server forced shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
