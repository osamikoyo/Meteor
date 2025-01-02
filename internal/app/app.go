package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-co-op/gocron"
	"github.com/osamikoyo/meteor/internal/cities"
	"github.com/osamikoyo/meteor/internal/data/models"
	"github.com/osamikoyo/meteor/internal/handler"
	"github.com/osamikoyo/meteor/internal/keys"
	"github.com/osamikoyo/meteor/internal/service"
	"github.com/osamikoyo/meteor/internal/transport"
	"github.com/osamikoyo/meteor/pkg/loger"
	"net/http"
	"sync"
	"time"
)

type App struct {
	timer     *gocron.Scheduler
	logger    loger.Logger
	service   service.Service
	apiRouter transport.ApiRouter
	handler   handler.Handler
}

func Init() *App {
	s := gocron.NewScheduler(time.Local)
	serv := service.New()
	logger := loger.New()

	return &App{
		timer:     s,
		service:   serv,
		logger:    logger,
		apiRouter: transport.New(),
		handler: handler.Handler{
			ST: serv,
		},
	}
}

func (a *App) route() {
	var ch chan error
	var data chan models.WeatherResponses
	var wg *sync.WaitGroup
	for _, c := range cities.Cities {
		wg.Add(1)
		go func() {
			a.apiRouter.URL = fmt.Sprintf("http://api.weatherapi.com/v1/current.json\\?key\\=\\%s\\&q\\=%s", keys.WEATHERAPI, c)
			a.apiRouter.ApiScanner(ch, data)
			select {
			case err := <-ch:
				a.logger.Error().Err(err)
				wg.Done()
			case d := <-data:
				if err := a.service.Add(d); err != nil {
					a.logger.Error().Err(err)
				}
			}
		}()
	}

}

func (a *App) Run() {
	a.timer.Every(1).Day().Do(a.route)
	a.timer.StartAsync()

	r := chi.NewRouter()

	a.handler.RegisterRoutes(r)

	server := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}

	a.logger.Info().Msg("Started http server on localhost:8080 :3")
	if err := server.ListenAndServe(); err != nil {
		a.logger.Error().Err(err)
	}
}
