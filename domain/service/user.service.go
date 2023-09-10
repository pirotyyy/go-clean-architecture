package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUser(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return us.repo.Insert(ctx, name)
}

func (us *userService) GetUser(ctx context.Context, token string) (*model.User, error) {
	return us.repo.Select(ctx, token)
}

func (us *userService) UpdateUser(ctx context.Context, name string, token string) (*model.User, error) {
	return us.repo.Update(ctx, name, token)
}
