package websitestore

import (
	"context"
	"database/sql"
	"math"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

var _ website.WebsiteStorage = &MemDB{}

type MemDB struct {
	sync.Mutex
	m map[string]website.Website
}

/* func NewWebsites() *MemDB {
	return &MemDB{
		m: make(map[string]website.Website),
	}
} */

func NewWebsites() *MemDB {
	websites := make(map[string]website.Website)
	websites["google"] = website.Website{
		URL:               "https://www.google.com",
		LastCheck:         time.Now(),
		AccessTime:        time.Millisecond * 298,
		AccessTimeCounter: 10,
	}

	websites["yandex"] = website.Website{
		URL:               "https://www.ya.ru",
		LastCheck:         time.Now(),
		AccessTime:        time.Millisecond * 132,
		AccessTimeCounter: 15,
	}

	return &MemDB{
		m: websites,
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

func (m *MemDB) Read(ctx context.Context, url string) (*website.Website, error) {
	t, ok := m.m[url]
	if ok {
		return &t, nil
	}
	return &website.Website{}, sql.ErrNoRows
}

func (m *MemDB) UpdateAccessCounter(ctx context.Context, url string) error {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	counter := m.m[url].AccessTimeCounter
	counter++
	m.m[url] = website.Website{
		AccessTimeCounter: counter,
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
		t.AccessTimeCounter++
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
	return url, nil
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
	return url, nil
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
	minAccessTime := time.Duration(math.MaxInt64)

	for _, w := range m.m {
		if w.AccessTime < minAccessTime {
			minURL = w.URL
			minAccessTime = w.AccessTime
		}
	}
	return minURL
}
