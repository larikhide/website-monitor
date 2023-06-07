package statstore

import (
	"context"
	"sync"

	"github.com/larikhide/website-monitor/internal/app/repos/stats"
)

var _ stats.StatsStorage = &MemDB{}

type MemDB struct {
	sync.Mutex
	s stats.Stats
}

func NewStatistics() *MemDB {
	return &MemDB{
		s: stats.Stats{},
	}
}

func (m *MemDB) GetMinAccessURLStats(ctx context.Context) (int64, error) {
	m.Lock()
	defer m.Unlock()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	fastestCounter := m.s.FastestCounter
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
	slowestCounter := m.s.FastestCounter
	return slowestCounter, nil
}
