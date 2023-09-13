package http

import (
	"ca-tech/domain/model"
	"ca-tech/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type characterHandler struct {
	usecase usecase.CharacterUsecase
}

func NewCharacterHandler(cu usecase.CharacterUsecase) *characterHandler {
	return &characterHandler{
		usecase: cu,
	}
}

func (ch *characterHandler) GetUserCharactersByToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token := c.Request().Header.Get("x-token")
		if token == "" {
			return c.JSON(http.StatusBadRequest, &model.ErrResponse{
				Message: "token is required"})
		}

		userCharacters, err := ch.usecase.GetUserCharactersByToken(ctx, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &model.ErrResponse{
				Message: err.Error(),
			})
		}

		characters := []*model.UserCharacter{}
		userCharacters = append(characters, userCharacters...)

		res := &model.CharacterListResponse{
			Characters: userCharacters,
		}

		return c.JSON(http.StatusOK, res)
	}
}
