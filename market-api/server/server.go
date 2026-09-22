// Package server wires the HTTP routes shared by the local server and Vercel.
package server

import (
	"net/http"

	"market-api/cache"
	"market-api/config"
	"market-api/handlers"
	"market-api/middleware"
)

// NewHandler creates the complete API router. It does not open a port, so it
// can be used both by the local executable and a Vercel Serverless Function.
func NewHandler(cfg *config.Config) http.Handler {
	c := cache.NewCache()
	mux := http.NewServeMux()

	mux.HandleFunc("/health", cors(handlers.HealthHandler))
	mux.HandleFunc("/api/health", cors(handlers.HealthHandler))
	mux.HandleFunc("/api/crypto", cors(handlers.CryptoHandler(c)))
	mux.HandleFunc("/api/stocks/id", cors(handlers.StocksHandler(c, cfg)))
	mux.HandleFunc("/api/stocks/us", cors(handlers.StocksHandler(c, cfg)))
	mux.HandleFunc("/api/gold", cors(handlers.GoldHandler(c)))
	mux.HandleFunc("/api/oil", cors(handlers.OilHandler(c, cfg)))
	mux.HandleFunc("/api/forex", cors(handlers.ForexHandler(c)))
	mux.HandleFunc("/api/reksadana", cors(handlers.ReksadanaHandler(c)))
	mux.HandleFunc("/api/obligasi", cors(handlers.ObligasiHandler(c)))
	mux.HandleFunc("/api/history/", cors(handlers.HistoryHandler(c, cfg)))
	mux.HandleFunc("/api/summary", cors(handlers.SummaryHandler(c, cfg)))

	return mux
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return middleware.CORSMiddleware(next)
}
