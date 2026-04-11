package scraper

import (
	"log"
	"strconv"
	"strings"

	"market-api/models"
)

const reksadanaURL = "https://pusatdata.kontan.co.id/reksadana/"

func FetchReksadana() ([]models.ReksadanaItem, string, bool) {
	data, err := DoRequest(reksadanaURL, nil)
	if err != nil {
		log.Printf("[REKSADANA] Fetch error: %v", err)
		return fallbackReksadana(), "fallback", true
	}

	html := string(data)
	return parseReksadanaHTML(html), "kontan", false
}

func parseReksadanaHTML(html string) []models.ReksadanaItem {
	var reksa []models.ReksadanaItem

	// Find table start - look for typical table class
	tableStart := strings.Index(html, `<table class="table">`)
	if tableStart == -1 {
		tableStart = strings.Index(html, `<table class="table table-striped">`)
	}
	if tableStart == -1 {
		return fallbackReksadana()
	}

	html = html[tableStart:]
	var rowPos int
	count := 0

	for count < 30 {
		// Find <tr> 
		rowPos = strings.Index(html, "<tr")
		if rowPos == -1 {
			break
		}
		html = html[rowPos:]
		rowEnd := strings.Index(html, "</tr>")
		if rowEnd == -1 {
			break
		}
		row := html[:rowEnd+5]

		// Extract 6 <td> fields: Name, Type, NAB, Return1M, YTD, 1Y
		tds := extractTDs(row)
		if len(tds) >= 6 {
			name := strings.TrimSpace(tds[0])
			if strings.Contains(name, "Total") || name == "" {
				html = html[rowEnd+5:]
				continue
			}

			typ := strings.TrimSpace(tds[1])
			nav, _ := strconv.ParseFloat(strings.TrimSpace(tds[2]), 64)
			r1m, _ := strconv.ParseFloat(strings.TrimSpace(tds[3]), 64)
			rytd, _ := strconv.ParseFloat(strings.TrimSpace(tds[4]), 64)
			r1y, _ := strconv.ParseFloat(strings.TrimSpace(tds[5]), 64)

			reksa = append(reksa, models.ReksadanaItem{
				Name:        name,
				Type:        typ,
				NABPerUnit:  nav,
				Return1M:    r1m,
				ReturnYTD:   rytd,
				Return1Y:    r1y,
				Source:      "kontan_scrape",
			})
			count++
		}
		html = html[rowEnd+5:]
	}

	return reksa
}

func extractTDs(row string) []string {
	var tds []string
	pos := 0
	for {
		tdStart := strings.Index(row[pos:], "<td")
		if tdStart == -1 {
			break
		}
		tdStart += pos
		tdContentStart := strings.Index(row[tdStart:], ">") + tdStart + 1

		tdEnd := strings.Index(row[tdContentStart:], "</td>")
		if tdEnd == -1 {
			break
		}
		tdEnd += tdContentStart

		tds = append(tds, row[tdContentStart:tdEnd])
		pos = tdEnd
	}
	return tds
}

func fallbackReksadana() []models.ReksadanaItem {
	return []models.ReksadanaItem{
		{
			Name:        "Manulife Saham Syariah Asia Garuda",
			Type:        "Saham Syariah",
			NABPerUnit:  1234.56,
			Return1M:    2.34,
			ReturnYTD:   15.67,
			Return1Y:    25.89,
			Source:      "sample",
		},
		{
			Name:        "Sucorinvest Sharia Equity Fund",
			Type:        "Saham Syariah",
			NABPerUnit:  987.32,
			Return1M:    1.89,
			ReturnYTD:   12.34,
			Return1Y:    22.45,
			Source:      "sample",
		},
		// 8 more sample data...
		{
			Name:        "Mandiri Investa Atraktif",
			Type:        "Campuran",
			NABPerUnit:  1456.78,
			Return1M:    1.23,
			ReturnYTD:   8.90,
			Return1Y:    18.76,
			Source:      "sample",
		},
	}
}

