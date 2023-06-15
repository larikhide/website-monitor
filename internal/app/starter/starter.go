package app

import (
	"context"
	"sync"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
	monitoring "github.com/larikhide/website-monitor/internal/app/services/monitoringService"
)

type App struct {
	wr *website.Websites
	sr *stats.Statistics
	mn *monitoring.MonitoringService
}

func NewApp(wr website.WebsiteRepository, sr stats.StatsRepository, mn monitoring.MonitoringService) *App {
	a := &App{
		wr: website.NewWebsites(wr),
		sr: stats.NewStatistics(sr),
		mn: monitoring.NewMonitoringService(wr, sr),
	}
	return a
}

type APIServer interface {
	Start(wr *website.Websites, sr *stats.Statistics)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	a.mn.PingWebsites(ctx)
	hs.Start(a.wr, a.sr)
	<-ctx.Done()
	hs.Stop()
}
