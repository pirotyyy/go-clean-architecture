package repository

import (
	"ca-tech/domain/model"
	"context"
)

type CharacterRepository interface {
	GetCharacters(ctx context.Context) ([]*model.Character, error)
}
