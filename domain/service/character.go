package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type CharacterService interface {
	GetCharacters(ctx context.Context) ([]*model.Character, error)
}

type characterService struct {
	charaRepo repository.CharacterRepository
}

func NewCharacterService(cr repository.CharacterRepository) CharacterService {
	return &characterService{
		charaRepo: cr,
	}
}

func (cs *characterService) GetCharacters(ctx context.Context) ([]*model.Character, error) {
	characters, err := cs.charaRepo.GetCharacters(ctx)
	if err != nil {
		return nil, err
	}
	return characters, nil
}
