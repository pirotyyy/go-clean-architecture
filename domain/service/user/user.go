package user

import (
	"ca-tech/domain/model/user"
	user2 "ca-tech/domain/repository/user"
	"context"
	"errors"
)

type UserService interface {
	CreateUser(ctx context.Context, name string) (*user.User, error)
	GetUser(ctx context.Context, token string) (*user.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*user.User, error)
}

type userService struct {
	userRepo user2.UserRepository
}

func NewUserService(ur user2.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) CreateUser(ctx context.Context, name string) (*user.User, error) {
	if !isNameValid(name) {
		return nil, errors.New("name format is invalid")
	}
	return us.userRepo.CreateUser(ctx, name)
}

func (us *userService) GetUser(ctx context.Context, token string) (*user.User, error) {
	if !isTokenValid(token) {
		return nil, errors.New("token format is invalid")
	}
	return us.userRepo.GetUserByToken(ctx, token)
}

func (us *userService) UpdateUser(ctx context.Context, name string, token string) (*user.User, error) {
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
