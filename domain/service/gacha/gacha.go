package gacha

import (
	"ca-tech/domain/model/gacha"
	"ca-tech/domain/repository/character"
	"ca-tech/domain/repository/user"
	"ca-tech/domain/repository/usercharacter"
	"context"
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type GachaService interface {
	Draw(ctx context.Context, times int64) ([]*gacha.GachaResult, error)
}

type gachaService struct {
	userRepo      user.UserRepository
	charaRepo     character.CharacterRepository
	userCharaRepo usercharacter.UserCharacterRepository
}

func NewGachaService(ur user.UserRepository, cr character.CharacterRepository, ucr usercharacter.UserCharacterRepository) GachaService {
	return &gachaService{
		userRepo:      ur,
		charaRepo:     cr,
		userCharaRepo: ucr,
	}
}

func (gs *gachaService) Draw(ctx context.Context, times int64) ([]*gacha.GachaResult, error) {
	var gachaResults []*gacha.GachaResult

	for i := int64(0); i < times; i++ {
		rarity := r.Int63n(100)
		index := r.Int63()
		gachaResult := &gacha.GachaResult{Rarity: rarity, Index: index}

		gachaResults = append(gachaResults, gachaResult)
	}

	return gachaResults, nil
}
