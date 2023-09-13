package repository

import (
	"ca-tech/domain/model"
	"context"
)

type CharacterRepository interface {
	GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error)
}
