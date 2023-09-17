package userCharacter

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"database/sql"
)

type userCharacterRepository struct {
	DB *sql.DB
}

func NewUserCharacterRepository(db *sql.DB) repository.UserCharacterRepository {
	return &userCharacterRepository{
		DB: db,
	}
}

func (ucr *userCharacterRepository) GetUserCharactersByUserId(ctx context.Context, userId int64) ([]*model.UserCharacter, error) {
	const (
		selectUserCharacterCommand = "SELECT user_character.id, user_character.character_id, game_character.name FROM user_character INNER JOIN game_character ON user_character.character_id = game_character.id WHERE user_character.user_id = ?"
	)

	rows, err := ucr.DB.QueryContext(ctx, selectUserCharacterCommand, userId)
	if err != nil {
		return nil, err
	}

	var userCharacters []*model.UserCharacter
	for rows.Next() {
		var uc model.UserCharacter
		if err := rows.Scan(&uc.UserCharacterID, &uc.CharacterID, &uc.Name); err != nil {
			return nil, err
		}
		userCharacters = append(userCharacters, &uc)
	}

	return userCharacters, nil
}

func (ucr *userCharacterRepository) Insert(ctx context.Context, userId int64, characterId int64) error {
	const insertCommand = "INSERT INTO user_character (user_id, character_id) VALUES (?, ?)"

	_, err := ucr.DB.ExecContext(ctx, insertCommand, userId, characterId)
	if err != nil {
		return err
	}

	return nil
}
