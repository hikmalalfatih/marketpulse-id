package handlers

import (
	"net/http"
	"strings"
	"time"

	"market-api/cache"
	"market-api/config"
	"market-api/models"
	"market-api/scraper"
)

func StocksHandler(c *cache.Cache, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		var key string
		var data []models.StockAsset
		var source string
		fallback := false

		if strings.Contains(path, "/stocks/id") {
			key = "api_stocks_id"
			if d, ok := c.Get(key); ok {
				writeResponse(w, true, d, true, "", time.Now())
				return
			}
			data, source, fallback = scraper.FetchIDStocks(cfg)
			c.Set(key, data, 5*time.Minute)
		} else if strings.Contains(path, "/stocks/us") {
			key = "api_stocks_us"
			if d, ok := c.Get(key); ok {
				writeResponse(w, true, d, true, "", time.Now())
				return
			}
			data, source, fallback = scraper.FetchUSStocks(cfg)
			c.Set(key, data, 5*time.Minute)
		} else {
			writeResponse(w, false, nil, false, "invalid_path", time.Now())
			return
		}

		writeResponse(w, true, data, !fallback, source, time.Now())
	}
}


