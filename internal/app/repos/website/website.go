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

// 1. Получить время доступа к определенному сайту.
type WebsiteRepository interface {
	//Create(ctx context.Context, website Website) (string, error)
	Read(ctx context.Context, url string) (*Website, error)
	Update(ctx context.Context, website *Website) error
	//Delete(ctx context.Context, url string) error

	//
	GetWebsitesList(ctx context.Context) ([]Website, error)

	GetPingRequestCount(ctx context.Context, url string) (int, error)
	IncrementPingRequestCount(ctx context.Context, url string) error
}

type Websites struct {
	wstore WebsiteRepository
}

func NewWebsites(wstore WebsiteRepository) *Websites {
	return &Websites{
		wstore: wstore,
	}
}
