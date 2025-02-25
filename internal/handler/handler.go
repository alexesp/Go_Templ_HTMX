package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func RegisterRoutes(r *chi.Mux) {
	home := homeHandler{}
	r.Get("/", handler(home.handleIndex))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			hanleError(w, r, err)
		}
	}
}

func hanleError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("Error during request", slog.String("err", err.Error()))
}
