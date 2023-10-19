package cachecharacter

import (
	"ca-tech/domain/model/cachecharacter"
	cachecharacterRepo "ca-tech/domain/repository/cachecharacter"
	"context"
)

type CacheCharacterService interface {
	GetCharactersCache(ctx context.Context, key string) (cachecharacter.CacheCharactersData, error)
	SetCharactersCache(ctx context.Context, cacheCharacterData cachecharacter.CacheCharactersData) error
}

type cacheCharacterService struct {
	cacheCharacterRepo cachecharacterRepo.CacheCharacterRepository
}

func NewCacheCharacterService(ccr cachecharacterRepo.CacheCharacterRepository) CacheCharacterService {
	return &cacheCharacterService{
		cacheCharacterRepo: ccr,
	}
}

func (ccs *cacheCharacterService) GetCharactersCache(ctx context.Context, key string) (cachecharacter.CacheCharactersData, error) {
	cacheCharacterData, err := ccs.cacheCharacterRepo.GetCharactersCache(ctx, key)
	if err != nil {
		return cachecharacter.CacheCharactersData{}, err
	}
	return cacheCharacterData, nil
}

func (ccs *cacheCharacterService) SetCharactersCache(ctx context.Context, cacheCharacterData cachecharacter.CacheCharactersData) error {
	err := ccs.cacheCharacterRepo.SetCharactersCache(ctx, cacheCharacterData)
	if err != nil {
		return err
	}
	return nil
}
