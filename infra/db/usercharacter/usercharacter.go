package userCharacter

import (
	userCharacterModel "ca-tech/domain/model/usercharacter"
	userCharacterRepo "ca-tech/domain/repository/usercharacter"
	"context"
	"database/sql"
)

type userCharacterRepository struct {
	DB *sql.DB
}

func NewUserCharacterRepository(db *sql.DB) userCharacterRepo.UserCharacterRepository {
	return &userCharacterRepository{
		DB: db,
	}
}

func (ucr *userCharacterRepository) GetUserCharactersByUserId(ctx context.Context, userId int64) ([]*userCharacterModel.UserCharacter, error) {
	const (
		selectUserCharacterCommand = "SELECT usercharacter.id, usercharacter.character_id, game_character.name FROM usercharacter INNER JOIN game_character ON usercharacter.character_id = game_character.id WHERE usercharacter.user_id = ?"
	)

	rows, err := ucr.DB.QueryContext(ctx, selectUserCharacterCommand, userId)
	if err != nil {
		return nil, err
	}

	var userCharacters []*userCharacterModel.UserCharacter
	for rows.Next() {
		var uc userCharacterModel.UserCharacter
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
