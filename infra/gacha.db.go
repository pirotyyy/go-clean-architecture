package infra

import (
	"ca-tech/domain/model"
	"ca-tech/domain/repository"
	"context"
	"database/sql"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

type gachaRepository struct {
	DB *sql.DB
}

func NewGachaRepository(db *sql.DB) repository.GachaRepository {
	return &gachaRepository{
		DB: db,
	}
}

var characterList = []*model.Character{}

func (gr *gachaRepository) Draw(ctx context.Context, times int64, token string) (characters []*model.Character, err error) {
	const (
		selectListCommand = "SELECT character_id, name, rarity FROM game_character"
		selectUserCommand = "SELECT user_id FROM user WHERE token = ?"
		insertCommand     = "INSERT INTO user_character (user_id, character_id) VALUES (?, ?)"
	)

	if len(characterList) == 0 {
		rows, err := gr.DB.QueryContext(ctx, selectListCommand)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var char model.Character
			if err := rows.Scan(&char.CharacterID, &char.Name, &char.Rarity); err != nil {
				return nil, err
			}
			characterList = append(characterList, &char)
		}
	}

	for i := int64(0); i < times; i++ {
		rarity := selectRarity()

		targetCharacters := getTargetCharacters(rarity)

		selectedIndex := r.Intn(len(targetCharacters))
		selectedChar := targetCharacters[selectedIndex]
		characters = append(characters, selectedChar)
	}

	var userID string
	if err := gr.DB.QueryRowContext(ctx, selectUserCommand, token).Scan(&userID); err != nil {
		return nil, err
	}

	for _, character := range characters {
		_, err := gr.DB.ExecContext(ctx, insertCommand, userID, character.CharacterID)
		if err != nil {
			return nil, err
		}
	}

	return characters, nil
}

func selectRarity() string {
	num := r.Intn(100)

	switch {
	case num < 50:
		return "N"
	case num < 70:
		return "R"
	case num < 90:
		return "SR"
	default:
		return "SSR"
	}
}

func getTargetCharacters(rarity string) []*model.Character {
	var targetCharacters []*model.Character

	for _, char := range characterList {
		if char.Rarity == rarity {
			targetCharacters = append(targetCharacters, char)
		}
	}

	return targetCharacters
}
