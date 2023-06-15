package monitoring

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type MonitoringService struct {
	websiteRepo website.WebsiteRepository
	statsRepo   stats.StatsRepository
	mu          *sync.Mutex
}

func NewMonitoringService(websiteRepo website.WebsiteRepository, statsRepo stats.StatsRepository) *MonitoringService {
	return &MonitoringService{
		websiteRepo: websiteRepo,
		statsRepo:   statsRepo,
		mu:          &sync.Mutex{},
	}
}

// TODO: add multithread
func (ms *MonitoringService) PingWebsites(ctx context.Context) error {
	// get list for pinging
	sites, err := ms.websiteRepo.GetWebsitesList(ctx)
	if err != nil {
		return ctx.Err()
	}

	// pingin all list
	for _, site := range sites {
		ping, err := PingURL(site.URL)
		if err != nil {
			return ctx.Err()
		}

		site.Ping = ping
		site.LastCheck = time.Now()

		// update every site into website repo
		err = ms.websiteRepo.Update(ctx, &site)
		if err != nil {
			return ctx.Err()
		}
	}

	// find min ping webite in updated website repo
	minPingWebsite, err := ms.websiteRepo.FindMinPingWebsite(ctx)
	if err != nil {
		return ctx.Err()
	}
	// find max ping webite in updated website repo
	maxPingWebsite, err := ms.websiteRepo.FindMaxPingWebsite(ctx)
	if err != nil {
		return ctx.Err()
	}

	oldStats, err := ms.statsRepo.Read(ctx)
	if err != nil {
		return ctx.Err()
	}

	newStats := &stats.Stats{
		MinPingURL:          minPingWebsite.URL,
		MaxPingURL:          maxPingWebsite.URL,
		MinPing:             minPingWebsite.Ping,
		MaxPing:             maxPingWebsite.Ping,
		MinPingRequestCount: oldStats.MinPingRequestCount,
		MaxPingRequestCount: oldStats.MaxPingRequestCount,
	}

	// update stats repo
	err = ms.statsRepo.Update(ctx, newStats)
	if err != nil {
		return ctx.Err()
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
