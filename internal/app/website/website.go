package website

import (
	"context"
	"fmt"
	"time"
)

type Website struct {
	URL        string
	Status     bool
	LastCheck  time.Time
	AccessTime time.Duration
	Slowest    bool
	Fastest    bool
}

// 1. Получить время доступа к определенному сайту.
// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
type WebsiteStorage interface {
	Read(ctx context.Context, url string) (*Website, error)
	UpdateAccessTime(ctx context.Context, lastCheck time.Time, accessTime int64) error

	GetAccessTime(ctx context.Context, URL string) (int64, error)
	GetMinAccessURL(ctx context.Context) (string, error)
	GetMaxAccessURL(ctx context.Context) (string, error)
}

type Websites struct {
	wstore WebsiteStorage
}

func NewWebsites(wstore WebsiteStorage) *Websites {
	return &Websites{
		wstore: wstore,
	}
}

func (ws *Websites) Read(ctx context.Context, url string) (*Website, error) {
	w, err := ws.wstore.Read(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return w, nil
}

func (ws *Websites) Update(ctx context.Context, lastCheck time.Time, accessTime int64) error {
	return nil
}

func (ws *Websites) GetAccessTime(ctx context.Context, URL string) (int64, error) {
	atime, err := ws.wstore.GetAccessTime(ctx, URL)
	if err != nil {
		return 0, fmt.Errorf("get from db errors: %w", err)
	}
	return atime, nil
}

func (ws *Websites) GetMinAccessURL(ctx context.Context) (string, error) {
	minAccess, err := ws.wstore.GetMinAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return minAccess, nil
}

func (ws *Websites) GetMaxAccessURL(ctx context.Context) (string, error) {
	maxAccess, err := ws.wstore.GetMaxAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return maxAccess, nil
}
