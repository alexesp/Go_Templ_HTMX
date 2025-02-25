package handler

import "github.com/go-chi/chi"

func RegisterRoutes(r *chi.Mux) {
	home := homeHandler{}
	r.Get("/", home.handleIndex)
}
