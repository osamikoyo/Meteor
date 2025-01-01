package data

import (
	"github.com/osamikoyo/meteor/internal/data/models"
	"time"
)

const TIMELAYOUT = "20.11.20"

func (d Database) Save(weather models.Day) error {
	return d.Create(&weather).Error
}

func (d Database) Get(date time.Time) (models.Day, error) {
	var day models.Day
	err := d.Where(
		&models.Day{
			Date: date,
		},
	).Find(&day).Error
	return day, err
}

func (d Database) GetByRange(date1 time.Time, date2 time.Time) (models.Period, error) {
	fdate1 := date1.Format(TIMELAYOUT)
	fdate2 := date2.Format(TIMELAYOUT)

	var days []models.Day

	datefirst, err := time.Parse(TIMELAYOUT, fdate1)
	if err != nil {
		return models.Period{}, err
	}
}
