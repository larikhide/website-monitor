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
	}
}

func (ms *MonitoringService) PingWebsites(ctx context.Context) error {
	websites, err := ms.websiteRepo.GetWebsitesList(ctx)
	if err != nil {
		return fmt.Errorf("failed to get websites list: %w", err)
	}

	wg := new(sync.WaitGroup)
	maxGorutines := 10
	errCh := make(chan error, maxGorutines)

	for _, wsite := range websites {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			website, err := ms.websiteRepo.Read(ctx, url)
			if err != nil {
				errCh <- fmt.Errorf("failed to read website %s: %v", url, err)
				return
			}

			pingDuration, err := PingURL(url)
			if err != nil {
				errCh <- fmt.Errorf("failed to ping website %s: %v", url, err)
				return
			}

			ms.mu.Lock()
			defer ms.mu.Unlock()

			// Update Website ping info
			website.Ping = pingDuration
			website.LastCheck = time.Now()
			website.PingRequestsCounter++

			err = ms.websiteRepo.Update(ctx, website)
			if err != nil {
				errCh <- fmt.Errorf("failed to update website info %s: %v", url, err)
				return
			}
		}(wsite.URL)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
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
