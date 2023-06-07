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
	Slowest           bool
	Fastest           bool
	AccessTimeCounter int64
	SlowestCounter    int64
	FastestCounter    int64
}

// 1. Получить время доступа к определенному сайту.
// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
// 4. Администраторы, которые хотят получать статистику количества запросов пользователей
// по трем вышеперечисленным эндпойнтам.
type WebsiteStorage interface {
	// staff only
	UpdateAccessTime(ctx context.Context, lastCheck time.Time, accessTime int64) error

	// for users
	GetAccessTime(ctx context.Context, url string) (time.Duration, error)
	GetMinAccessURL(ctx context.Context) (string, error)
	GetMaxAccessURL(ctx context.Context) (string, error)

	// for admins
	GetAccessTimeStats(ctx context.Context, url string) (int64, error)
	GetMinAccessURLStats(ctx context.Context) (string, error)
	GetMaxAccessURLStats(ctx context.Context) (string, error)
}

type Websites struct {
	wstore WebsiteStorage
}

func NewWebsites(wstore WebsiteStorage) *Websites {
	return &Websites{
		wstore: wstore,
	}
}

func (ws *Websites) UpdateAccessTime(ctx context.Context, lastCheck time.Time, accessTime int64) error {
	return nil
}
