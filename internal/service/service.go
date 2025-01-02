package service

import (
	"github.com/osamikoyo/meteor/internal/data/models"
)

type Service interface {
	Add(response models.WeatherResponses) error
	GetByRange(date1 string, date2 string) (models.Period, error)
	GetByDay(date string) (models.Day, error)
}
