package server

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	sv := http.Server{
		Addr:    ":8080",
		Handler: GetRoutes(),
	}

	go func() {
		defer wg.Done()
		if err := sv.ListenAndServe(); err != nil {
			return
		}
	}()

	<-ctx.Done()
	err := sv.Shutdown(ctx)
	if err != nil {
		slog.Error("error shutting down", "error", err)
	}
}
