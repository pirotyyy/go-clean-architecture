package repository

import (
	"ca-tech/domain/model"
	"context"
)

// 抽象化
type UserRepository interface {
	// name でなくmodelでやりとりしたい
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByToken(ctx context.Context, token string) (*model.User, error)
	Update(ctx context.Context, user *model.User, token string) (*model.User, error)
}
