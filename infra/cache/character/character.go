package character

import (
	cachecharacterModel "ca-tech/domain/model/cachecharacter"
	cachecharacterRepo "ca-tech/domain/repository/cachecharacter"
	"context"
	"github.com/redis/go-redis/v9"
)

type CacheCharacterRepository struct {
	Cache *redis.Client
}

func NewCacheCharacterRepository(cache *redis.Client) cachecharacterRepo.CacheCharacterRepository {
	return &CacheCharacterRepository{
		Cache: cache,
	}
}

func (c *CacheCharacterRepository) GetCharactersCache(ctx context.Context, key string) (cachecharacterModel.CacheCharactersData, error) {
	cacheData, err := c.Cache.Get(ctx, key).Result()

	if err != nil {
		return cachecharacterModel.CacheCharactersData{}, err
	}

	cacheCharacterData := cachecharacterModel.CacheCharactersData{
		Key:        key,
		JsonString: cacheData,
	}

	return cacheCharacterData, nil
}

func (c *CacheCharacterRepository) SetCharactersCache(ctx context.Context, cacheCharacterData cachecharacterModel.CacheCharactersData) error {
	err := c.Cache.Set(ctx, cacheCharacterData.Key, cacheCharacterData.JsonString, 0).Err()

	if err != nil {
		return err
	}
	return nil
}
