package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/osamikoyo/meteor/internal/data/models"
	"github.com/osamikoyo/meteor/internal/keys"
	"io/ioutil"
	"net/http"
)

type ApiRouter struct {
	URL string
	ctx context.Context
	key string
}

func New(city string) *ApiRouter {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json\\?key\\=\\%s\\&q\\=%s", keys.WEATHERAPI)

	return &ApiRouter{URL: url, key: keys.WEATHERAPI}
}

func (a *ApiRouter) ApiScanner(ch chan error, dataCh chan models.Current) {
	req, err := http.NewRequestWithContext(a.ctx, "GET", a.URL, nil)
	if err != nil {
		ch <- err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ch <- err
	}

	var data models.Current
	if err := json.Unmarshal(body, &data); err != nil {
		ch <- err
	}

	ch <- nil
	dataCh <- data
}
