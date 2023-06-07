package memstore

import (
	"context"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/website"
)

var _ website.WebsiteStorage = &Websites{}

type Websites struct {
	sync.Mutex
	m map[string]website.Website
}

func NewWebsites() *Websites {
	return &Websites{
		m: make(map[string]website.Website),
	}
}

func (ws *Websites) Read(ctx context.Context, url string) (*Website, error) {

}
func (ws *Websites) UpdateAccessTime(ctx context.Context, lastCheck time.Time, accessTime int64) error {

}
