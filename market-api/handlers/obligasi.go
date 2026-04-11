package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/scraper"
)

func ObligasiHandler(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_obligasi"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		items, source, fallback := scraper.FetchObligasi()
		c.Set(key, items, 60*time.Minute)
		writeResponse(w, true, items, !fallback, source, time.Now())
	}
}


