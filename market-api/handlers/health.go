package handlers

import (
 "encoding/json"
 "net/http"
 "time"

 "market-api/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
 resp := models.APIResponse{
  Success:   true,
  Data:      map[string]string{"status": "ok", "version": "1.0"},
  Source:    "internal",
  UpdatedAt: time.Now().Format(time.RFC3339),
 }
 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(resp)
}

