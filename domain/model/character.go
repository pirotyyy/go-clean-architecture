package model

type Rarity string

const (
	SSR Rarity = "SSR"
	SR  Rarity = "SR"
	R   Rarity = "R"
	N   Rarity = "N"
)

type Character struct {
	CharacterID int64  `json:"id"`
	Name        string `json:"name"`
	Rarity      Rarity `json:"rarity"`
}

type CharacterListResponse struct {
	Characters []*UserCharacter `json:"characters"`
}
