package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	Port        string
	HTTPTimeout int
	IDXTickers  []string
	USTickers   []string
	OilSymbols  []string
}

var defaultConfig = &Config{
	Port:        ":8080",
	HTTPTimeout: 10,
	IDXTickers: []string{
		"BBCA.JK", "BBRI.JK", "BMRI.JK", "TLKM.JK", "ASII.JK",
		"UNVR.JK", "GOTO.JK", "BYAN.JK", "ADRO.JK", "EXCL.JK",
		"ICBP.JK", "KLBF.JK", "SMGR.JK", "PTBA.JK", "ANTM.JK",
	},
	USTickers: []string{
		"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA",
		"META", "NVDA", "JPM", "V", "BRK-B",
	},
	OilSymbols: []string{"BZ=F", "CL=F", "NG=F"}, // Brent, WTI, Natural Gas
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultConfig.Port
	} else if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	cfg := &Config{
		Port:        port,
		HTTPTimeout: defaultConfig.HTTPTimeout,
		IDXTickers:  defaultConfig.IDXTickers,
		USTickers:   defaultConfig.USTickers,
		OilSymbols:  defaultConfig.OilSymbols,
	}

	log.Printf("Config loaded: Port=%s, IDX=%d tickers, US=%d tickers",
		cfg.Port, len(cfg.IDXTickers), len(cfg.USTickers))

	return cfg
}
