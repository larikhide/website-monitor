package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type MonitoringService struct {
	websiteRepo website.WebsiteRepository
}

func NewMonitoringService(websiteRepo website.WebsiteRepository) *MonitoringService {
	return &MonitoringService{
		websiteRepo: websiteRepo,
	}
}

// TODO: add concurrency and syncronization
func (ms *MonitoringService) PingWebsites(ctx context.Context) error {
	websites, err := ms.websiteRepo.GetWebsitesList(ctx)
	if err != nil {
		return fmt.Errorf("failed to get websites list: %w", err)
	}
	for _, wsite := range websites {
		website, err := ms.websiteRepo.Read(ctx, wsite.URL)
		if err != nil {
			return fmt.Errorf("failed to read website: %w", err)
		}

		pingDuration, err := PingURL(wsite.URL)
		if err != nil {
			return fmt.Errorf("failed to ping website %s: %w", wsite.URL, err)
		}

		//update Website ping info
		website.Ping = pingDuration
		website.LastCheck = time.Now()
		website.PingRequestsCounter++

		err = ms.websiteRepo.Update(ctx, website)
		if err != nil {
			return fmt.Errorf("failed to update website info %s: %w", wsite.URL, err)
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
