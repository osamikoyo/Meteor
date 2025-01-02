package service

import (
	"github.com/osamikoyo/meteor/internal/data"
	"github.com/osamikoyo/meteor/internal/data/models"
	"time"
)

const CLOUDYRANGE = 3

type WeatherStorage struct {
	ST data.Storage
}

func New() *WeatherStorage {
	return &WeatherStorage{
		ST: data.New(),
	}
}

func (s *WeatherStorage) Add(response models.WeatherResponses) error {
	var w models.Day
	w.Temp = int16(response.Current.TempC)
	w.WindSpeed = uint64(response.Current.WindMph)
	w.Region = response.Location.Region

	if response.Current.Cloud < CLOUDYRANGE {
		w.Cloudy = false
	} else {
		w.Cloudy = true
	}

	date, err := time.Parse(data.TIMELAYOUT, time.Now().String())
	if err != nil {
		return err
	}

	w.Date = date

	return s.ST.Save(w)
}

func (s *WeatherStorage) GetByRange(date1 string, date2 string) (models.Period, error) {
	datefirst, err := time.Parse(data.TIMELAYOUT, date1)
	if err != nil {
		return models.Period{}, err
	}

	datesecond, err := time.Parse(data.TIMELAYOUT, date2)
	if err != nil {
		return models.Period{}, err
	}

	period, err := s.ST.GetByRange(datefirst, datesecond)
	return period, err
}

func (s *WeatherStorage) GetByDay(date string) (models.Day, error) {
	times, err := time.Parse(data.TIMELAYOUT, date)
	if err != nil {
		return models.Day{}, err
	}

	day, err := s.ST.Get(times)
	if err != nil {
		return models.Day{}, err
	}

	return day, nil
}
