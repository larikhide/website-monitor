package routergin

import (
	"github.com/gin-gonic/gin"
	"github.com/larikhide/website-monitor/internal/api/handler"
)

type RouterGin struct {
	*gin.Engine
	hs *handler.Handlers
}

func NewRouterGin(hs *handler.Handlers) *RouterGin {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}

	r.GET("/ping/:url", ret.GetAccessTime)
	r.GET("/maxping", ret.GetMaxAccessURL)
	r.GET("/minping", ret.GetMinAccessURL)

	r.GET("/ping/:url/stats", ret.GetAccessTimeStats)
	r.GET("/maxping/stats", ret.GetMinAccessURLStats)
	r.GET("/minping/stats", ret.GetMinAccessURLStats)

	ret.Engine = r
	return ret
}

func (r *RouterGin) GetAccessTime(c *gin.Context)   {}
func (r *RouterGin) GetMaxAccessURL(c *gin.Context) {}
func (r *RouterGin) GetMinAccessURL(c *gin.Context) {}

func (r *RouterGin) GetAccessTimeStats(c *gin.Context)   {}
func (r *RouterGin) GetMaxAccessURLStats(c *gin.Context) {}
func (r *RouterGin) GetMinAccessURLStats(c *gin.Context) {}
