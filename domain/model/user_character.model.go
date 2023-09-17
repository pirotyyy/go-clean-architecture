package model

type UserCharacter struct {
	UserCharacterID int64  `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
}
