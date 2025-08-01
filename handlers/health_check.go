package handlers

import (
	"net/http"
	"vaqua/db"
	"log"
)

type HealthHandler struct{}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(); err != nil {
		log.Printf("DB connection error: %v", err)
		http.Error(w, "Database not reachable", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}