package scraper

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"market-api/models"
)

const goldURL = "https://data-asg.goldprice.org/dbXRates/USD"

type goldResponse struct {
	Items []struct {
		XauPrice string `json:"xauPrice"`
		XagPrice string `json:"xagPrice"`
	} `json:"items"`
}

const (
	ozToGram = 31.1035
	usdidr   = 15850.0
)

func FetchGold() (models.GoldData, string, bool) {
	data, err := DoRequest(goldURL, map[string]string{
		"Origin":  "https://goldprice.org",
		"Referer": "https://goldprice.org/",
	})
	if err != nil {
		log.Printf("[GOLD] Fetch error: %v", err)
		return fallbackGold(), "fallback", true
	}

	var resp goldResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		log.Printf("[GOLD] JSON parse error: %v", err)
		return fallbackGold(), "fallback", true
	}

	if len(resp.Items) == 0 {
		return fallbackGold(), "fallback", true
	}

	item := resp.Items[0]
	goldOz, _ := strconv.ParseFloat(item.XauPrice, 64)
	silverOz, _ := strconv.ParseFloat(item.XagPrice, 64)

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

	return gold, "goldprice.org", false
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

