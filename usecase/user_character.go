package usecase

import (
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"context"
)

type UserCharacterUsecase interface {
	GetUserCharactersByToken(ctx context.Context, token string) ([]*model.UserCharacter, error)
}

type userCharacterUsecase struct {
	svc service.UserCharacterService
}

func NewUserCharacterUsecase(ucs service.UserCharacterService) UserCharacterUsecase {
	return &userCharacterUsecase{
		svc: ucs,
	}
}

func (ucu *userCharacterUsecase) GetUserCharactersByToken(ctx context.Context, token string) ([]*model.UserCharacter, error) {
	return ucu.svc.GetUserCharactersByToken(ctx, token)
}
