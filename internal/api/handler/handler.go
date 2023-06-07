package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/larikhide/website-monitor/internal/app/monitor"
	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type URL struct {
	URL        string        `json:"url"`
	AccessTime time.Duration `json:"access_time"`
}

type Handlers struct {
	websiteDB *website.Websites
	statsDB   *stats.Statistics
}

func NewHandlers(wdb *website.Websites, sdb *stats.Statistics) *Handlers {
	r := &Handlers{
		websiteDB: wdb,
		statsDB:   sdb,
	}
	return r
}

func (hs *Handlers) HandleAccessTime(ctx context.Context, u URL) (URL, error) {

	bu := website.Website{
		URL:        u.URL,
		AccessTime: u.AccessTime,
	}

	accessTime, err := monitor.AccessTime(bu.URL)
	if err != nil {
		return URL{}, fmt.Errorf("error when pinging: %w", err)
	}

	bu.AccessTime = accessTime

	return URL{
		URL:        bu.URL,
		AccessTime: bu.AccessTime,
	}, nil
}

func (hs *Handlers) HandleMinAccessURL(ctx context.Context) (URL, error) {
	nbu, err := hs.websiteDB.ReadMinAccessURL(ctx)
	if err != nil {
		return URL{}, fmt.Errorf("url not found: %w", err)
	}
	return URL{
		URL: nbu,
	}, nil
}

func (hs *Handlers) HandleMaxAccessURL(ctx context.Context) (URL, error) {
	nbu, err := hs.websiteDB.ReadMaxAccessURL(ctx)
	if err != nil {
		return URL{}, fmt.Errorf("url not found: %w", err)
	}
	return URL{
		URL: nbu,
	}, nil
}
