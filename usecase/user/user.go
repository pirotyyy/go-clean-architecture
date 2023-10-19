package user

import (
	userModel "ca-tech/domain/model/user"
	userService "ca-tech/domain/service/user"
	"context"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, name string) (*userModel.User, error)
	GetUser(ctx context.Context, token string) (*userModel.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*userModel.User, error)
}

type userUsecase struct {
	svc userService.UserService
}

func NewUserUsecase(us userService.UserService) UserUsecase {
	return &userUsecase{
		svc: us,
	}
}

func (uu *userUsecase) CreateUser(ctx context.Context, name string) (*userModel.User, error) {
	return uu.svc.CreateUser(ctx, name)
}

func (uu *userUsecase) GetUser(ctx context.Context, token string) (*userModel.User, error) {
	return uu.svc.GetUser(ctx, token)
}

func (uu *userUsecase) UpdateUser(ctx context.Context, name string, token string) (*userModel.User, error) {
	return uu.svc.UpdateUser(ctx, name, token)
}
