package scraper

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var defaultHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Accept":          "application/json, text/html, */*",
	"Accept-Language": "en-US,en;q=0.9,id;q=0.8",
	"Accept-Encoding": "gzip, deflate, br",
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func DoRequest(url string, extraHeaders map[string]string) ([]byte, error) {
	client := NewHTTPClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add default headers
	for k, v := range defaultHeaders {
		req.Header.Set(k, v)
	}

	// Override with extra
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

