package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/larikhide/website-monitor/internal/api/handlers"
	"github.com/larikhide/website-monitor/internal/api/router/routergin"
	"github.com/larikhide/website-monitor/internal/api/server"
	"github.com/larikhide/website-monitor/internal/app/repos/stats"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
	monitoring "github.com/larikhide/website-monitor/internal/app/services/monitoringService"
	app "github.com/larikhide/website-monitor/internal/app/starter"
	"github.com/larikhide/website-monitor/internal/db/mem/statstore"
	"github.com/larikhide/website-monitor/internal/db/mem/websitestore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	websiteStore := websitestore.NewWebsites()

	filePath := "./websites.txt"

	websiteStore.PopulateFromSourceFile(ctx, filePath)
	log.Printf("website store : %v", websiteStore)
	statsStore := statstore.NewStatistics()
	monitor := monitoring.NewMonitoringService(websiteStore, statsStore)
	a := app.NewApp(websiteStore, statsStore, *monitor)
	ws := website.NewWebsites(websiteStore)
	ss := stats.NewStatistics(statsStore)
	uh := handlers.NewUserHandlers(ws, ss)
	ah := handlers.NewAdminHandlers(ws, ss)
	router := routergin.NewRouterGin(uh, ah)
	srv := server.NewServer(":8000", router)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
