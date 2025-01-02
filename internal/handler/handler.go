package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osamikoyo/meteor/internal/service"
	"github.com/osamikoyo/meteor/pkg/loger"
	"net/http"
)

type Handler struct {
	ST service.Service
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func errorRoute(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			loger.New().Error().Err(err)
		}
	}
}

func (h Handler) RegisterRoutes(r *chi.Mux) {
	r.Use(middleware.Logger)

	r.Get("/api/get/{date}", errorRoute(h.getByDay))
	r.Get("/api/get/range/{first}-{second}", errorRoute(h.getByRange))
}
