package service

import (
	"github.com/osamikoyo/meteor/internal/data"
	"github.com/osamikoyo/meteor/internal/data/models"
)

type WeatherStorage struct {
	ST data.Storage
}

type Service interface {
	Add(models.Day) error
}

func (s *WeatherStorage) Add(w models.Day) error {
	return s.ST.Save(w)
}
