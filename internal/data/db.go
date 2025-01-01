package data

import (
	"github.com/osamikoyo/meteor/internal/data/models"
	"github.com/osamikoyo/meteor/pkg/loger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Storage interface {
	Save(weather models.Day) error
	Get(date time.Time) (models.Day, error)
	GetByRange(date1 time.Time, date2 time.Time) (models.Period, error)
}

type Database struct {
	*gorm.DB
}

func New() Storage {
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		loger.New().Error().Err(err)
	}
	if err = db.AutoMigrate(&models.Day{}); err != nil {
		loger.New().Error().Err(err)
	}

	return Database{db}
}
