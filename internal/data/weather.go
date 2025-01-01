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
			precipitation    float32
			tempday          float32
			tempnight        float32
			countsnowtrue    uint16
			countsnowfalse   uint16
			countcloudytrue  uint16
			countcloudyfalse uint16
			cloudy           bool
			snow             bool
		)

		for _, day := range days {
			windspeed = windspeed + float32(day.WindSpeed)
			precipitation = precipitation + float32(day.Precipitation)
			tempday = tempday + float32(day.TempDay)
			tempnight = tempnight + float32(day.TempNight)
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
		precipitation = precipitation / float32(len(days))
		tempday = tempday / float32(len(days))
		tempnight = tempnight / float32(len(days))

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
			FirstDate:           days[0].Date,
			SecondDate:          days[len(days)-1].Date,
			MiddlePrecipitation: precipitation,
			MiddleTempDay:       tempday,
			MiddleTempNight:     tempnight,
			MiddleWindSpeed:     windspeed,
			Snow:                snow,
			Cloudy:              cloudy,
		}

		return per, nil
	}
}
