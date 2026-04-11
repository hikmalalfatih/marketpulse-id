package scraper

import (
	"encoding/json"
	"log"

	"market-api/config"
	"market-api/models"
)

	type OilYahooChart struct {
	Chart struct {
		Result []struct {
			Meta struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				ChartPreviousClose float64 `json:"chartPreviousClose"`
				Symbol             string  `json:"symbol"`
				Currency           string  `json:"currency"`
				RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

func FetchOil(cfg *config.Config) ([]models.OilData, string, bool) {
	var oils []models.OilData

	for _, symbol := range cfg.OilSymbols {
		url := "https://query1.finance.yahoo.com/v8/finance/chart/" + symbol + "?interval=1d&range=1d"
		data, err := DoRequest(url, map[string]string{
			"Referer": "https://finance.yahoo.com/",
		})
		if err != nil {
			log.Printf("[OIL %s] Fetch error: %v", symbol, err)
			continue
		}

	var ch OilYahooChart
		if err := json.Unmarshal(data, &ch); err != nil {
			log.Printf("[OIL %s] JSON parse error: %v", symbol, err)
			continue
		}

		if len(ch.Chart.Result) > 0 {
			meta := ch.Chart.Result[0].Meta
			name := symbolName(symbol)
			priceIDR := meta.RegularMarketPrice * 15850 // approx USD/IDR
			oil := models.OilData{
				Name:           name,
				Symbol:         symbol,
				PriceUSD:       meta.RegularMarketPrice,
				PriceIDR:       priceIDR,
				Change:         meta.RegularMarketPrice - meta.ChartPreviousClose,
				ChangePercent:  meta.RegularMarketChangePercent,
				Unit:           "per barrel",
			}
			oils = append(oils, oil)
		}
	}

	if len(oils) == 0 {
		return fallbackOil(), "fallback", true
	}

	return oils, "yahoo", false
}

func symbolName(symbol string) string {
	switch symbol {
	case "BZ=F":
		return "Brent Crude"
	case "CL=F":
		return "WTI Crude"
	case "NG=F":
		return "Natural Gas"
	default:
		return symbol
	}
}

func fallbackOil() []models.OilData {
	return []models.OilData{
		{
			Name:          "Brent Crude",
			Symbol:        "BZ=F",
			PriceUSD:      85.5,
			PriceIDR:      85.5 * 15850,
			ChangePercent: 1.2,
			Unit:          "per barrel",
		},
		{
			Name:          "WTI Crude",
			Symbol:        "CL=F",
			PriceUSD:      81.2,
			PriceIDR:      81.2 * 15850,
			ChangePercent: 0.8,
			Unit:          "per barrel",
		},
	}
}

