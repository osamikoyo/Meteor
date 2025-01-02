package transport

import (
	"context"
	"encoding/json"
	"github.com/osamikoyo/meteor/internal/data/models"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiRouter struct {
	URL string
	Ctx context.Context
	key string
}

func New() ApiRouter {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()
	return ApiRouter{Ctx: ctx}
}

func (a *ApiRouter) ApiScanner(ch chan error, dataCh chan models.WeatherResponses) {
	req, err := http.NewRequestWithContext(a.Ctx, "GET", a.URL, nil)
	if err != nil {
		ch <- err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ch <- err
	}

	var data models.WeatherResponses
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- err
	}

	ch <- nil
	dataCh <- data
}
