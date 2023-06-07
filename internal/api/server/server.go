package server

import (
	"context"
	"net/http"
	"time"

	"github.com/larikhide/website-monitor/internal/app"
	"github.com/larikhide/website-monitor/internal/app/repos/website"
)

var _ app.APIServer = &Server{}

type Server struct {
	srv http.Server
	ws  *website.Websites
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start(ws *website.Websites) {
	s.ws = ws
	go s.srv.ListenAndServe()
}
