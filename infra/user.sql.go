package infra

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	const (
		insert  = "INSERT INTO user (name, token) VALUES (?, ?)"
		confirm = "SELECT name, created_at FROM user WHERE user_id = ?"
	)

	res, err := ur.DB.ExecContext(ctx, insert, user.Name, tokenGenerator())
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		UserId: id,
	}
	if err := ur.DB.QueryRowContext(ctx, confirm, newUser.UserId).Scan(&user.Name, &user.CreatedAt); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (ur *userRepository) GetByToken(ctx context.Context, token string) (*model.User, error) {
	const (
		selectCommand = "SELECT user_id, name, created_at FROM user WHERE token = ?"
	)

	var user = &model.User{}
	if err := ur.DB.QueryRowContext(ctx, selectCommand, token).Scan(&user.UserId, &user.Name, &user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) Update(ctx context.Context, user *model.User, token string) (*model.User, error) {
	const (
		update  = "UPDATE user SET name = ? WHERE token = ?"
		confirm = "SELECT user_id, name, created_at FROM user WHERE token = ?"
	)

	_, err := ur.DB.ExecContext(ctx, update, user.Name, token)
	if err != nil {
		return nil, err
	}

	var updatedUser = &model.User{}
	if err := ur.DB.QueryRowContext(ctx, confirm, token).Scan(&updatedUser.UserId, &updatedUser.Name, &updatedUser.CreatedAt); err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
