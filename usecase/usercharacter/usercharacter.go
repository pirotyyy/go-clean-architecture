package usercharacter

import (
	userCharacterModel "ca-tech/domain/model/usercharacter"
	userCharacterService "ca-tech/domain/service/usercharacter"
	"context"
)

type UserCharacterUsecase interface {
	GetUserCharactersByToken(ctx context.Context, token string) ([]*userCharacterModel.UserCharacter, error)
}

type userCharacterUsecase struct {
	svc userCharacterService.UserCharacterService
}

func NewUserCharacterUsecase(ucs userCharacterService.UserCharacterService) UserCharacterUsecase {
	return &userCharacterUsecase{
		svc: ucs,
	}
}

func (ucu *userCharacterUsecase) GetUserCharactersByToken(ctx context.Context, token string) ([]*userCharacterModel.UserCharacter, error) {
	return ucu.svc.GetUserCharactersByToken(ctx, token)
}
