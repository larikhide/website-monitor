package stats

import (
	"context"
	"fmt"
	"time"
)

type Stats struct {
	MinPingURL          string
	MaxPingURL          string
	MinPing             time.Duration
	MaxPing             time.Duration
	MinPingRequestCount int64
	MaxPingRequestCount int64
}

type StatsRepository interface {
	Read(ctx context.Context) (*Stats, error)
	Update(ctx context.Context, upds *Stats) error
}

type Statistics struct {
	statstore StatsRepository
}

func NewStatistics(sstore StatsRepository) *Statistics {
	return &Statistics{
		statstore: sstore,
	}
}

func (st *Statistics) Read(ctx context.Context) (*Stats, error) {
	stat, err := st.statstore.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("read stats error: %w", err)
	}
	return stat, nil
}

func (st *Statistics) Update(ctx context.Context, stats *Stats) error {
	err := st.statstore.Update(ctx, stats)
	if err != nil {
		return fmt.Errorf("update stats error: %w", err)
	}
	return nil
}
