package transport

import (
	"context"
	"encoding/json"
	"github.com/osamikoyo/meteor/internal/data/models"
	"io/ioutil"
	"net/http"
)

type ApiRouter struct {
	URL string
	ctx context.Context
	key string
}

func (a *ApiRouter) ApiScanner(ch chan error, dataCh chan models.WeatherResponses) {
	req, err := http.NewRequestWithContext(a.ctx, "GET", a.URL, nil)
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
