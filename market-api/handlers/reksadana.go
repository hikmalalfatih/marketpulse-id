package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/scraper"
)

func ReksadanaHandler(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_reksadana"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		items, source, fallback := scraper.FetchReksadana()
		c.Set(key, items, 60*time.Minute)
		writeResponse(w, true, items, !fallback, source, time.Now())
	}
}


