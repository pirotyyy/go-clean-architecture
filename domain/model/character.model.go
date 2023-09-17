package model

type Character struct {
	CharacterID int64  `json:"id"`
	Name        string `json:"name"`
	Rarity      string `json:"rarity"`
}

type CharacterListResponse struct {
	Characters []*UserCharacter `json:"characters"`
}
