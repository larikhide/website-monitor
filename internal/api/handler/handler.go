package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type Website struct {
	URL        string        `json:"url"`
	AccessTime time.Duration `json:"access_time"`
}

type Stats struct {
	SlowestCounter int64 `json:"slowest_counter"`
	FastestCounter int64 `json:"fastest_counter"`
}

type Handlers struct {
	websiteDB *website.Websites
	statsDB   *stats.Statistics
}

func NewHandlers(wdb *website.Websites, sdb *stats.Statistics) *Handlers {
	hs := &Handlers{
		websiteDB: wdb,
		statsDB:   sdb,
	}
	return hs
}

// /ping?url=...
func (hs *Handlers) ReadAccessTime(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	accessTime, err := hs.websiteDB.ReadAccessTime(r.Context(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(
		Website{
			URL:        url,
			AccessTime: accessTime,
		},
	)
}

// /minping
func (hs *Handlers) ReadMinAccessURL(w http.ResponseWriter, r *http.Request) {
	url, err := hs.websiteDB.ReadMinAccessURL(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}
	_ = json.NewEncoder(w).Encode(url)
}

// /maxping
func (hs *Handlers) ReadMaxAccessURL(w http.ResponseWriter, r *http.Request) {
	url, err := hs.websiteDB.ReadMaxAccessURL(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}
	_ = json.NewEncoder(w).Encode(url)
}
func (hs *Handlers) ReadAccessTimeStats(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(
		Website{
			URL:        r.URL.Query().Get("url"),
			AccessTime: 123,
		},
	)
}
func (hs *Handlers) ReadMinAccessURLStats(w http.ResponseWriter, r *http.Request) {}
func (hs *Handlers) ReadMaxAccessURLStats(w http.ResponseWriter, r *http.Request) {}
