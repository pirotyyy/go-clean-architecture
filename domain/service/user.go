package service

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"errors"
)

type UserService interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUser(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	if !isNameValid(name) {
		return nil, errors.New("name format is invalid")
	}
	return us.userRepo.CreateUser(ctx, name)
}

func (us *userService) GetUser(ctx context.Context, token string) (*model.User, error) {
	if !isTokenValid(token) {
		return nil, errors.New("token format is invalid")
	}
	return us.userRepo.GetUserByToken(ctx, token)
}

func (us *userService) UpdateUser(ctx context.Context, name string, token string) (*model.User, error) {
	if !isNameValid(name) {
		return nil, errors.New("name format is invalid")
	}
	if !isTokenValid(token) {
		return nil, errors.New("token format is invalid")
	}
	return us.userRepo.UpdateUser(ctx, name, token)
}

func isNameValid(name string) bool {
	if name == "" || (len(name) > 20) {
		return false
	}
	return true
}

func isTokenValid(token string) bool {
	if token == "" {
		return false
	}
	return true
}
