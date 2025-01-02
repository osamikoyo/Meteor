package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiRouter struct {
	URL string
	ctx context.Context
	key string
}

type WeatherResponse struct {
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

func New(url string, key string) *ApiRouter {
	return &ApiRouter{URL: url, key: key}
}

func (a *ApiRouter) ApiScanner(ch chan error) {
	req, err := http.NewRequestWithContext(a.ctx, "GET", a.URL, nil)
	if err != nil {
		ch <- err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ch <- err
	}

	var data Current
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- err
	}
}
