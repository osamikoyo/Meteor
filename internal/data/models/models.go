package models

import "time"

type Day struct {
	Date          time.Time `gorm:"type:date;primaryKey"`
	Temp          int16
	WindSpeed     uint64
	Snow          bool
	Precipitation uint64
	Cloudy        bool
	Region        string
}

type Period struct {
	FirstDate           time.Time
	SecondDate          time.Time
	MiddleTempDay       float32
	MiddleTempNight     float32
	MiddleWindSpeed     float32
	Snow                bool
	MiddlePrecipitation float32
	Cloudy              bool
}

type WeatherResponses struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	IsDay     int       `json:"is_day"`
	Condition Condition `json:"condition"`
	WindMph   float64   `json:"wind_mph"`
	PrecipMm  float64   `json:"precip_mm"`
	PrecipIn  float64   `json:"precip_in"`
	Cloud     int       `json:"cloud"`
}

type Condition struct {
	Text string `json:"text"`
	Code int    `json:"code"`
}
