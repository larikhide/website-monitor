package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type AdminHandlers struct {
	websiteDB *website.Websites
	statsDB   *stats.Statistics
}

func NewAdminHandlers(wdb *website.Websites, sdb *stats.Statistics) *AdminHandlers {
	ah := &AdminHandlers{
		websiteDB: wdb,
		statsDB:   sdb,
	}
	return ah
}

func (as *AdminHandlers) GetPingRequestCountHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	wsite, err := as.websiteDB.Read(r.Context(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}
	_ = json.NewEncoder(w).Encode(struct {
		URL                 string `json:"url"`
		PingRequestsCounter int64  `json:"ping_requests"`
	}{URL: wsite.URL, PingRequestsCounter: wsite.PingRequestsCounter})
}

func (as *AdminHandlers) GetMinPingStatsHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := as.statsDB.Read(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}
	_ = json.NewEncoder(w).Encode(struct {
		MinPingRequests int64 `json:"min_ping_requests"`
	}{MinPingRequests: stts.MinPingRequestCount})
}

func (as *AdminHandlers) GetMaxPingStatsHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := as.statsDB.Read(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(struct {
		MaxPingRequests int64 `json:"max_ping_requests"`
	}{MaxPingRequests: stts.MaxPingRequestCount})
}

func (as *AdminHandlers) GetAllStats(w http.ResponseWriter, r *http.Request) {
	stts, err := as.statsDB.Read(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(stts)
}
