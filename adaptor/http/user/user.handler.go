package user

import (
	"ca-tech/domain/model"
	"ca-tech/usecase"
	"encoding/json"
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

		if req.Name == "" {
			errRes := &model.ErrResponse{
				Message: "name is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		user := &model.User{
			Name: req.Name,
		}
		newUser, err := uh.usecase.CreateUser(ctx, user)
		if err != nil {
			errRes := &model.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &model.UserCreateResponse{
			Token: newUser.Token,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token := c.Request().Header.Get("x-token")
		if token == "" {
			errRes := &model.ErrResponse{
				Message: "token is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		user, err := uh.usecase.GetUser(ctx, token)
		if err != nil {
			errRes := &model.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &model.UserGetResponse{
			Name: user.Name,
		}

		return c.JSON(http.StatusOK, res)
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

		token := c.Request().Header.Get("x-token")
		if token == "" {
			errRes := &model.ErrResponse{
				Message: "token is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		user := &model.User{
			Name: req.Name,
		}
		updatedUser, err := uh.usecase.UpdateUser(ctx, user, token)
		if err != nil {
			errRes := &model.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &model.UserUpdateResponse{
			Name: updatedUser.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}
