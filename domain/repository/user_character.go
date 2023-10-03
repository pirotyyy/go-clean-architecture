package repository

import (
	"ca-tech/domain/model"
	"context"
)

type UserCharacterRepository interface {
	GetUserCharactersByUserId(ctx context.Context, userId int64) ([]*model.UserCharacter, error)
	Insert(ctx context.Context, userId int64, characterId int64) error
}
