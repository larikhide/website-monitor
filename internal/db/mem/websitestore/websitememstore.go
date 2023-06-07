package memstore

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

var _ website.WebsiteStorage = &MemDB{}

type MemDB struct {
	sync.Mutex
	m map[string]website.Website
}

func NewWebsites() *MemDB {
	return &MemDB{
		m: make(map[string]website.Website),
	}
}

func (m *MemDB) UpdateAccessTime(ctx context.Context, url string, ut time.Time, ping time.Duration) error {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// uid := uuid.New()
	// u.ID = uid
	// m.m[u.ID] = u
	m.m[url] = website.Website{
		LastCheck:  ut,
		AccessTime: ping,
	}

	return nil

}

func (m *MemDB) GetAccessTime(ctx context.Context, url string) (time.Duration, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	t, ok := m.m[url]
	if ok {
		return t.AccessTime, nil
	}
	return 0, sql.ErrNoRows
}

func (m *MemDB) GetAccessTimeStats(ctx context.Context, url string) (int64, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	t, ok := m.m[url]
	if ok {
		return t.AccessTimeCounter, nil
	}
	return 0, sql.ErrNoRows
}

func (m *MemDB) GetMaxAccessURL(ctx context.Context) (string, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	url := m.findMaxAccessTimeURL()
	return url, sql.ErrNoRows
}

func (m *MemDB) GetMinAccessURL(ctx context.Context) (string, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	url := m.findMinAccessTimeURL()
	return url, sql.ErrNoRows
}

// TODO: must lock or not?
func (m *MemDB) findMaxAccessTimeURL() string {
	var maxURL string
	var maxAccessTime time.Duration

	for _, w := range m.m {
		if w.AccessTime > maxAccessTime {
			maxURL = w.URL
			maxAccessTime = w.AccessTime
		}
	}
	return maxURL
}

// TODO: must lock or not?
func (m *MemDB) findMinAccessTimeURL() string {
	var minURL string
	var minAccessTime time.Duration

	for _, w := range m.m {
		if w.AccessTime == 0 || w.AccessTime < minAccessTime {
			minURL = w.URL
			minAccessTime = w.AccessTime
		}
	}
	return minURL
}
