package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, token string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return us.userRepo.CreateUser(ctx, user)
}

func (us *userService) GetUser(ctx context.Context, token string) (*model.User, error) {
	return us.userRepo.GetUserByToken(ctx, token)
}

func (us *userService) UpdateUser(ctx context.Context, user *model.User, token string) (*model.User, error) {
	return us.userRepo.UpdateUser(ctx, user, token)
}
