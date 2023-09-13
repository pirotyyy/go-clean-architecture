package model

type Character struct {
	CharacterID string `json:"id"`
	Name        string `json:"name"`
	Rarity      string `json:"rarity"`
}

type CharacterListResponse struct {
	Characters []*UserCharacter `json:"characters"`
}

type UserCharacter struct {
	UserCharacterID string `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
}
