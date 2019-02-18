package rest

import (
	"github.com/caioever/monografia/backend/app/health"
	"github.com/labstack/echo"
	"net/http"
)

func (restBase *restBase) Health(c echo.Context) error {
	status := health.HealthCheck(restBase.Context)
	return c.JSON(http.StatusOK, status)
}