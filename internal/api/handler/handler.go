package handler

import (
	"github.com/DanielTitkov/cm-stat-engine/internal/app"
	"github.com/DanielTitkov/cm-stat-engine/internal/configs"
	"github.com/DanielTitkov/cm-stat-engine/internal/logger"
	"github.com/labstack/echo"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {

	v1 := e.Group("/api/v1")
	// v1.POST("/getToken", h.GetTokenHandler)
	// v1.POST("/createUser", h.CreateUserHandler)

	// actuator urls
	v1Actuator := v1.Group("/actuator")
	v1Actuator.POST("/health", h.HealthHandler)
	v1Actuator.POST("/stats", h.StatsHandler)
}
