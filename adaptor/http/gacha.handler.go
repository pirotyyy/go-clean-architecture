package http

import (
	"ca-tech/domain/model"
	"ca-tech/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type gachaHandler struct {
	usecase usecase.GachaUsecase
}

func NewGachaHandler(gu usecase.GachaUsecase) *gachaHandler {
	return &gachaHandler{
		usecase: gu,
	}
}

func (gh *gachaHandler) Draw() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &model.GachaDrawRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		token := c.Request().Header.Get("x-token")
		if token == "" {
			return c.JSON(http.StatusBadRequest, &model.ErrResponse{
				Message: "token is required",
			})
		}

		characters, err := gh.usecase.Draw(ctx, req.Times, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &model.ErrResponse{
				Message: err.Error(),
			})
		}

		results := []model.GachaResult{}
		for _, character := range characters {
			results = append(results, model.GachaResult{
				CharacterID: character.CharacterID,
				Name:        character.Name,
				Rarity:      character.Rarity,
			})
		}

		res := &model.GachaDrawResponse{
			Results: results,
		}

		return c.JSON(http.StatusOK, res)
	}
}
