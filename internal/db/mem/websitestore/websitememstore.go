package websitestore

import (
	"context"
	"database/sql"
	"math"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

var _ website.WebsiteRepository = &MemDB{}

type MemDB struct {
	sync.Mutex
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
		Name:   "google",
		URL:    "https://www.google.com",
		Status: true,
	}

	websites["yandex"] = website.Website{
		Name:   "yandex",
		URL:    "https://www.ya.ru",
		Status: true,
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

func (m *MemDB) Update(ctx context.Context, wsite *website.Website) error {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	_, ok := m.m[wsite.Name]
	if ok {
		m.m[wsite.Name] = website.Website{
			Name:                wsite.Name,
			URL:                 wsite.URL,
			Status:              wsite.Status,
			LastCheck:           wsite.LastCheck,
			Ping:                wsite.Ping,
			PingRequestsCounter: wsite.PingRequestsCounter,
		}

		return nil
	}
	return sql.ErrNoRows
}

func (m *MemDB) GetWebsitesList(ctx context.Context) ([]website.Website, error) {
	m.Lock()
	defer m.Unlock()

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

func (m *MemDB) FindMinPingWebsite(ctx context.Context) (*website.Website, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return &website.Website{}, ctx.Err()
	default:
	}

	var minURL string
	minPing := time.Duration(math.MaxInt64)

	for _, w := range m.m {
		if w.Ping < minPing {
			minURL = w.Name
			minPing = w.Ping
		}
	}

	wsite := m.m[minURL]
	return &wsite, nil
}

func (m *MemDB) FindMaxPingWebsite(ctx context.Context) (*website.Website, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return &website.Website{}, ctx.Err()
	default:
	}

	var maxURL string
	var maxPing time.Duration

	for _, w := range m.m {
		if w.Ping > maxPing {
			maxURL = w.Name
			maxPing = w.Ping
		}
	}
	wsite := m.m[maxURL]
	return &wsite, nil
}
