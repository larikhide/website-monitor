package stats

import (
	"context"
)

type Stats struct {
	MinPingURL          string
	MaxPingURL          string
	MinPingRequestCount int
	MaxPingRequestCount int
}

// 2. Получить имя сайта с минимальным временем доступа.
// 3. Получить имя сайта с максимальным временем доступа.
// 5. Получить статистику запросов к минимальному времени доступа.
// 6. Получить статистику запросов к максимальному времени доступа.
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
