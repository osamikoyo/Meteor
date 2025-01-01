package app

import (
	"github.com/osamikoyo/meteor/internal/service"
	"github.com/osamikoyo/meteor/pkg/loger"
)

type App struct {
	logger  loger.Logger
	service service.Service
}
