package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/scraper"
)

func ForexHandler(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_forex"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		rates, source, fallback := scraper.FetchForex()
		c.Set(key, rates, 30*time.Minute)
		writeResponse(w, true, rates, !fallback, source, time.Now())
	}
}
