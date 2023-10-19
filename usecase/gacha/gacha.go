package gacha

import (
	characterModel "ca-tech/domain/model/character"
	cacheCharacterService "ca-tech/domain/service/cachecharacter"
	characterService "ca-tech/domain/service/character"
	gachaService "ca-tech/domain/service/gacha"
	userService "ca-tech/domain/service/user"
	userCharacterService "ca-tech/domain/service/usercharacter"
	"context"
	"encoding/json"
)

type GachaUsecase interface {
	Draw(ctx context.Context, times int64, token string) ([]*characterModel.Character, error)
}

type gachaUsecase struct {
	userService       userService.UserService
	charaSerivce      characterService.CharacterService
	gachaService      gachaService.GachaService
	userCharaService  userCharacterService.UserCharacterService
	cacheCharaService cacheCharacterService.CacheCharacterService
}

func NewGachaUsecase(us userService.UserService, cs characterService.CharacterService, gs gachaService.GachaService, ucs userCharacterService.UserCharacterService, ccs cacheCharacterService.CacheCharacterService) GachaUsecase {
	return &gachaUsecase{
		userService:       us,
		charaSerivce:      cs,
		gachaService:      gs,
		userCharaService:  ucs,
		cacheCharaService: ccs,
	}
}

func (gu *gachaUsecase) Draw(ctx context.Context, times int64, token string) ([]*characterModel.Character, error) {
	var resultCharacters []*characterModel.Character

	// user
	user, err := gu.userService.GetUser(ctx, token)
	if err != nil {
		return nil, err
	}

	// draw gacha
	gachaResults, err := gu.gachaService.Draw(ctx, times)
	if err != nil {
		return nil, err
	}

	for _, gachaResult := range gachaResults {
		rarity := getRarity(gachaResult.Rarity)
		characterCacheData, err := gu.cacheCharaService.GetCharactersCache(ctx, string(rarity))
		if err != nil {
			return nil, err
		}

		var characters []*characterModel.Character
		err = json.Unmarshal([]byte(characterCacheData.JsonString), &characters)
		if err != nil {
			return nil, err
		}
		targetCharacter := characters[gachaResult.Index%int64(len(characters))]
		resultCharacters = append(resultCharacters, targetCharacter)
	}

	// save usercharacter
	for _, resultCharacter := range resultCharacters {
		err := gu.userCharaService.Save(ctx, user.UserId, resultCharacter.CharacterID)
		if err != nil {
			return nil, err
		}
	}

	return resultCharacters, nil
}

func getRarity(randNum int64) characterModel.Rarity {

	switch {
	case randNum < 50:
		return characterModel.N
	case randNum < 70:
		return characterModel.R
	case randNum < 90:
		return characterModel.SR
	default:
		return characterModel.SSR
	}
}
