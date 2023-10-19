package usercharacter

import (
	"ca-tech/domain/model/usercharacter"
	"ca-tech/domain/repository/user"
	user_character2 "ca-tech/domain/repository/usercharacter"
	"context"
)

type UserCharacterService interface {
	GetUserCharactersByToken(ctx context.Context, token string) ([]*usercharacter.UserCharacter, error)
	Save(ctx context.Context, userId int64, characterId int64) error
}

type userCharacterService struct {
	userRepo      user.UserRepository
	userCharaRepo user_character2.UserCharacterRepository
}

func NewUserCharacterRepository(ur user.UserRepository, ucr user_character2.UserCharacterRepository) UserCharacterService {
	return &userCharacterService{
		userRepo:      ur,
		userCharaRepo: ucr,
	}
}

func (ucs *userCharacterService) GetUserCharactersByToken(ctx context.Context, token string) ([]*usercharacter.UserCharacter, error) {
	user, err := ucs.userRepo.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return ucs.userCharaRepo.GetUserCharactersByUserId(ctx, user.UserId)
}

func (ucs *userCharacterService) Save(ctx context.Context, userId int64, characterId int64) error {
	return ucs.userCharaRepo.Insert(ctx, userId, characterId)
}
