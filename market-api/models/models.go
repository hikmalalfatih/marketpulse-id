package models

type APIResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Cached    bool        `json:"cached"`
	Stale     bool        `json:"stale,omitempty"`
	Source    string      `json:"source"`
	UpdatedAt string      `json:"updated_at"`
	Error     string      `json:"error,omitempty"`
}

type CryptoAsset struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Symbol        string    `json:"symbol"`
	Price         float64   `json:"price"`
	Change1h      float64   `json:"change_1h"`
	Change24h     float64   `json:"change_24h"`
	Change7d      float64   `json:"change_7d"`
	MarketCap     float64   `json:"market_cap"`
	Volume24h     float64   `json:"volume_24h"`
	ImageURL      string    `json:"image_url"`
	SparklineData []float64 `json:"sparkline_7d"`
}

type StockAsset struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	PreviousClose float64 `json:"previous_close"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	Currency      string  `json:"currency"`
	Market        string  `json:"market"`
}

type HistoryPoint struct {
	Timestamp int64   `json:"timestamp"`
	Date      string  `json:"date"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Volume    float64 `json:"volume"`
}

type GoldData struct {
	PriceUSDPerOz    float64 `json:"price_usd_per_oz"`
	PriceUSDPerGram  float64 `json:"price_usd_per_gram"`
	PriceIDRPerOz    float64 `json:"price_idr_per_oz"`
	PriceIDRPerGram  float64 `json:"price_idr_per_gram"`
	SilverUSDPerOz   float64 `json:"silver_usd_per_oz"`
	SilverIDRPerGram float64 `json:"silver_idr_per_gram"`
	ChangePercent    float64 `json:"change_percent"`
	UpdatedAt        string  `json:"updated_at"`
}

type OilData struct {
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	PriceUSD      float64 `json:"price_usd"`
	PriceIDR      float64 `json:"price_idr"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	Unit          string  `json:"unit"`
}

type ForexRate struct {
	From      string             `json:"from"`
	Rates     map[string]float64 `json:"rates"`
	UpdatedAt string             `json:"updated_at"`
}

type ReksadanaItem struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	NABPerUnit float64 `json:"nab_per_unit"`
	Return1M   float64 `json:"return_1m"`
	ReturnYTD  float64 `json:"return_ytd"`
	Return1Y   float64 `json:"return_1y"`
	AUM        float64 `json:"aum,omitempty"`
	Source     string  `json:"data_source"`
}

type ObligasiItem struct {
	Kode       string  `json:"kode"`
	Nama       string  `json:"nama"`
	Jenis      string  `json:"jenis"`
	Kupon      float64 `json:"kupon"`
	JatuhTempo string  `json:"jatuh_tempo"`
	HargaPasar float64 `json:"harga_pasar"`
	Yield      float64 `json:"yield"`
	Tenor      string  `json:"tenor"`
	Source     string  `json:"data_source"`
}

type SummaryData struct {
	CryptoSummary []CryptoAsset   `json:"crypto"`
	IDXStocks     []StockAsset    `json:"stocks_id"`
	USStocks      []StockAsset    `json:"stocks_us"`
	Gold          GoldData        `json:"gold"`
	Oil           []OilData       `json:"oil"`
	Forex         ForexRate       `json:"forex"`
	Reksadana     []ReksadanaItem `json:"reksadana"`
	Obligasi      []ObligasiItem  `json:"obligasi"`
	UpdatedAt     string          `json:"updated_at"`
}
