package website

import (
	"context"
	"time"
)

type Website struct {
	URL                 string
	Status              bool
	LastCheck           time.Time
	Ping                time.Duration
	PingRequestsCounter int64
}

// for users
// 1. Получить время доступа к определенному сайту.
// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
type WebsiteStorage interface {
	//Create(ctx context.Context, website Website) (string, error)
	Read(ctx context.Context, url string) (*Website, error)
	Update(ctx context.Context, website *Website) error
	//Delete(ctx context.Context, url string) error

	GetWebsitesList(ctx context.Context) ([]Website, error)
}

type Websites struct {
	wstore WebsiteStorage
}

func NewWebsites(wstore WebsiteStorage) *Websites {
	return &Websites{
		wstore: wstore,
	}
}
