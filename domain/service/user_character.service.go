package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type UserCharacterService interface {
	GetUserCharactersByToken(ctx context.Context, token string) ([]*model.UserCharacter, error)
}

type userCharacterService struct {
	userRepo      repository.UserRepository
	userCharaRepo repository.UserCharacterRepository
}

func NewUserCharacterRepository(ur repository.UserRepository, ucr repository.UserCharacterRepository) UserCharacterService {
	return &userCharacterService{
		userRepo:      ur,
		userCharaRepo: ucr,
	}
}

func (ucs *userCharacterService) GetUserCharactersByToken(ctx context.Context, token string) ([]*model.UserCharacter, error) {
	user, err := ucs.userRepo.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return ucs.userCharaRepo.GetUserCharactersByUserId(ctx, user.UserId)
}
