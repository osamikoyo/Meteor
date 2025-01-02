package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h Handler) getByDay(w http.ResponseWriter, r *http.Request) error {
	dayNumber := chi.URLParam(r, "date")

	day, err := h.ST.GetByDay(dayNumber)
	if err != nil {
		return err
	}

	response, err := json.Marshal(&day)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, response)
	return err
}

func (h Handler) getByRange(w http.ResponseWriter, r *http.Request) error {
	dayNumber := chi.URLParam(r, "start")
	dayNumber2 := chi.URLParam(r, "end")

	period, err := h.ST.GetByRange(dayNumber, dayNumber2)

	response, err := json.Marshal(period)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, response)
	return err
}
