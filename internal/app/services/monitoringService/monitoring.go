package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type MonitoringService struct {
	websiteRepo website.WebsiteRepository
	mu          sync.Mutex
}

func NewMonitoringService(websiteRepo website.WebsiteRepository) *MonitoringService {
	return &MonitoringService{
		websiteRepo: websiteRepo,
		mu:          sync.Mutex{},
	}
}

// TODO: add multithread
func (ms *MonitoringService) PingWebsites(ctx context.Context) error {
	sites, err := ms.websiteRepo.GetWebsitesList(ctx)
	if err != nil {
		return fmt.Errorf("failed to get websites list: %w", err)
	}

	for _, site := range sites {
		ping, err := PingURL(site.URL)
		if err != nil {
			return fmt.Errorf("failed to get ping: %w", err)
		}

		site.Ping = ping
		site.LastCheck = time.Now()

		err = ms.websiteRepo.Update(ctx, &site)
		if err != nil {
			return fmt.Errorf("failed to update site info: %w", err)
		}
	}
	return nil
}

func PingURL(url string) (time.Duration, error) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	return elapsed, nil
}
