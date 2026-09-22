package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"market-api/models"
)

const (
	goldURL   = "https://query1.finance.yahoo.com/v8/finance/chart/GC=F?range=1d&interval=1d"
	silverURL = "https://query1.finance.yahoo.com/v8/finance/chart/SI=F?range=1d&interval=1d"
)

type yahooChartResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

const (
	ozToGram = 31.1035
	usdidr   = 15850.0
)

func FetchGold() (models.GoldData, string, bool) {
	goldOz, err := fetchYahooMarketPrice(goldURL)
	if err != nil {
		log.Printf("[GOLD] Fetch error: %v", err)
		return fallbackGold(), "fallback", true
	}

	silverOz, err := fetchYahooMarketPrice(silverURL)
	if err != nil {
		log.Printf("[GOLD] Silver fetch error: %v", err)
		return fallbackGold(), "fallback", true
	}

	gold := models.GoldData{
		PriceUSDPerOz:    goldOz,
		PriceUSDPerGram:  goldOz / ozToGram,
		PriceIDRPerOz:    goldOz * usdidr,
		PriceIDRPerGram:  (goldOz / ozToGram) * usdidr,
		SilverUSDPerOz:   silverOz,
		SilverIDRPerGram: (silverOz / ozToGram) * usdidr,
		ChangePercent:    0.0, // Single snapshot
		UpdatedAt:        time.Now().Format(time.RFC3339),
	}

	return gold, "yahoo-finance", false
}

func fetchYahooMarketPrice(url string) (float64, error) {
	data, err := DoRequest(url, nil)
	if err != nil {
		return 0, err
	}

	var response yahooChartResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return 0, err
	}
	if len(response.Chart.Result) == 0 || response.Chart.Result[0].Meta.RegularMarketPrice <= 0 {
		return 0, fmt.Errorf("market price missing from response")
	}

	return response.Chart.Result[0].Meta.RegularMarketPrice, nil
}

func fallbackGold() models.GoldData {
	return models.GoldData{
		PriceUSDPerOz:    2650.0,
		PriceUSDPerGram:  85.2,
		PriceIDRPerOz:    2650 * usdidr,
		PriceIDRPerGram:  85.2 * usdidr,
		SilverUSDPerOz:   32.0,
		SilverIDRPerGram: (32.0 / ozToGram) * usdidr,
		UpdatedAt:        time.Now().Format(time.RFC3339),
	}
}
