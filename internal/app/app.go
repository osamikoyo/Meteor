package app

import (
	"github.com/go-co-op/gocron"
	"github.com/osamikoyo/meteor/internal/service"
	"github.com/osamikoyo/meteor/pkg/loger"
	"time"
)

type App struct {
	timer   *gocron.Scheduler
	logger  loger.Logger
	service service.Service
}

func Init() *App {
	s := gocron.NewScheduler(time.Local)
	serv := service.New()
	logger := loger.New()
	return &App{
		timer:   s,
		service: serv,
		logger:  logger,
	}
}
