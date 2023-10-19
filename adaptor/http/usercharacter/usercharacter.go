package usercharacter

import (
	characterModel "ca-tech/domain/model/character"
	errModel "ca-tech/domain/model/error"
	userCharacterModel "ca-tech/domain/model/usercharacter"
	userCharacterUsecase "ca-tech/usecase/usercharacter"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userCharacterHandler struct {
	usecase userCharacterUsecase.UserCharacterUsecase
}

func NewUserCharacterHandler(ucu userCharacterUsecase.UserCharacterUsecase) *userCharacterHandler {
	return &userCharacterHandler{
		usecase: ucu,
	}
}

func (uch *userCharacterHandler) GetUserCharactersByToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		token := c.Request().Header.Get("x-token")
		if token == "" {
			return c.JSON(http.StatusBadRequest, &errModel.ErrResponse{
				Message: "token is required"})
		}

		userCharacters, err := uch.usecase.GetUserCharactersByToken(ctx, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &errModel.ErrResponse{
				Message: err.Error(),
			})
		}

		var characters []*userCharacterModel.UserCharacter
		userCharacters = append(characters, userCharacters...)

		res := &characterModel.CharacterListResponse{
			Characters: userCharacters,
		}

		return c.JSON(http.StatusOK, res)
	}
}
