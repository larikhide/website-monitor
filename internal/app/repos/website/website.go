package website

import (
	"context"
	"fmt"
	"time"
)

type Website struct {
	URL               string
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
	Read(ctx context.Context, url string) (*Website, error)
	UpdateAccessTime(ctx context.Context, url string, lastCheck time.Time, accessTime time.Duration) error
	UpdateAccessCounter(ctx context.Context, url string) error

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

// staff
func (ws *Websites) Read(ctx context.Context, url string) (*Website, error) {
	website, err := ws.wstore.Read(ctx, url)
	if err != nil {
		return &Website{}, fmt.Errorf("get from db errors: %w", err)
	}
	return website, err
}

func (ws *Websites) UpdateAccessTime(ctx context.Context, url string, lastCheck time.Time, accessTime time.Duration) error {
	//TODO
	return nil
}

func (ws *Websites) UpdateAccessCounter(ctx context.Context, url string) (*Website, error) {
	website, err := ws.wstore.Read(ctx, url)
	if err != nil {
		return &Website{}, fmt.Errorf("get from db errors: %w", err)
	}
	counter := website.AccessTimeCounter
	counter++
	return &Website{URL: website.URL,
		LastCheck:         website.LastCheck,
		AccessTime:        website.AccessTime,
		AccessTimeCounter: counter,
	}, nil
}
