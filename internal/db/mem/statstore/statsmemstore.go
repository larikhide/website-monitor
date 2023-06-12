package statstore

import (
	"context"
	"sync"
	"time"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
)

var _ stats.StatsRepository = &MemDB{}

type MemDB struct {
	sync.Mutex
	s stats.Stats
}

// TODO: just mock for check. move to _test
// func NewStatistics() *MemDB {
// 	return &MemDB{
// 		s: stats.Stats{},
// 	}
// }

func NewStatistics() *MemDB {
	return &MemDB{
		s: stats.Stats{
			MinPingURL:          "google",
			MaxPingURL:          "yandex",
			MinPing:             123 * time.Millisecond,
			MaxPing:             333 * time.Millisecond,
			MinPingRequestCount: 54,
			MaxPingRequestCount: 48,
		},
	}
}

func (m *MemDB) Read(ctx context.Context) (*stats.Stats, error) {
	s := m.s
	// TODO: check for any error?
	return &s, nil
}

func (m *MemDB) Update(ctx context.Context, upds *stats.Stats) error {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	m.s = stats.Stats{
		MinPingURL:          upds.MinPingURL,
		MaxPingURL:          upds.MaxPingURL,
		MinPing:             upds.MaxPing,
		MaxPing:             upds.MaxPing,
		MinPingRequestCount: upds.MinPingRequestCount,
		MaxPingRequestCount: upds.MaxPingRequestCount,
	}
	return nil
}

func (m *MemDB) GetMinAccessURLStats(ctx context.Context) (int64, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	fastestCounter := m.s.MinPingRequestCount
	return fastestCounter, nil
}

func (m *MemDB) GetMaxAccessURLStats(ctx context.Context) (int64, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	slowestCounter := m.s.MaxPingRequestCount
	return slowestCounter, nil
}
