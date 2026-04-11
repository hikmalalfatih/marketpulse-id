package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"market-api/cache"
	"market-api/config"
	"market-api/handlers"
	"market-api/middleware"
)

var globalCache *cache.Cache
var globalConfig *config.Config

func main() {
	globalCache = cache.NewCache()
	globalConfig = config.LoadConfig()

	// HTTP client timeout
	http.DefaultClient.Timeout = 10 * time.Second

	log.Printf("🚀 MarketPulse ID API v1.0 starting on %s", globalConfig.Port)

	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/health", corsMiddleware(handlers.HealthHandler))

	// Market data endpoints
	mux.HandleFunc("/api/crypto", corsMiddleware(handlers.CryptoHandler(globalCache)))
	mux.HandleFunc("/api/stocks/id", corsMiddleware(handlers.StocksHandler(globalCache, globalConfig)))
	mux.HandleFunc("/api/stocks/us", corsMiddleware(handlers.StocksHandler(globalCache, globalConfig)))
	mux.HandleFunc("/api/gold", corsMiddleware(handlers.GoldHandler(globalCache)))
	mux.HandleFunc("/api/oil", corsMiddleware(handlers.OilHandler(globalCache, globalConfig)))
	mux.HandleFunc("/api/forex", corsMiddleware(handlers.ForexHandler(globalCache)))
	mux.HandleFunc("/api/reksadana", corsMiddleware(handlers.ReksadanaHandler(globalCache)))
	mux.HandleFunc("/api/obligasi", corsMiddleware(handlers.ObligasiHandler(globalCache)))
	mux.HandleFunc("/api/history/", corsMiddleware(handlers.HistoryHandler(globalCache, globalConfig)))
	mux.HandleFunc("/api/summary", corsMiddleware(handlers.SummaryHandler(globalCache, globalConfig)))

	srv := &http.Server{
		Addr:         globalConfig.Port,
		Handler:      mux,
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

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return middleware.CORSMiddleware(next)
}

