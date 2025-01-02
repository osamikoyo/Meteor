package service

import (
	"github.com/osamikoyo/meteor/internal/data"
	"github.com/osamikoyo/meteor/internal/data/models"
	"time"
)

const TIMELAYOUT = "11.11.24"

type WeatherStorage struct {
	ST data.Storage
}

func (s *WeatherStorage) Add(w models.Day) error {
	date, err := time.Parse(TIMELAYOUT, time.Now().String())
	if err != nil {
		return err
	}

	w.Date = date

	return s.ST.Save(w)
}

func (s *WeatherStorage) GetByRange(date1 string, date2 string) (models.Period, error) {
	datefirst, err := time.Parse(TIMELAYOUT, date1)
	if err != nil {
		return models.Period{}, err
	}

	datesecond, err := time.Parse(TIMELAYOUT, date2)
	if err != nil {
		return models.Period{}, err
	}

	period, err := s.ST.GetByRange(datefirst, datesecond)
	return period, err
}

func (s *WeatherStorage) GetByDay(date string) (models.Day, error) {
	times, err := time.Parse(TIMELAYOUT, date)
	if err != nil {
		return models.Day{}, err
	}

	day, err := s.ST.Get(times)
	if err != nil {
		return models.Day{}, err
	}

	return day, nil
}
