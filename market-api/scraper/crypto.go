package scraper

import (
	"encoding/json"
	"log"

	"market-api/models"
)

const cryptoURL = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=20&page=1&sparkline=true&price_change_percentage=1h%2C24h%2C7d"

type coingeckoCrypto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	CurrentPrice   float64   `json:"current_price"`
	PriceChange1h  float64   `json:"price_change_percentage_1h_in_currency"`
	PriceChange24h float64   `json:"price_change_percentage_24h_in_currency"`
	PriceChange7d  float64   `json:"price_change_percentage_7d_in_currency"`
	MarketCap      float64   `json:"market_cap"`
	TotalVolume    float64   `json:"total_volume"`
	Image          string    `json:"image"`
	Sparkline7d    sparkline `json:"sparkline_in_7d"`
}

type sparkline struct {
	Price []float64 `json:"price"`
}

func FetchCrypto() ([]models.CryptoAsset, string, bool) {
	data, err := DoRequest(cryptoURL, nil)
	if err != nil {
		log.Printf("[CRYPTO] Fetch error: %v", err)
		return fallbackCrypto(), "fallback", true
	}

	var cryptos []coingeckoCrypto
	if err := json.Unmarshal(data, &cryptos); err != nil {
		log.Printf("[CRYPTO] JSON parse error: %v", err)
		return fallbackCrypto(), "fallback", true
	}

	if len(cryptos) == 0 {
		return fallbackCrypto(), "fallback", true
	}

	var result []models.CryptoAsset
	for _, c := range cryptos {
		result = append(result, models.CryptoAsset{
			ID:            c.ID,
			Name:          c.Name,
			Symbol:        c.Symbol,
			Price:         c.CurrentPrice,
			Change1h:      c.PriceChange1h,
			Change24h:     c.PriceChange24h,
			Change7d:      c.PriceChange7d,
			MarketCap:     c.MarketCap,
			Volume24h:     c.TotalVolume,
			ImageURL:      c.Image,
			SparklineData: c.Sparkline7d.Price,
		})
	}

	return result, "coingecko", false
}

