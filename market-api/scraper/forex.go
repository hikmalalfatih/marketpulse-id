package scraper

import (
	"encoding/json"
	"log"
	"time"

	"market-api/models"
)

const forexURL = "https://open.er-api.com/v6/latest/USD"

type erResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchForex() (models.ForexRate, string, bool) {
	data, err := DoRequest(forexURL, nil)
	if err != nil {
		log.Printf("[FOREX] Fetch error: %v", err)
		return fallbackForex(), "fallback", true
	}

	var resp erResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		log.Printf("[FOREX] JSON parse error: %v", err)
		return fallbackForex(), "fallback", true
	}

	rates := models.ForexRate{
		From:      "USD",
		Rates:     resp.Rates,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	// Ensure IDR present
	if _, ok := rates.Rates["IDR"]; !ok {
		rates.Rates["IDR"] = 16000.0 // fallback
	}

	return rates, "er-api", false
}

func fallbackForex() models.ForexRate {
	return models.ForexRate{
		From: "USD",
		Rates: map[string]float64{
			"IDR":  15850.0,
			"EUR":  0.92,
			"GBP":  0.79,
			"SGD":  1.34,
			"JPY":  150.0,
			"CNY":  7.1,
			"AUD":  1.5,
			"SAR":  3.75,
			"MYR":  4.7,
		},
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}

