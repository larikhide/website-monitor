package handler

import (
	"context"

	"github.com/larikhide/website-monitor/internal/app/website"
)

type Handler struct {
	ws *website.Websites
}

func NewHandler(ws *website.Websites) *Handler {
	h := &Handler{
		ws: ws,
	}
	return h
}

func (hs *Handler) GetAccessTime(ctx context.Context, shortURL string) (int64, error) {

}
