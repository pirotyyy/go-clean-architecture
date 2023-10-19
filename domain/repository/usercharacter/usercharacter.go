package usercharacter

import (
	"ca-tech/domain/model/usercharacter"
	"context"
)

type UserCharacterRepository interface {
	GetUserCharactersByUserId(ctx context.Context, userId int64) ([]*usercharacter.UserCharacter, error)
	Insert(ctx context.Context, userId int64, characterId int64) error
}
