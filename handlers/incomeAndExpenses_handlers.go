package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vaqua/service"
)

type IncomeAndExpensesHandler struct {
	Service *service.IncomeAndExpensesService
}

func (h *IncomeAndExpensesHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	// Expect URL like: /accounts/1/summary
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "accounts" || parts[2] != "summary" {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	// Parse optional ?from=YYYY-MM-DD&to=YYYY-MM-DD
	fromT, fromStr, err := parseDate(r, "from")
	if err != nil {
		http.Error(w, "bad 'from' date", http.StatusBadRequest)
		return
	}
	toT, toStr, err := parseDate(r, "to")
	if err != nil {
		http.Error(w, "bad 'to' date", http.StatusBadRequest)
		return
	}

	// Call service
	summary, err := h.Service.GetSummary(int64(id), fromT, toT, fromStr, toStr)
	if err != nil {
		http.Error(w, "could not get summary", http.StatusInternalServerError)
		return
	}

	// Respond as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func parseDate(r *http.Request, key string) (*time.Time, *string, error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return nil, nil, nil
	}
	t, err := time.Parse("2006-01-02", raw)
	if err != nil {
		return nil, nil, err
	}
	return &t, &raw, nil
}
