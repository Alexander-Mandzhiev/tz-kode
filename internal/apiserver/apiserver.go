package apiserver

import (
	"context"
	"net/http"
	"tz-kode/internal/config"
)

type APIServer struct {
	httpserver *http.Server
}

func (s *APIServer) Start(cfg *config.Config, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:           cfg.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
		IdleTimeout:    cfg.IdleTimeout,
	}

	return s.httpserver.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
