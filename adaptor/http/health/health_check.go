package health

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler interface {
	HealthCheck() echo.HandlerFunc
}

type healthCheckHandler struct {
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

type healthCheckResponse struct {
	Message string `json:"message"`
}

func (hch *healthCheckHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {

		message := fmt.Sprintf("Status OK! %s", c.Path())
		return c.JSON(
			http.StatusOK,
			healthCheckResponse{
				Message: message,
			},
		)
	}
}
