package app

import (
	"context"
	"sync"

	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

type App struct {
	ws *website.Websites
}

func NewApp(ws website.WebsiteStorage) *App {
	a := &App{
		ws: website.NewWebsites(ws),
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
