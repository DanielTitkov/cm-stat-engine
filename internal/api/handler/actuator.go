package handler

import (
	"net/http"

	"github.com/DanielTitkov/cm-stat-engine/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "service is running",
	})
}

func (h *Handler) StatsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, model.ServiceStatsResponse{})
}
