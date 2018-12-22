package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	return r
}
