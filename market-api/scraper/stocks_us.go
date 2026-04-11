package scraper

import (
	"encoding/json"
	"log"

	"market-api/config"
	"market-api/models"
)

type usYahooChart struct {
	Chart struct {
		Result []struct {
			Meta struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				PreviousClose      float64 `json:"previousClose"`
				ChartPreviousClose float64 `json:"chartPreviousClose"`
				Symbol             string  `json:"symbol"`
				LongName           string  `json:"longName"`
				ShortName          string  `json:"shortName"`
				Currency           string  `json:"currency"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

func FetchUSStocks(cfg *config.Config) ([]models.StockAsset, string, bool) {
	var allStocks []models.StockAsset

	for _, ticker := range cfg.USTickers {
		url := "https://query1.finance.yahoo.com/v8/finance/chart/" + ticker + "?interval=1d&range=1d"
		data, err := DoRequest(url, map[string]string{
			"Referer": "https://finance.yahoo.com/",
		})
		if err != nil {
			log.Printf("[US-STOCK %s] Fetch error: %v", ticker, err)
			continue
		}

		var ch usYahooChart
		if err := json.Unmarshal(data, &ch); err != nil {
			log.Printf("[US-STOCK %s] JSON parse error: %v", ticker, err)
			continue
		}

		if len(ch.Chart.Result) > 0 && ch.Chart.Result[0].Meta.RegularMarketPrice > 0 {
			meta := ch.Chart.Result[0].Meta
			prev := meta.PreviousClose
			if prev == 0 {
				prev = meta.ChartPreviousClose
			}
			name := meta.LongName
			if name == "" {
				name = meta.ShortName
			}
			cur := meta.Currency
			if cur == "" {
				cur = "USD"
			}
			stock := models.StockAsset{
				Symbol:        meta.Symbol,
				Name:          name,
				Price:         meta.RegularMarketPrice,
				PreviousClose: prev,
				Change:        meta.RegularMarketPrice - prev,
				ChangePercent: safePercent(meta.RegularMarketPrice, prev),
				Currency:      cur,
				Market:        "US",
			}
			allStocks = append(allStocks, stock)
		}
	}

	if len(allStocks) == 0 {
		return fallbackUSStocks(), "fallback", true
	}

	return allStocks, "yahoo", false
}

func fallbackUSStocks() []models.StockAsset {
	return []models.StockAsset{
		{Symbol: "AAPL", Name: "Apple Inc.", Price: 228.50, PreviousClose: 226.00, Change: 2.50, ChangePercent: 1.11, Currency: "USD", Market: "US"},
		{Symbol: "MSFT", Name: "Microsoft Corporation", Price: 415.20, PreviousClose: 412.00, Change: 3.20, ChangePercent: 0.78, Currency: "USD", Market: "US"},
		{Symbol: "GOOGL", Name: "Alphabet Inc.", Price: 175.80, PreviousClose: 174.50, Change: 1.30, ChangePercent: 0.74, Currency: "USD", Market: "US"},
		{Symbol: "AMZN", Name: "Amazon.com Inc.", Price: 198.40, PreviousClose: 196.00, Change: 2.40, ChangePercent: 1.22, Currency: "USD", Market: "US"},
		{Symbol: "TSLA", Name: "Tesla Inc.", Price: 248.50, PreviousClose: 252.00, Change: -3.50, ChangePercent: -1.39, Currency: "USD", Market: "US"},
		{Symbol: "META", Name: "Meta Platforms Inc.", Price: 512.30, PreviousClose: 508.00, Change: 4.30, ChangePercent: 0.85, Currency: "USD", Market: "US"},
		{Symbol: "NVDA", Name: "NVIDIA Corporation", Price: 875.40, PreviousClose: 868.00, Change: 7.40, ChangePercent: 0.85, Currency: "USD", Market: "US"},
		{Symbol: "JPM", Name: "JPMorgan Chase & Co.", Price: 198.60, PreviousClose: 197.00, Change: 1.60, ChangePercent: 0.81, Currency: "USD", Market: "US"},
		{Symbol: "V", Name: "Visa Inc.", Price: 278.90, PreviousClose: 276.50, Change: 2.40, ChangePercent: 0.87, Currency: "USD", Market: "US"},
		{Symbol: "BRK-B", Name: "Berkshire Hathaway Inc.", Price: 412.50, PreviousClose: 410.00, Change: 2.50, ChangePercent: 0.61, Currency: "USD", Market: "US"},
	}
}
