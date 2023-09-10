package repository

import (
	"ca-tech/domain/model"
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, name string) (*model.User, error)
	Select(ctx context.Context, token string) (*model.User, error)
	Update(ctx context.Context, name string, token string) (*model.User, error)
}
