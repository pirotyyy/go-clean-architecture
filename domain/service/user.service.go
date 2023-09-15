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
	repo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return us.repo.CreateUser(ctx, user)
}

func (us *userService) GetUser(ctx context.Context, token string) (*model.User, error) {
	return us.repo.GetUserByToken(ctx, token)
}

func (us *userService) UpdateUser(ctx context.Context, user *model.User, token string) (*model.User, error) {
	return us.repo.UpdateUser(ctx, user, token)
}
