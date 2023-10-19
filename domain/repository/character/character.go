package character

import (
	"ca-tech/domain/model/character"
	"context"
)

type CharacterRepository interface {
	GetCharacters(ctx context.Context) ([]*character.Character, error)
}
