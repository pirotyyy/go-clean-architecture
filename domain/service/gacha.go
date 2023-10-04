package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"math/rand"
	"time"
)

var (
	r             = rand.New(rand.NewSource(time.Now().UnixNano()))
	characterList []*model.Character
)

type GachaService interface {
	InitCharacterList(ctx context.Context) error
	Draw(ctx context.Context, times int64) ([]*model.GachaResult, error)
}

type gachaService struct {
	userRepo      repository.UserRepository
	charaRepo     repository.CharacterRepository
	userCharaRepo repository.UserCharacterRepository
}

func NewGachaService(ur repository.UserRepository, cr repository.CharacterRepository, ucr repository.UserCharacterRepository) GachaService {
	return &gachaService{
		userRepo:      ur,
		charaRepo:     cr,
		userCharaRepo: ucr,
	}
}

func (gs *gachaService) InitCharacterList(ctx context.Context) error {
	var err error
	characterList, err = gs.charaRepo.GetCharacters(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (gs *gachaService) Draw(ctx context.Context, times int64) ([]*model.GachaResult, error) {
	var gachaResults []*model.GachaResult

	for i := int64(0); i < times; i++ {
		rarity := r.Int63n(100)
		index := r.Int63()
		gachaResult := &model.GachaResult{Rarity: rarity, Index: index}

		gachaResults = append(gachaResults, gachaResult)
	}

	return gachaResults, nil
}
