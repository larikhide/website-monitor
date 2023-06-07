package stats

import (
	"context"
	"fmt"
)

type Stats struct {
	SlowestCounter int64
	FastestCounter int64
}

// for admins
// 2. Получить статистику запросов к минимальному времени доступа.
// 3. Получить статистику запросов к максимальному времени доступа.
type StatsStorage interface {
	GetMinAccessURLStats(ctx context.Context) (string, error)
	GetMaxAccessURLStats(ctx context.Context) (string, error)
}

type Statistics struct {
	statstore StatsStorage
}

func (ss *Statistics) ReadMinAccessURLStats(ctx context.Context) (string, error) {
	minAccessStats, err := ss.statstore.GetMinAccessURLStats(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return minAccessStats, nil
}
func (ss *Statistics) GetMaxAccessURLStats(ctx context.Context) (string, error) {
	maxAccessStats, err := ss.statstore.GetMaxAccessURLStats(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return maxAccessStats, nil
}
