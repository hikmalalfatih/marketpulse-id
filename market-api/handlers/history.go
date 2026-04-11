package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"market-api/cache"
	"market-api/config"
	"market-api/models"
	"market-api/scraper"
)

func HistoryHandler(c *cache.Cache, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/history/"), "/")
		if len(parts) < 2 {
			writeJSON(w, models.APIResponse{
				Success:   false,
				Error:     "invalid path: use /api/history/{type}/{symbol}",
				UpdatedAt: time.Now().Format(time.RFC3339),
			})
			return
		}

		typ := parts[0]
		symbol := parts[1]

		key := "history_" + typ + "_" + symbol
		if data, ok := c.Get(key); ok {
			writeJSON(w, models.APIResponse{
				Success:   true,
				Data:      data,
				Cached:    true,
				UpdatedAt: time.Now().Format(time.RFC3339),
			})
			return
		}

		history, source, _ := fetchHistory(typ, symbol, cfg)
		c.Set(key, history, 15*time.Minute)
		writeJSON(w, models.APIResponse{
			Success:   true,
			Data:      history,
			Cached:    false,
			Source:    source,
			UpdatedAt: time.Now().Format(time.RFC3339),
		})
	}
}

func fetchHistory(typ, symbol string, cfg *config.Config) ([]models.HistoryPoint, string, bool) {
	switch typ {
	case "crypto":
		return fetchCryptoHistory(symbol)
	case "stock_id":
		return fetchYahooHistory(symbol + ".JK")
	case "stock_us":
		return fetchYahooHistory(symbol)
	case "gold":
		return fetchYahooHistory("GC%3DF")
	case "oil":
		switch symbol {
		case "brent":
			return fetchYahooHistory("BZ%3DF")
		case "wti":
			return fetchYahooHistory("CL%3DF")
		case "natgas":
			return fetchYahooHistory("NG%3DF")
		}
		return fallbackHistory(), "invalid oil symbol", true
	default:
		return fallbackHistory(), "invalid type", true
	}
}

func fetchCryptoHistory(symbol string) ([]models.HistoryPoint, string, bool) {
	u := "https://api.coingecko.com/api/v3/coins/" + symbol + "/market_chart?vs_currency=usd&days=90&interval=daily"
	data, err := scraper.DoRequest(u, nil)
	if err != nil {
		return fallbackHistory(), "error", true
	}

	type cgResp struct {
		Prices [][]float64 `json:"prices"`
	}
	var r cgResp
	if err := json.Unmarshal(data, &r); err != nil {
		return fallbackHistory(), "parse", true
	}

	var h []models.HistoryPoint
	for _, p := range r.Prices {
		if len(p) == 2 {
			h = append(h, models.HistoryPoint{
				Timestamp: int64(p[0]) / 1000,
				Date:      time.UnixMilli(int64(p[0])).Format("2006-01-02"),
				Close:     p[1],
				Open:      p[1],
				High:      p[1],
				Low:       p[1],
			})
		}
	}
	if len(h) == 0 {
		return fallbackHistory(), "empty", true
	}
	return h, "coingecko", false
}

func fetchYahooHistory(symbol string) ([]models.HistoryPoint, string, bool) {
	u := "https://query1.finance.yahoo.com/v8/finance/chart/" + symbol + "?interval=1d&range=3mo"
	data, err := scraper.DoRequest(u, map[string]string{
		"Referer": "https://finance.yahoo.com/",
	})
	if err != nil {
		return fallbackHistory(), "error", true
	}

	// Yahoo Finance v8 response structure
	type quoteData struct {
		Open   []float64 `json:"open"`
		High   []float64 `json:"high"`
		Low    []float64 `json:"low"`
		Close  []float64 `json:"close"`
		Volume []float64 `json:"volume"`
	}
	type yahooResult struct {
		Timestamp  []int64 `json:"timestamp"`
		Indicators struct {
			Quote []quoteData `json:"quote"`
		} `json:"indicators"`
	}
	type yahooResp struct {
		Chart struct {
			Result []yahooResult `json:"result"`
			Error  interface{}   `json:"error"`
		} `json:"chart"`
	}

	var r yahooResp
	if err := json.Unmarshal(data, &r); err != nil {
		return fallbackHistory(), "parse", true
	}

	if len(r.Chart.Result) == 0 || len(r.Chart.Result[0].Timestamp) == 0 {
		return fallbackHistory(), "no data", true
	}

	res := r.Chart.Result[0]
	if len(res.Indicators.Quote) == 0 {
		return fallbackHistory(), "no quote", true
	}
	quote := res.Indicators.Quote[0]
	n := len(res.Timestamp)

	var h []models.HistoryPoint
	for i := 0; i < n; i++ {
		if i >= len(quote.Close) || quote.Close[i] == 0 {
			continue
		}
		open := safeGet(quote.Open, i)
		high := safeGet(quote.High, i)
		low := safeGet(quote.Low, i)
		vol := safeGet(quote.Volume, i)
		h = append(h, models.HistoryPoint{
			Timestamp: res.Timestamp[i],
			Date:      time.Unix(res.Timestamp[i], 0).Format("2006-01-02"),
			Open:      open,
			High:      high,
			Low:       low,
			Close:     quote.Close[i],
			Volume:    vol,
		})
	}
	if len(h) == 0 {
		return fallbackHistory(), "empty", true
	}
	return h, "yahoo", false
}

func safeGet(arr []float64, i int) float64 {
	if i < len(arr) {
		return arr[i]
	}
	return 0
}

func fallbackHistory() []models.HistoryPoint {
	base := time.Now().AddDate(0, -3, 0)
	var h []models.HistoryPoint
	price := 100.0
	for i := 0; i < 90; i++ {
		t := base.AddDate(0, 0, i)
		price += (float64(i%7) - 3) * 0.5
		h = append(h, models.HistoryPoint{
			Timestamp: t.Unix(),
			Date:      t.Format("2006-01-02"),
			Open:      price - 0.5,
			High:      price + 1,
			Low:       price - 1,
			Close:     price,
			Volume:    1000000,
		})
	}
	return h
}
