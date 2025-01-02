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

	datesecond, err := time.Parse(TIMELAYOUT, fdate2)
	if err != nil {
		return models.Period{}, err
	}

	if err := d.Where("date >= ? AND date <= ?", datefirst, datesecond).Find(&days).Error; err != nil {
		return models.Period{}, err
	} else {
		var (
			windspeed        float32
			temp             float32
			countsnowtrue    uint16
			countsnowfalse   uint16
			countcloudytrue  uint16
			countcloudyfalse uint16
			cloudy           bool
			snow             bool
		)

		for _, day := range days {
			windspeed = windspeed + float32(day.WindSpeed)
			temp = temp + float32(day.Temp)
			if day.Snow {
				countsnowtrue++
			} else {
				countsnowfalse++
			}
			if day.Cloudy {
				countcloudytrue++
			} else {
				countcloudyfalse++
			}
		}

		windspeed = windspeed / float32(len(days))
		temp = temp / float32(len(days))

		if countcloudyfalse < countcloudytrue {
			cloudy = true
		} else {
			cloudy = false
		}

		if countsnowfalse < countsnowtrue {
			snow = true
		} else {
			snow = false
		}

		per := models.Period{
			FirstDate:       days[0].Date,
			SecondDate:      days[len(days)-1].Date,
			MiddleTempDay:   temp,
			MiddleWindSpeed: windspeed,
			Snow:            snow,
			Cloudy:          cloudy,
		}

		return per, nil
	}
}