func fallbackCrypto() []models.CryptoAsset {
	return []models.CryptoAsset{
		{ID: "bitcoin", Name: "Bitcoin", Symbol: "btc", Price: 84200, Change1h: 0.3, Change24h: 2.41, Change7d: 5.2, MarketCap: 1.65e12, Volume24h: 38e9, ImageURL: "https://assets.coingecko.com/coins/images/1/large/bitcoin.png"},
		{ID: "ethereum", Name: "Ethereum", Symbol: "eth", Price: 3200, Change1h: -0.1, Change24h: 1.82, Change7d: 3.1, MarketCap: 385e9, Volume24h: 18e9, ImageURL: "https://assets.coingecko.com/coins/images/279/large/ethereum.png"},
		{ID: "tether", Name: "Tether", Symbol: "usdt", Price: 1.0, Change1h: 0.01, Change24h: 0.02, Change7d: 0.01, MarketCap: 110e9, Volume24h: 85e9, ImageURL: "https://assets.coingecko.com/coins/images/325/large/Tether.png"},
		{ID: "binancecoin", Name: "BNB", Symbol: "bnb", Price: 580, Change1h: 0.2, Change24h: 1.1, Change7d: 2.3, MarketCap: 84e9, Volume24h: 2.1e9, ImageURL: "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png"},
		{ID: "solana", Name: "Solana", Symbol: "sol", Price: 175, Change1h: 0.5, Change24h: 3.2, Change7d: 8.1, MarketCap: 82e9, Volume24h: 4.5e9, ImageURL: "https://assets.coingecko.com/coins/images/4128/large/solana.png"},
		{ID: "ripple", Name: "XRP", Symbol: "xrp", Price: 0.62, Change1h: -0.2, Change24h: -1.3, Change7d: 2.1, MarketCap: 34e9, Volume24h: 1.8e9, ImageURL: "https://assets.coingecko.com/coins/images/44/large/xrp-symbol-white-128.png"},
		{ID: "usd-coin", Name: "USDC", Symbol: "usdc", Price: 1.0, Change1h: 0.0, Change24h: 0.01, Change7d: 0.02, MarketCap: 32e9, Volume24h: 7e9, ImageURL: "https://assets.coingecko.com/coins/images/6319/large/usdc.png"},
		{ID: "cardano", Name: "Cardano", Symbol: "ada", Price: 0.48, Change1h: 0.1, Change24h: 0.8, Change7d: 1.5, MarketCap: 17e9, Volume24h: 450e6, ImageURL: "https://assets.coingecko.com/coins/images/975/large/cardano.png"},
		{ID: "avalanche-2", Name: "Avalanche", Symbol: "avax", Price: 38, Change1h: 0.3, Change24h: 2.1, Change7d: 4.5, MarketCap: 16e9, Volume24h: 620e6, ImageURL: "https://assets.coingecko.com/coins/images/12559/large/Avalanche_Circle_RedWhite_Trans.png"},
		{ID: "dogecoin", Name: "Dogecoin", Symbol: "doge", Price: 0.165, Change1h: 0.4, Change24h: 1.9, Change7d: 3.8, MarketCap: 24e9, Volume24h: 1.2e9, ImageURL: "https://assets.coingecko.com/coins/images/5/large/dogecoin.png"},
		{ID: "polkadot", Name: "Polkadot", Symbol: "dot", Price: 7.8, Change1h: -0.1, Change24h: 0.5, Change7d: 1.2, MarketCap: 11e9, Volume24h: 280e6, ImageURL: "https://assets.coingecko.com/coins/images/12171/large/polkadot.png"},
		{ID: "chainlink", Name: "Chainlink", Symbol: "link", Price: 14.5, Change1h: 0.2, Change24h: 1.4, Change7d: 3.2, MarketCap: 9e9, Volume24h: 420e6, ImageURL: "https://assets.coingecko.com/coins/images/877/large/chainlink-new-logo.png"},
		{ID: "tron", Name: "TRON", Symbol: "trx", Price: 0.12, Change1h: 0.1, Change24h: 0.6, Change7d: 1.8, MarketCap: 10e9, Volume24h: 380e6, ImageURL: "https://assets.coingecko.com/coins/images/1094/large/tron-logo.png"},
		{ID: "shiba-inu", Name: "Shiba Inu", Symbol: "shib", Price: 0.0000245, Change1h: 0.5, Change24h: 2.8, Change7d: 5.1, MarketCap: 14e9, Volume24h: 680e6, ImageURL: "https://assets.coingecko.com/coins/images/11939/large/shiba.png"},
		{ID: "litecoin", Name: "Litecoin", Symbol: "ltc", Price: 88, Change1h: 0.1, Change24h: 0.9, Change7d: 2.1, MarketCap: 6.5e9, Volume24h: 320e6, ImageURL: "https://assets.coingecko.com/coins/images/2/large/litecoin.png"},
		{ID: "uniswap", Name: "Uniswap", Symbol: "uni", Price: 9.8, Change1h: 0.3, Change24h: 1.6, Change7d: 3.4, MarketCap: 5.9e9, Volume24h: 180e6, ImageURL: "https://assets.coingecko.com/coins/images/12504/large/uniswap-uni.png"},
		{ID: "stellar", Name: "Stellar", Symbol: "xlm", Price: 0.115, Change1h: 0.0, Change24h: 0.4, Change7d: 1.1, MarketCap: 3.2e9, Volume24h: 95e6, ImageURL: "https://assets.coingecko.com/coins/images/100/large/Stellar_symbol_black_RGB.png"},
		{ID: "monero", Name: "Monero", Symbol: "xmr", Price: 168, Change1h: -0.2, Change24h: -0.8, Change7d: 0.5, MarketCap: 3.1e9, Volume24h: 72e6, ImageURL: "https://assets.coingecko.com/coins/images/69/large/monero_logo.png"},
		{ID: "ethereum-classic", Name: "Ethereum Classic", Symbol: "etc", Price: 28.5, Change1h: 0.1, Change24h: 0.7, Change7d: 1.9, MarketCap: 4.1e9, Volume24h: 145e6, ImageURL: "https://assets.coingecko.com/coins/images/453/large/ethereum-classic-logo.png"},
		{ID: "cosmos", Name: "Cosmos", Symbol: "atom", Price: 8.9, Change1h: 0.2, Change24h: 1.1, Change7d: 2.5, MarketCap: 3.5e9, Volume24h: 165e6, ImageURL: "https://assets.coingecko.com/coins/images/1481/large/cosmos_hub.png"},
	}
}
