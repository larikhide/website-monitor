package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/larikhide/website-monitor/internal/app/monitor"
	"github.com/larikhide/website-monitor/internal/app/website"
)

type URL struct {
	URL        string        `json:"url"`
	AccessTime time.Duration `json:"access_time"`
}

type Handlers struct {
	db *website.Websites
}

func NewHandlers(db *website.Websites) *Handlers {
	r := &Handlers{
		db: db,
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
	nbu, err := hs.db.GetMinAccessURL(ctx)
	if err != nil {
		return URL{}, fmt.Errorf("url not found: %w", err)
	}
	return URL{
		URL: nbu,
	}, nil
}

func (hs *Handlers) HandleMaxAccessURL(ctx context.Context) (URL, error) {
	nbu, err := hs.db.GetMaxAccessURL(ctx)
	if err != nil {
		return URL{}, fmt.Errorf("url not found: %w", err)
	}
	return URL{
		URL: nbu,
	}, nil
}
