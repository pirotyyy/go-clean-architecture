package user

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

func (ur *userRepository) CreateUser(ctx context.Context, name string) (*model.User, error) {
	const (
		insert  = "INSERT INTO user (name, token) VALUES (?, ?)"
		confirm = "SELECT name, created_at FROM user WHERE id = ?"
	)

	token := tokenGenerator()
	res, err := ur.DB.ExecContext(ctx, insert, name, token)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		UserId: id,
		Token:  token,
	}
	if err := ur.DB.QueryRowContext(ctx, confirm, newUser.UserId).Scan(&newUser.Name, &newUser.CreatedAt); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (ur *userRepository) GetUserByToken(ctx context.Context, token string) (*model.User, error) {
	const (
		selectCommand = "SELECT id, name, created_at FROM user WHERE token = ?"
	)

	var user = &model.User{}
	if err := ur.DB.QueryRowContext(ctx, selectCommand, token).Scan(&user.UserId, &user.Name, &user.CreatedAt); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, name string, token string) (*model.User, error) {
	const (
		update  = "UPDATE user SET name = ? WHERE token = ?"
		confirm = "SELECT id, name, created_at FROM user WHERE token = ?"
	)

	_, err := ur.DB.ExecContext(ctx, update, name, token)
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
