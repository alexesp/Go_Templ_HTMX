package app

import (
	"context"
	"log/slog"
	"net/http"
)

func Run(ctx context.Context) error {
	s := http.Server{
		Addr: ":8080",
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server")
		s.Shutdown(ctx)
	}()
	slog.Info("Starting server")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
