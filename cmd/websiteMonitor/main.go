package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/larikhide/website-monitor/internal/api/handler"
	"github.com/larikhide/website-monitor/internal/api/server"
	"github.com/larikhide/website-monitor/internal/app"
	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
	"github.com/larikhide/website-monitor/internal/db/mem/statstore"
	"github.com/larikhide/website-monitor/internal/db/mem/websitestore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	websiteStore := websitestore.NewWebsites()
	statsStore := statstore.NewStatistics()
	a := app.NewApp(websiteStore, statsStore)
	ws := website.NewWebsites(websiteStore)
	ss := stats.NewStatistics(statsStore)
	h := handler.NewHandlers(ws, ss)
	//router := router.NewRouter(h)
	srv := server.NewServer(":8000", router)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}