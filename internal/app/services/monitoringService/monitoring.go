package monitoring

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type PingResult struct {
	wsite website.Website
	Error error
}

type MonitoringService struct {
	websiteRepo website.WebsiteRepository
	statsRepo   stats.StatsRepository
}

func NewMonitoringService(websiteRepo website.WebsiteRepository, statsRepo stats.StatsRepository) *MonitoringService {
	return &MonitoringService{
		websiteRepo: websiteRepo,
		statsRepo:   statsRepo,
	}
}

func (ms *MonitoringService) StartMonitoring(ctx context.Context) {
	//initial ping
	err := ms.PingAndUpdate(ctx)
	if err != nil {
		log.Printf("initial ping failed: %v", err)
	}

	// sync channel
	done := make(chan struct{})

	go func() {
		defer close(done)

		tick := time.NewTicker(time.Minute)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				err := ms.PingAndUpdate(ctx)
				if err != nil {
					log.Printf("Monitoring service error: %v", err)
				}
			}
		}

	}()
}

func (ms *MonitoringService) PingAndUpdate(ctx context.Context) error {
	// get list for pinging
	sites, err := ms.websiteRepo.GetWebsitesList(ctx)
	if err != nil {
		return err
	}

	//channel to receive ping results
	pingResults := make(chan PingResult)

	// ping all sites in the list
	for _, site := range sites {

		go func(site website.Website) {
			pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			ping, err := PingURL(site.URL, pingCtx)
			if err != nil {
				site.Status = false
				log.Printf("error pinging %v: %v", site.URL, err)
			} else {
				site.Status = true
			}

			site.Ping = ping
			site.LastCheck = time.Now()

			// Send the ping result to the channel
			pingResults <- PingResult{
				wsite: site,
				Error: err,
			}
		}(site)
	}

	ms.UpdateRepos(ctx, sites, pingResults)
	log.Printf("ping and update have been finished")
	return nil
}

func (ms *MonitoringService) UpdateRepos(ctx context.Context, sites []website.Website, pingResults chan PingResult) error {
	// Collect ping results from the channel
	for range sites {
		result := <-pingResults

		// Update every site into website repo
		err := ms.websiteRepo.Update(ctx, &result.wsite)
		if err != nil {
			return err
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

func PingURL(url string, ctx context.Context) (time.Duration, error) {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	return elapsed, nil
}
