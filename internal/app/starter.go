package app

import (
	"context"
	"sync"

	"github.com/larikhide/website-monitor/internal/app/monitor"
	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type App struct {
	ws *website.Websites
	ss *stats.Statistics
	mn *monitor.Monitor
}

func NewApp(ws website.WebsiteStorage, ss stats.StatsStorage, mn monitor.Monitor) *App {
	a := &App{
		ws: website.NewWebsites(ws),
		ss: stats.NewStatistics(ss),
		mn: monitor.NewMonitor(),
	}
	return a
}

type APIServer interface {
	Start(ws *website.Websites)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	a.mn.StartMonitoring()
	hs.Start(a.ws)
	<-ctx.Done()
	hs.Stop()
}
