package cachecharacter

type CacheCharactersData struct {
	Key        string
	JsonString string
}

type GetCacheCharacterDataResponse struct {
	CacheCharacters CacheCharactersData
}
