package website

import (
	"context"
	"fmt"
	"time"
)

func (ws *Websites) ReadAccessTime(ctx context.Context, url string) (time.Duration, error) {
	atime, err := ws.wstore.GetAccessTime(ctx, url)
	if err != nil {
		return 0, fmt.Errorf("get from db errors: %w", err)
	}
	return atime, nil
}

func (ws *Websites) ReadMinAccessURL(ctx context.Context) (string, error) {
	minAccess, err := ws.wstore.GetMinAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return minAccess, nil
}

func (ws *Websites) ReadMaxAccessURL(ctx context.Context) (string, error) {
	maxAccess, err := ws.wstore.GetMaxAccessURL(ctx)
	if err != nil {
		return "", fmt.Errorf("get from db errors: %w", err)
	}
	return maxAccess, nil
}
