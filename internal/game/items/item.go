package items

import "github.com/janmmiranda/holo-ventures/internal/game/dice"

type Item struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	IsWeapon         bool             `json:"isWeapon"`
	WeaponDetails    WeaponDetails    `json:"weaponDetails"`
	Degradable       bool             `json:"degradable"`
	DegradableTimes  int              `json:"degradableTimes"`
	IsEquippable     bool             `json:"isEquippable"`
	EquipmentDetails EquipmentDetails `json:"equipmentDetails"`
}

type WeaponDetails struct {
	AttackRoll      dice.Dice
	AttackType      string
	HitRoll         dice.Dice
	HasSpecial      bool
	SpecialResource string
	SpecialRoll     dice.Dice
}

type EquipmentDetails struct {
	ArmorModifier          int
	IncreaseStats          bool
	EquimentStatsModifiers []EquipmentStatModifier
}

type EquipmentStatModifier struct {
	StatName string
	Modifier int
}
