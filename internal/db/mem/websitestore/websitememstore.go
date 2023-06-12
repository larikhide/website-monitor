package websitestore

import (
	"context"
	"database/sql"
	"math"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

var _ website.WebsiteRepository = &MemDB{}

type MemDB struct {
	// sync.Mutex
	m map[string]website.Website
}

/* func NewWebsites() *MemDB {
	return &MemDB{
		m: make(map[string]website.Website),
	}
} */

// TODO: just mock for check. remove to _test
func NewWebsites() *MemDB {
	websites := make(map[string]website.Website)
	websites["google"] = website.Website{
		URL:                 "https://www.google.com",
		Status:              true,
		LastCheck:           time.Now(),
		Ping:                time.Millisecond * 298,
		PingRequestsCounter: 10,
	}

	websites["yandex"] = website.Website{
		URL:                 "https://www.ya.ru",
		Status:              true,
		LastCheck:           time.Now(),
		Ping:                time.Millisecond * 132,
		PingRequestsCounter: 15,
	}

	return &MemDB{
		m: websites,
	}
}

func (m *MemDB) Read(ctx context.Context, url string) (*website.Website, error) {
	t, ok := m.m[url]
	if ok {
		return &t, nil
	}
	return &website.Website{}, sql.ErrNoRows
}

// TODO: check for correct funcionality
func (m *MemDB) Update(ctx context.Context, wsite *website.Website) error {
	// m.Lock()
	// defer m.Unlock()

	// select {
	// case <-ctx.Done():
	// 	return ctx.Err()
	// default:
	// }

	m.m[wsite.URL] = website.Website{
		URL:                 wsite.URL,
		Status:              wsite.Status,
		LastCheck:           wsite.LastCheck,
		Ping:                wsite.Ping,
		PingRequestsCounter: wsite.PingRequestsCounter,
	}

	return nil
}

func (m *MemDB) GetWebsitesList(ctx context.Context) ([]website.Website, error) {
	// m.Lock()
	// defer m.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	wlist := make([]website.Website, 0, len(m.m))
	for _, w := range m.m {
		wlist = append(wlist, w)
	}

	// TODO: check for created empty map?
	return wlist, nil
}

func (m *MemDB) GetMinAccessURL(ctx context.Context) (string, error) {
	// m.Lock()
	// defer m.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	url := m.findMinAccessTimeURL()
	return url, nil
}

func (m *MemDB) GetMaxAccessURL(ctx context.Context) (string, error) {
	// m.Lock()
	// defer m.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	url := m.findMaxAccessTimeURL()
	return url, nil
}

// TODO: must lock or not?
func (m *MemDB) findMinAccessTimeURL() string {
	var minURL string
	minAccessTime := time.Duration(math.MaxInt64)

	for _, w := range m.m {
		if w.Ping < minAccessTime {
			minURL = w.URL
			minAccessTime = w.Ping
		}
	}
	return minURL
}

// TODO: must lock or not?
func (m *MemDB) findMaxAccessTimeURL() string {
	var maxURL string
	var maxAccessTime time.Duration

	for _, w := range m.m {
		if w.Ping > maxAccessTime {
			maxURL = w.URL
			maxAccessTime = w.Ping
		}
	}
	return maxURL
}
