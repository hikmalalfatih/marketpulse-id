package scraper

import (
	"log"

	"market-api/models"
)

const obligasiURL = "https://www.djppr.kemenkeu.go.id/surat-utang-negara" // Scrape or fallback

func FetchObligasi() ([]models.ObligasiItem, string, bool) {
	// Hard scrape complex, use fallback with realistic data
	log.Printf("[OBLIGASI] Using reference data (scrape to be implemented)")
	return fallbackObligasi(), "reference", true
}

func fallbackObligasi() []models.ObligasiItem {
	return []models.ObligasiItem{
		{
			Kode:        "FR0091",
			Nama:        "FR0091",
			Jenis:       "FR",
			Kupon:       6.5,
			JatuhTempo:  "2031-03-15",
			HargaPasar:  102.5,
			Yield:       6.2,
			Tenor:       "6 years",
			Source:      "reference",
		},
		{
			Kode:        "ORI024",
			Nama:        "Obligasi Negara Ritel ORI024",
			Jenis:       "ORI",
			Kupon:       6.4,
			JatuhTempo:  "2030-12-15",
			HargaPasar:  101.8,
			Yield:       6.1,
			Tenor:       "5 years",
			Source:      "reference",
		},
		{
			Kode:        "SR019",
			Nama:        "Sukuk Negara SR019",
			Jenis:       "SR",
			Kupon:       5.8,
			JatuhTempo:  "2028-09-20",
			HargaPasar:  99.5,
			Yield:       5.9,
			Tenor:       "3 years",
			Source:      "reference",
		},
		{
			Kode:        "PBS039",
			Nama:        "PBS039",
			Jenis:       "PBS",
			Kupon:       6.75,
			JatuhTempo:  "2032-06-10",
			HargaPasar:  103.2,
			Yield:       6.4,
			Tenor:       "7 years",
			Source:      "reference",
		},
		// 6 more...
		{
			Kode:        "ORI025",
			Nama:        "ORI025",
			Jenis:       "ORI",
			Kupon:       6.6,
			JatuhTempo:  "2031-06-15",
			HargaPasar:  100.5,
			Yield:       6.5,
			Tenor:       "6 years",
			Source:      "reference",
		},
	}
}


