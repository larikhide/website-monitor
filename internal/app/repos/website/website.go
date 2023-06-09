package website

import (
	"context"
	"fmt"
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

func (ws *Websites) Read(ctx context.Context, url string) (*Website, error) {
	wsite, err := ws.wstore.Read(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return wsite, nil
}

func (ws *Websites) Update(ctx context.Context, website *Website) error {
	err := ws.wstore.Update(ctx, website)
	if err != nil {
		return fmt.Errorf("update website error: %w", err)
	}
	return nil
}
