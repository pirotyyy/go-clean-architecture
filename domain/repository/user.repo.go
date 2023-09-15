package repository

import (
	"ca-tech/domain/model"
	"context"
)

// 抽象化
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByToken(ctx context.Context, token string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, token string) (*model.User, error)
}
