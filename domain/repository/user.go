package repository

import (
	"ca-tech/domain/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*model.User, error)
}
