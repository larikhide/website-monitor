package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

//TODO: по хорошему описать представления респонсов
// type Website struct {
// 	URL        string        `json:"url"`
// 	AccessTime time.Duration `json:"access_time"`
// }

// type Stats struct {
// 	SlowestCounter int64 `json:"slowest_counter"`
// 	FastestCounter int64 `json:"fastest_counter"`
// }

type UserHandlers struct {
	websiteDB *website.Websites
	statsDB   *stats.Statistics
}

func NewUserHandlers(wdb *website.Websites, sdb *stats.Statistics) *UserHandlers {
	uh := &UserHandlers{
		websiteDB: wdb,
		statsDB:   sdb,
	}
	return uh
}

// /ping?url=...
func (uh *UserHandlers) GetPingURLHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	wsite, err := uh.websiteDB.Read(r.Context(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	wsite.PingRequestsCounter++
	if err := uh.websiteDB.Update(r.Context(), wsite); err != nil {
		log.Printf("failed to update ping request counter: %v", err)
	}

	_ = json.NewEncoder(w).Encode(
		struct {
			URL  string        `json:"url"`
			Ping time.Duration `json:"ping"`
		}{URL: wsite.URL, Ping: wsite.Ping},
	)
}

// /minping
func (uh *UserHandlers) GetMinPingURLHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := uh.statsDB.Read(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	stts.MinPingRequestCount++
	if err := uh.statsDB.Update(r.Context(), stts); err != nil {
		log.Printf("failed to update min ping request counter: %v", err)
	}
	_ = json.NewEncoder(w).Encode(struct {
		URL  string        `json:"url"`
		Ping time.Duration `json:"ping"`
	}{URL: stts.MinPingURL, Ping: stts.MinPing})
}

// /maxping
func (uh *UserHandlers) GetMaxPingURLHandler(w http.ResponseWriter, r *http.Request) {
	stts, err := uh.statsDB.Read(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}
	stts.MaxPingRequestCount++
	if err := uh.statsDB.Update(r.Context(), stts); err != nil {
		log.Printf("failed to update max ping request counter: %v", err)
	}
	_ = json.NewEncoder(w).Encode(struct {
		URL  string        `json:"url"`
		Ping time.Duration `json:"ping"`
	}{URL: stts.MinPingURL, Ping: stts.MaxPing})
}
