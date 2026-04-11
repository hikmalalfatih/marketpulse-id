package scraper

import (
	"encoding/json"
	"log"
	"strings"

	"market-api/config"
	"market-api/models"
)

type idxYahooChart struct {
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

func FetchIDStocks(cfg *config.Config) ([]models.StockAsset, string, bool) {
	var allStocks []models.StockAsset

	for _, ticker := range cfg.IDXTickers {
		url := "https://query1.finance.yahoo.com/v8/finance/chart/" + ticker + "?interval=1d&range=1d"
		data, err := DoRequest(url, map[string]string{
			"Referer": "https://finance.yahoo.com/",
		})
		if err != nil {
			log.Printf("[IDX-STOCK %s] Fetch error: %v", ticker, err)
			continue
		}

		var ch idxYahooChart
		if err := json.Unmarshal(data, &ch); err != nil {
			log.Printf("[IDX-STOCK %s] JSON parse error: %v", ticker, err)
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
			// Strip .JK suffix for display
			sym := strings.TrimSuffix(meta.Symbol, ".JK")
			stock := models.StockAsset{
				Symbol:        sym,
				Name:          name,
				Price:         meta.RegularMarketPrice,
				PreviousClose: prev,
				Change:        meta.RegularMarketPrice - prev,
				ChangePercent: safePercent(meta.RegularMarketPrice, prev),
				Currency:      "IDR",
				Market:        "IDX",
			}
			allStocks = append(allStocks, stock)
		}
	}

	if len(allStocks) == 0 {
		return fallbackIDStocks(), "fallback", true
	}

	return allStocks, "yahoo", false
}

func safePercent(current, prev float64) float64 {
	if prev == 0 {
		return 0
	}
	return ((current - prev) / prev) * 100
}

func fallbackIDStocks() []models.StockAsset {
	return []models.StockAsset{
		{Symbol: "BBCA", Name: "Bank Central Asia Tbk", Price: 10250, PreviousClose: 10200, Change: 50, ChangePercent: 0.49, Currency: "IDR", Market: "IDX"},
		{Symbol: "BBRI", Name: "Bank Rakyat Indonesia Tbk", Price: 5475, PreviousClose: 5450, Change: 25, ChangePercent: 0.46, Currency: "IDR", Market: "IDX"},
		{Symbol: "BMRI", Name: "Bank Mandiri Tbk", Price: 7200, PreviousClose: 7150, Change: 50, ChangePercent: 0.70, Currency: "IDR", Market: "IDX"},
		{Symbol: "TLKM", Name: "Telkom Indonesia Tbk", Price: 3890, PreviousClose: 3920, Change: -30, ChangePercent: -0.77, Currency: "IDR", Market: "IDX"},
		{Symbol: "ASII", Name: "Astra International Tbk", Price: 5200, PreviousClose: 5175, Change: 25, ChangePercent: 0.48, Currency: "IDR", Market: "IDX"},
		{Symbol: "UNVR", Name: "Unilever Indonesia Tbk", Price: 2840, PreviousClose: 2860, Change: -20, ChangePercent: -0.70, Currency: "IDR", Market: "IDX"},
		{Symbol: "GOTO", Name: "GoTo Gojek Tokopedia Tbk", Price: 68, PreviousClose: 66, Change: 2, ChangePercent: 3.03, Currency: "IDR", Market: "IDX"},
		{Symbol: "BYAN", Name: "Bayan Resources Tbk", Price: 18500, PreviousClose: 18250, Change: 250, ChangePercent: 1.37, Currency: "IDR", Market: "IDX"},
		{Symbol: "ADRO", Name: "Adaro Energy Indonesia Tbk", Price: 2640, PreviousClose: 2620, Change: 20, ChangePercent: 0.76, Currency: "IDR", Market: "IDX"},
		{Symbol: "EXCL", Name: "XL Axiata Tbk", Price: 2180, PreviousClose: 2160, Change: 20, ChangePercent: 0.93, Currency: "IDR", Market: "IDX"},
		{Symbol: "ICBP", Name: "Indofood CBP Sukses Makmur Tbk", Price: 11050, PreviousClose: 11000, Change: 50, ChangePercent: 0.45, Currency: "IDR", Market: "IDX"},
		{Symbol: "KLBF", Name: "Kalbe Farma Tbk", Price: 1620, PreviousClose: 1610, Change: 10, ChangePercent: 0.62, Currency: "IDR", Market: "IDX"},
		{Symbol: "SMGR", Name: "Semen Indonesia Tbk", Price: 5750, PreviousClose: 5800, Change: -50, ChangePercent: -0.86, Currency: "IDR", Market: "IDX"},
		{Symbol: "PTBA", Name: "Bukit Asam Tbk", Price: 3120, PreviousClose: 3100, Change: 20, ChangePercent: 0.65, Currency: "IDR", Market: "IDX"},
		{Symbol: "ANTM", Name: "Aneka Tambang Tbk", Price: 1680, PreviousClose: 1660, Change: 20, ChangePercent: 1.20, Currency: "IDR", Market: "IDX"},
	}
}
