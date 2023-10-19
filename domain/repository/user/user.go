package user

import (
	"ca-tech/domain/model/user"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name string) (*user.User, error)
	GetUserByToken(ctx context.Context, token string) (*user.User, error)
	UpdateUser(ctx context.Context, name string, token string) (*user.User, error)
}
