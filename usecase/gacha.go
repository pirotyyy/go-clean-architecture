package usecase

import (
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"context"
)

type GachaUsecase interface {
	Draw(ctx context.Context, times int64, token string) ([]*model.Character, error)
}

type gachaUsecase struct {
	userService      service.UserService
	charaSerivce     service.CharacterService
	gachaService     service.GachaService
	userCharaService service.UserCharacterService
}

func NewGachaUsecase(us service.UserService, cs service.CharacterService, gs service.GachaService, ucs service.UserCharacterService) GachaUsecase {
	return &gachaUsecase{
		userService:      us,
		charaSerivce:     cs,
		gachaService:     gs,
		userCharaService: ucs,
	}
}

func (gu *gachaUsecase) Draw(ctx context.Context, times int64, token string) ([]*model.Character, error) {
	var resultCharacters []*model.Character

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

	characters, err := gu.charaSerivce.GetCharacters(ctx)
	if err != nil {
		return nil, err
	}

	for _, gachaResult := range gachaResults {
		rarity := getRarity(gachaResult.Rarity)
		targetCharacters := getTargetCharacters(characters, rarity)
		targetCharacter := targetCharacters[gachaResult.Index%int64(len(targetCharacters))]
		resultCharacters = append(resultCharacters, targetCharacter)
	}

	// save user_character
	for _, resultCharacter := range resultCharacters {
		err := gu.userCharaService.Insert(ctx, user.UserId, resultCharacter.CharacterID)
		if err != nil {
			return nil, err
		}
	}

	return resultCharacters, nil
}

func getTargetCharacters(characters []*model.Character, rarity model.Rarity) []*model.Character {
	var targetCharacters []*model.Character

	for _, char := range characters {
		if char.Rarity == rarity {
			targetCharacters = append(targetCharacters, char)
		}
	}

	return targetCharacters
}

func getRarity(randNum int64) model.Rarity {

	switch {
	case randNum < 50:
		return model.N
	case randNum < 70:
		return model.R
	case randNum < 90:
		return model.SR
	default:
		return model.SSR
	}
}
