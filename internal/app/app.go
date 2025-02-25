package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/alexesp/Go_Templ_HTMX.git/internal/config"
	"github.com/alexesp/Go_Templ_HTMX.git/internal/handler"
	"github.com/go-chi/chi"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("Shutting down server")
		s.Shutdown(ctx)
	}()
	slog.Info("Starting server!", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
