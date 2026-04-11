package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"market-api/models"
)

func writeResponse(w http.ResponseWriter, success bool, data interface{}, cached bool, source string, updated time.Time) {
	resp := models.APIResponse{
		Success:   success,
		Data:      data,
		Cached:    cached,
		Source:    source,
		UpdatedAt: updated.Format(time.RFC3339),
	}
	writeJSON(w, resp)
}

func writeJSON(w http.ResponseWriter, resp models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
