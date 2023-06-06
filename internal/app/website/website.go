package website

import (
	"context"
	"time"
)

type Website struct {
	URL        string
	ShortName  string
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
	webstore WebsiteStorage
}

func NewUsers(webstore WebsiteStorage) *Websites {
	return &Websites{
		webstore: webstore,
	}
}
