package handler

import (
	"net/http"

	"market-api/config"
	"market-api/server"
)

// app is reused while a Vercel Function instance stays warm, including its
// in-memory cache. Vercel may create a new instance at any time.
var app = server.NewHandler(config.LoadConfig())

// Handler is the Vercel Go Function entry point for every /api/* request.
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
