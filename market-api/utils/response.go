package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"market-api/models"
)

func WriteJSON(w http.ResponseWriter, resp models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(resp)
}

func WriteError(w http.ResponseWriter, errMsg string) {
	resp := models.APIResponse{
		Success:   false,
		Error:     errMsg,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	WriteJSON(w, resp)
}

