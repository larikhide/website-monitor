package app

import (
	"context"
	"sync"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type App struct {
	ws *website.Websites
	ss *stats.Statistics
}

func NewApp(ws website.WebsiteStorage, ss stats.StatsStorage) *App {
	a := &App{
		ws: website.NewWebsites(ws),
		ss: stats.NewStatistics(ss),
	}
	return a
}

type APIServer interface {
	Start(ws *website.Websites)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.ws)
	<-ctx.Done()
	hs.Stop()
}
