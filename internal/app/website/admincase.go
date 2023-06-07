package website

import (
	"context"
	"fmt"
)

func (ws *Websites) ReadAccessTimeStats(ctx context.Context, url string) (int64, error) {
	atimeStats, err := ws.wstore.GetAccessTimeStats(ctx, url)
	if err != nil {
		return 0, fmt.Errorf("get from db errors: %w", err)
	}
	return atimeStats, nil
}

func (ws *Websites) ReadMinAccessURLStats(ctx context.Context) (string, error) {
	minAccessStats, err := ws.wstore.GetMinAccessURLStats(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return minAccessStats, nil
}
func (ws *Websites) GetMaxAccessURLStats(ctx context.Context) (string, error) {
	maxAccessStats, err := ws.wstore.GetMaxAccessURLStats(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return maxAccessStats, nil
}
