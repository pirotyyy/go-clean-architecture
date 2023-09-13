package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type CharacterService interface {
	GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error)
}

type characterService struct {
	repo repository.CharacterRepository
}

func NewCharacterService(cr repository.CharacterRepository) CharacterService {
	return &characterService{
		repo: cr,
	}
}

func (cs *characterService) GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error) {
	return cs.repo.GetUserCharactersByToken(ctx, token)
}
