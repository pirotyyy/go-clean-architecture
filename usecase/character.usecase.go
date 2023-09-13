package usecase

import (
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"context"
)

type CharacterUsecase interface {
	GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error)
}

type characterUsecase struct {
	svc service.CharacterService
}

func NewCharacterUsecase(cs service.CharacterService) CharacterUsecase {
	return &characterUsecase{
		svc: cs,
	}
}

func (cu *characterUsecase) GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error) {
	return cu.svc.GetUserCharactersByToken(ctx, token)
}
