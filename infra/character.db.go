package infra

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"database/sql"
)

type characterRepository struct {
	DB *sql.DB
}

func NewCharacterRepository(db *sql.DB) repository.CharacterRepository {
	return &characterRepository{
		DB: db,
	}
}

func (cr *characterRepository) GetUserCharactersByToken(ctx context.Context, token string) (userCharacters []*model.UserCharacter, err error) {
	const (
		selectUserCommand          = "SELECT user_id FROM user WHERE token = ?"
		selectUserCharacterCommand = "SELECT user_character.user_character_id, user_character.character_id, game_character.name FROM user_character INNER JOIN game_character ON user_character.character_id = game_character.character_id WHERE user_character.user_id = ?"
	)

	var userId string
	if err := cr.DB.QueryRowContext(ctx, selectUserCommand, token).Scan(&userId); err != nil {
		return nil, err
	}

	rows, err := cr.DB.QueryContext(ctx, selectUserCharacterCommand, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var uc model.UserCharacter
		if err := rows.Scan(&uc.UserCharacterID, &uc.CharacterID, &uc.Name); err != nil {
			return nil, err
		}
		userCharacters = append(userCharacters, &uc)
	}

	return userCharacters, nil
}
