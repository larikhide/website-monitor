package memstore

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/website"
)

var _ website.WebsiteStorage = &Websites{}

type Websites struct {
	sync.Mutex
	m map[string]website.Website
}

func NewWebsites() *Websites {
	return &Websites{
		m: make(map[string]website.Website),
	}
}

// TODO: must lock or not?
func (ws *Websites) findMaxAccessTimeURL() string {
	var maxURL string
	var maxAccessTime time.Duration

	for _, w := range ws.m {
		if w.AccessTime > maxAccessTime {
			maxURL = w.URL
			maxAccessTime = w.AccessTime
		}
	}
	return maxURL
}

func (ws *Websites) GetAccessTime(ctx context.Context, url string) (time.Duration, error) {
	ws.Lock()
	defer ws.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	t, ok := ws.m[url]
	if ok {
		return t.AccessTime, nil
	}
	return 0, sql.ErrNoRows
}

func (ws *Websites) GetAccessTimeStats(ctx context.Context, url string) (int64, error) {
	ws.Lock()
	defer ws.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	t, ok := ws.m[url]
	if ok {
		return t.AccessTimeCounter, nil
	}
	return 0, sql.ErrNoRows
}

func (ws *Websites) GetMaxAccessURL(ctx context.Context) (string, error) {
	ws.Lock()
	defer ws.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	url := ws.findMaxAccessTimeURL()
	return url, sql.ErrNoRows
}
