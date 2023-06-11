package routergin

import (
	"github.com/gin-gonic/gin"
	"github.com/larikhide/website-monitor/internal/api/handlers"
)

type RouterGin struct {
	*gin.Engine
	uh *handlers.UserHandlers
	ah *handlers.AdminHandlers
}

func NewRouterGin(uh *handlers.UserHandlers, ah *handlers.AdminHandlers) *RouterGin {
	r := gin.Default()
	ret := &RouterGin{
		uh: uh,
		ah: ah,
	}

	r.GET("/ping", ret.getPing)
	r.GET("/minping", ret.getMinPingURL)
	r.GET("/maxping", ret.getMaxPingURL)

	r.GET("/ping/stats", ret.getPingRequestCount)
	r.GET("/minping/stats", ret.getMinPingStats)
	r.GET("/maxping/stats", ret.getMaxPingStats)

	ret.Engine = r
	return ret
}

func (r *RouterGin) getPing(c *gin.Context) {
	r.uh.GetPingURLHandler(c.Writer, c.Request)
}

func (r *RouterGin) getMinPingURL(c *gin.Context) {
	r.uh.GetMinPingURLHandler(c.Writer, c.Request)
}

func (r *RouterGin) getMaxPingURL(c *gin.Context) {
	r.uh.GetMaxPingURLHandler(c.Writer, c.Request)
}

func (r *RouterGin) getPingRequestCount(c *gin.Context) {
	r.ah.GetPingRequestCountHandler(c.Writer, c.Request)
}

func (r *RouterGin) getMinPingStats(c *gin.Context) {
	r.ah.GetMinPingStatsHandler(c.Writer, c.Request)
}

func (r *RouterGin) getMaxPingStats(c *gin.Context) {
	r.ah.GetMaxPingStatsHandler(c.Writer, c.Request)
}
