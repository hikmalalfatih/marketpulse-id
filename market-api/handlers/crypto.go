package handlers

import (
	"net/http"
	"time"

	"market-api/cache"
	"market-api/scraper"
)

func CryptoHandler(c *cache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_crypto"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		assets, source, fallback := scraper.FetchCrypto()
		c.Set(key, assets, 3*time.Minute)
		writeResponse(w, true, assets, !fallback, source, time.Now())
	}
}
