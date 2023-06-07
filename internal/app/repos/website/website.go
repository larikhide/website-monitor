package website

import (
	"context"
	"time"
)

type Website struct {
	URL               string
	Status            bool
	LastCheck         time.Time
	AccessTime        time.Duration
	AccessTimeCounter int64
}

// for users
// 1. Получить время доступа к определенному сайту.
// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
type WebsiteStorage interface {
	//staff
	UpdateAccessTime(ctx context.Context, url string, lastCheck time.Time, accessTime time.Duration) error

	//for users
	GetAccessTime(ctx context.Context, url string) (time.Duration, error)
	GetMinAccessURL(ctx context.Context) (string, error)
	GetMaxAccessURL(ctx context.Context) (string, error)

	//for admins
	GetAccessTimeStats(ctx context.Context, url string) (int64, error)
}

type Websites struct {
	wstore WebsiteStorage
}

func NewWebsites(wstore WebsiteStorage) *Websites {
	return &Websites{
		wstore: wstore,
	}
}

func (ws *Websites) UpdateAccessTime(ctx context.Context, url string, lastCheck time.Time, accessTime time.Duration) error {
	//TODO
	return nil
}
