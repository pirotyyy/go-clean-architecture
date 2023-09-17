package character

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"database/sql"
)

/*
	character-tableに対する操作を定義する
*/

type characterRepository struct {
	DB *sql.DB
}

func NewCharacterRepository(db *sql.DB) repository.CharacterRepository {
	return &characterRepository{
		DB: db,
	}
}

func (cr *characterRepository) GetCharacters(ctx context.Context) ([]*model.Character, error) {
	const selectCommand = "SELECT id, name, rarity FROM game_character"

	rows, err := cr.DB.QueryContext(ctx, selectCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*model.Character
	for rows.Next() {
		var c model.Character
		if err := rows.Scan(&c.CharacterID, &c.Name, &c.Rarity); err != nil {
			return nil, err
		}
		characters = append(characters, &c)
	}

	return characters, nil
}
