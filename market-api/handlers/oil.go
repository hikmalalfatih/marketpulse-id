package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/config"
	"market-api/scraper"
)

func OilHandler(c *cache.Cache, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_oil"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		oils, source, fallback := scraper.FetchOil(cfg)
		c.Set(key, oils, 5*time.Minute)
		writeResponse(w, true, oils, !fallback, source, time.Now())
	}
}


