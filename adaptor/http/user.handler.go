package http

import (
	"ca-tech/domain/model"
	"ca-tech/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &model.UserCreateRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		user, err := uh.usecase.CreateUser(ctx, req.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create user: %v", err))
		}

		return c.JSON(http.StatusOK, user.Name)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		user, err := uh.usecase.GetUser(ctx, c.Request().Header.Get("x-token"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to get user: %v", err))
		}

		return c.JSON(http.StatusOK, user.Name)
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &model.UserUpdateRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		user, err := uh.usecase.UpdateUser(ctx, req.Name, c.Request().Header.Get("x-token"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %v", err))
		}

		return c.JSON(http.StatusOK, user.Name)
	}
}
