package models

import "time"

type Day struct {
	Date          time.Time `gorm:"type:date;primaryKey"`
	TempNight     int16
	TempDay       int16
	WindSpeed     uint64
	Snow          bool
	Precipitation uint64
	Cloudy        bool
}

type Period struct {
	FirstDate           time.Time
	SecondDate          time.Time
	MiddleTemp          int16
	MiddleWindSpeed     uint32
	Show                bool
	MiddlePrecipitation uint64
	Cloudy              bool
}
