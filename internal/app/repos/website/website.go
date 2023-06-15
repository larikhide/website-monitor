package website

import (
	"context"
	"fmt"
	"time"
)

type Website struct {
	Name                string
	URL                 string
	Status              bool //TODO: обработать ошибку, если не доступен
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

	//GetPingRequestCount(ctx context.Context, url string) (int, error)
	//IncrementPingRequestCount(ctx context.Context, url string) error

	FindMinPingWebsite(ctx context.Context) (*Website, error)
	FindMaxPingWebsite(ctx context.Context) (*Website, error)
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

func (ws *Websites) GetWebsitesList(ctx context.Context) ([]Website, error) {
	list, err := ws.wstore.GetWebsitesList(ctx)
	if err != nil {
		return nil, fmt.Errorf("get list of websites error: %w", err)
	}
	return list, nil
}

func (ws *Websites) FindMinPingWebsite(ctx context.Context) (*Website, error) {
	wsite, err := ws.wstore.FindMinPingWebsite(ctx)
	if err != nil {
		return nil, fmt.Errorf("get min ping website error: %w", err)
	}
	return wsite, nil
}

func (ws *Websites) FindMaxPingWebsite(ctx context.Context) (*Website, error) {
	wsite, err := ws.wstore.FindMaxPingWebsite(ctx)
	if err != nil {
		return nil, fmt.Errorf("get max ping website error: %w", err)
	}
	return wsite, nil
}
