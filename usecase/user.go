package usecase

import (
	"ca-tech/domain/model"
	"ca-tech/domain/service"
	"context"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUser(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*model.User, error)
}

type userUsecase struct {
	svc service.UserService
}

func NewUserUsecase(us service.UserService) UserUsecase {
	return &userUsecase{
		svc: us,
	}
}

func (uu *userUsecase) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return uu.svc.CreateUser(ctx, name)
}

func (uu *userUsecase) GetUser(ctx context.Context, token string) (*model.User, error) {
	return uu.svc.GetUser(ctx, token)
}

func (uu *userUsecase) UpdateUser(ctx context.Context, name string, token string) (*model.User, error) {
	return uu.svc.UpdateUser(ctx, name, token)
}
