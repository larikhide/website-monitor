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

func NewHandlers(wdb *website.Websites, sdb *stats.Statistics) *UserHandlers {
	hs := &UserHandlers{
		websiteDB: wdb,
		statsDB:   sdb,
	}
	return hs
}

// /ping?url=...
func (uh *UserHandlers) GetPingURLHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	website, err := uh.websiteDB.Read(r.Context(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	website.PingRequestsCounter++
	uh.websiteDB.Update(r.Context(), website)

	_ = json.NewEncoder(w).Encode(
		struct {
			URL  string        `json:"url"`
			Ping time.Duration `json:"ping"`
		}{URL: website.URL, Ping: website.Ping},
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
