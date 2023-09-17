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
	Draw(ctx context.Context, times int64, token string) ([]*model.Character, error)
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

func (gs *gachaService) Draw(ctx context.Context, times int64, token string) ([]*model.Character, error) {
	var characters []*model.Character

	user, err := gs.userRepo.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < times; i++ {
		rarity := selectRarity()
		targetCharacters := getTargetCharacters(rarity)
		character := targetCharacters[r.Intn(len(targetCharacters))]

		characters = append(characters, character)
	}

	for _, character := range characters {
		err := gs.userCharaRepo.Insert(ctx, user.UserId, character.CharacterID)
		if err != nil {
			return nil, err
		}
	}

	return characters, nil
}

func selectRarity() string {
	num := r.Intn(100)

	switch {
	case num < 50:
		return "N"
	case num < 70:
		return "R"
	case num < 90:
		return "SR"
	default:
		return "SSR"
	}
}

func getTargetCharacters(rarity string) []*model.Character {
	var targetCharacters []*model.Character

	for _, char := range characterList {
		if char.Rarity == rarity {
			targetCharacters = append(targetCharacters, char)
		}
	}

	return targetCharacters
}
