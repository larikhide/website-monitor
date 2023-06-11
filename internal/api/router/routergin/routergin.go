package routergin

import (
	"github.com/gin-gonic/gin"
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

	r.GET("/ping", ret.getAccessTime)
	r.GET("/maxping", ret.getMaxAccessURL)
	r.GET("/minping", ret.getMinAccessURL)

	r.GET("/ping/stats", ret.getAccessTimeStats)
	r.GET("/maxping/stats", ret.getMaxAccessURLStats)
	r.GET("/minping/stats", ret.getMinAccessURLStats)

	ret.Engine = r
	return ret
}

func (r *RouterGin) getAccessTime(c *gin.Context) {
	r.hs.ReadAccessTime(c.Writer, c.Request)
}
func (r *RouterGin) getMaxAccessURL(c *gin.Context) {
	r.hs.ReadMaxAccessURL(c.Writer, c.Request)
}
func (r *RouterGin) getMinAccessURL(c *gin.Context) {
	r.hs.ReadMinAccessURL(c.Writer, c.Request)
}

func (r *RouterGin) getAccessTimeStats(c *gin.Context) {
	r.hs.ReadAccessTimeStats(c.Writer, c.Request)
}
func (r *RouterGin) getMaxAccessURLStats(c *gin.Context) {
	r.hs.ReadMaxAccessURLStats(c.Writer, c.Request)
}
func (r *RouterGin) getMinAccessURLStats(c *gin.Context) {
	r.hs.ReadMinAccessURLStats(c.Writer, c.Request)
}
