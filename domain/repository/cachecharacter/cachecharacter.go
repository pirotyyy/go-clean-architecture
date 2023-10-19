package cachecharacter

import (
	"ca-tech/domain/model/cachecharacter"
	"context"
)

type CacheCharacterRepository interface {
	GetCharactersCache(ctx context.Context, key string) (cachecharacter.CacheCharactersData, error)
	SetCharactersCache(ctx context.Context, cacheCharacterData cachecharacter.CacheCharactersData) error
}
