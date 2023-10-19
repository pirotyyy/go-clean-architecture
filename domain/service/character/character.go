package character

import (
	characterModel "ca-tech/domain/model/character"
	characterRepo "ca-tech/domain/repository/character"
	"context"
)

type CharacterService interface {
	GetCharacters(ctx context.Context) ([]*characterModel.Character, error)
}

type characterService struct {
	charaRepo characterRepo.CharacterRepository
}

func NewCharacterService(cr characterRepo.CharacterRepository) CharacterService {
	return &characterService{
		charaRepo: cr,
	}
}

func (cs *characterService) GetCharacters(ctx context.Context) ([]*characterModel.Character, error) {
	characters, err := cs.charaRepo.GetCharacters(ctx)
	if err != nil {
		return nil, err
	}
	return characters, nil
}
