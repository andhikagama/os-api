package shared

import (
	"fmt"
	"os"

	"github.com/andhikagama/os-api/config"
	"github.com/andhikagama/os-api/shared/resource"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// Resource .
type Resource struct {
	DB        *gorm.DB
	Config    *config.EnvConfig
	Echo      *echo.Echo
	Logger    *logrus.Logger
	Validator *resource.Validator
}

// New ...
func New(config *config.EnvConfig) Resource {
	r := Resource{
		Echo:      resource.NewEcho(),
		Validator: resource.NewValidator(),
		Config:    config,
		Logger:    resource.NewLogger(),
	}

	db, err := resource.NewDB(config)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("mysql connection failed. Err: %v", err.Error()))
		os.Exit(1)
	}

	r.DB = db

	return r
}
