package gacha

import (
	characterModel "ca-tech/domain/model/character"
	errorModel "ca-tech/domain/model/error"
	gachaUsecase "ca-tech/usecase/gacha"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type gachaHandler struct {
	usecase gachaUsecase.GachaUsecase
}

type DrawRequest struct {
	Times int64 `json:"times"`
}

type DrawResult struct {
	CharacterID int64                 `json:"character_id"`
	Name        string                `json:"name"`
	Rarity      characterModel.Rarity `json:"rarity"`
}

type DrawResponse struct {
	Results []DrawResult `json:"results"`
}

func NewGachaHandler(gu gachaUsecase.GachaUsecase) *gachaHandler {
	return &gachaHandler{
		usecase: gu,
	}
}

func (gh *gachaHandler) Draw() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := &DrawRequest{}
		dec := json.NewDecoder(c.Request().Body)
		if err := dec.Decode(&req); err != nil {
			log.Println(err)
		}

		token := c.Request().Header.Get("x-token")
		if token == "" {
			return c.JSON(http.StatusBadRequest, &errorModel.ErrResponse{
				Message: "token is required",
			})
		}

		characters, err := gh.usecase.Draw(ctx, req.Times, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &errorModel.ErrResponse{
				Message: err.Error(),
			})
		}

		var results []DrawResult
		for _, character := range characters {
			results = append(results, DrawResult{
				CharacterID: character.CharacterID,
				Name:        character.Name,
				Rarity:      character.Rarity,
			})
		}

		res := &DrawResponse{
			Results: results,
		}

		return c.JSON(http.StatusOK, res)
	}
}
