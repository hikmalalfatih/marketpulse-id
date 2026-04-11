package handlers

import (
	"net/http"
	"sync"
	"time"

	"market-api/cache"
	"market-api/config"
	"market-api/models"
	"market-api/scraper"
)

func SummaryHandler(c *cache.Cache, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := "api_summary"
		if data, ok := c.Get(key); ok {
			writeResponse(w, true, data, true, "", time.Now())
			return
		}

		sumData := fetchSummary(cfg)
		c.Set(key, sumData, 5*time.Minute)
		writeResponse(w, true, sumData, false, "combined", time.Now())
	}
}

func fetchSummary(cfg *config.Config) *models.SummaryData {
	var sum models.SummaryData
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		assets, _, _ := scraper.FetchCrypto()
		mu.Lock()
		sum.CryptoSummary = assets
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stocks, _, _ := scraper.FetchIDStocks(cfg)
		mu.Lock()
		sum.IDXStocks = stocks
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stocks, _, _ := scraper.FetchUSStocks(cfg)
		mu.Lock()
		sum.USStocks = stocks
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		g, _, _ := scraper.FetchGold()
		mu.Lock()
		sum.Gold = g
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		oils, _, _ := scraper.FetchOil(cfg)
		mu.Lock()
		sum.Oil = oils
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f, _, _ := scraper.FetchForex()
		mu.Lock()
		sum.Forex = f
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r, _, _ := scraper.FetchReksadana()
		mu.Lock()
		sum.Reksadana = r
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		o, _, _ := scraper.FetchObligasi()
		mu.Lock()
		sum.Obligasi = o
		mu.Unlock()
	}()

	wg.Wait()
	sum.UpdatedAt = time.Now().Format(time.RFC3339)
	return &sum
}
