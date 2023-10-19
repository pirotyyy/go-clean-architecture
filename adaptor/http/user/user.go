package user

import (
	errModel "ca-tech/domain/model/error"
	userModel "ca-tech/domain/model/user"
	userUsecase "ca-tech/usecase/user"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase userUsecase.UserUsecase
}

func NewUserHandler(uu userUsecase.UserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &userModel.UserCreateRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		if req.Name == "" {
			errRes := &errModel.ErrResponse{
				Message: "name is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		newUser, err := uh.usecase.CreateUser(ctx, req.Name)
		if err != nil {
			errRes := &errModel.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &userModel.UserCreateResponse{
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
			errRes := &errModel.ErrResponse{
				Message: "token is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		user, err := uh.usecase.GetUser(ctx, token)
		if err != nil {
			errRes := &errModel.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &userModel.UserGetResponse{
			Name: user.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &userModel.UserUpdateRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		token := c.Request().Header.Get("x-token")
		if token == "" {
			errRes := &errModel.ErrResponse{
				Message: "token is empty",
			}
			return c.JSON(http.StatusBadRequest, errRes)
		}

		updatedUser, err := uh.usecase.UpdateUser(ctx, req.Name, token)
		if err != nil {
			errRes := &errModel.ErrResponse{
				Message: err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, errRes)
		}

		res := &userModel.UserUpdateResponse{
			Name: updatedUser.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}
