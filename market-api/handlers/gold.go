package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/scraper"
)

func GoldHandler(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_gold"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		gold, source, fallback := scraper.FetchGold()
		c.Set(key, gold, 10*time.Minute)
		writeResponse(w, true, gold, !fallback, source, time.Now())
	}
}


