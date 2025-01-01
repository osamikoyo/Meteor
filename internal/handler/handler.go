package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/osamikoyo/meteor/internal/data"
	"github.com/osamikoyo/meteor/pkg/loger"
	"net/http"
)

type Handler struct {
	ST data.Storage
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

}
