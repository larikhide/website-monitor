package website

import (
	"context"
	"fmt"
	"time"
)

type Website struct {
	URL        string
	ShortURL   string
	Status     bool
	LastCheck  time.Time
	AccessTime int64
	MinTime    int64
	MaxTime    int64
}

// 1. Получить время доступа к определенному сайту.
// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
type WebsiteStorage interface {
	GetAccessTime(ctx context.Context, shortURL string) (int64, error)
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

func (ws *Websites) GetAccessTime(ctx context.Context, shortURL string) (int64, error) {
	atime, err := ws.wstore.GetAccessTime(ctx, shortURL)
	if err != nil {
		return 0, fmt.Errorf("get access time error: %w", err)
	}
	return atime, nil
}

func (ws *Websites) GetMinAccessURL(ctx context.Context) (string, error) {
	minAccess, err := ws.wstore.GetMinAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("min access time error: %w", err)
	}
	return minAccess, nil
}

func (ws *Websites) GetMaxAccessURL(ctx context.Context) (string, error) {
	maxAccess, err := ws.wstore.GetMaxAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("max access time error: %w", err)
	}
	return maxAccess, nil
}
