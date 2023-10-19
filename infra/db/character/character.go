package character

import (
	characterModel "ca-tech/domain/model/character"
	characterRepo "ca-tech/domain/repository/character"
	"context"
	"database/sql"
)

/*
	character-tableに対する操作を定義する
*/

type characterRepository struct {
	DB *sql.DB
}

func NewCharacterRepository(db *sql.DB) characterRepo.CharacterRepository {
	return &characterRepository{
		DB: db,
	}
}

func (cr *characterRepository) GetCharacters(ctx context.Context) ([]*characterModel.Character, error) {
	const selectCommand = "SELECT id, name, rarity FROM game_character"

	rows, err := cr.DB.QueryContext(ctx, selectCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*characterModel.Character
	for rows.Next() {
		var c characterModel.Character
		if err := rows.Scan(&c.CharacterID, &c.Name, &c.Rarity); err != nil {
			return nil, err
		}
		characters = append(characters, &c)
	}

	return characters, nil
}
