package entities

import "github.com/janmmiranda/holo-ventures/internal/game/items"

type Entity struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	BaseStats         Stats        `json:"baseStats"`
	Inventory         []items.Item `json:"inventory"`
	Class             string       `json:"class"`
	Race              string       `json:"race"`
	ExpPoints         int          `json:"expPoints"`
	Inspiration       bool         `json:"inspiration"`
	InspirationRounds int          `json:"inspirationRounds"`
}

type Stats struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

func NewCharacter() *Entity {
	return &Entity{
		ID:   "rand",
		Name: "Leo",
		BaseStats: Stats{
			Strength:     14,
			Dexterity:    12,
			Constitution: 10,
			Intelligence: 8,
			Wisdom:       10,
			Charisma:     12,
		},
		Class:             "Warrior",
		Race:              "Human",
		ExpPoints:         0,
		Inspiration:       false,
		InspirationRounds: 0,
	}
}
