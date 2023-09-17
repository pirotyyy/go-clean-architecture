package model

type GachaDrawRequest struct {
	Times int64 `json:"times"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	CharacterID int64  `json:"character_id"`
	Name        string `json:"name"`
	Rarity      string `json:"rarity"`
}
