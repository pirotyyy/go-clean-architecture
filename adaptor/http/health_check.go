package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckResponse struct {
	Message string `json:"message"`
}

func healthChcek(c echo.Context) error {
	message := fmt.Sprintf("Status OK! %s", c.Path())
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{
			Message: message,
		},
	)
}
