package app

import (
	"github.com/DanielTitkov/cm-stat-engine/internal/configs"
	"github.com/DanielTitkov/cm-stat-engine/internal/logger"
	"github.com/robfig/cron"
)

type (
	// App combines services and holds business logic
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		Cron   *cron.Cron
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
) (*App, error) {
	c := cron.New()
	c.Start()

	app := App{
		cfg:    cfg,
		logger: logger,
		Cron:   c,
	}

	return &app, nil
}
