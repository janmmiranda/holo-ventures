package game

type DamageType struct {
	name string
}

var (
	Piercing    = DamageType{name: "Piercing"}
	Slashing    = DamageType{name: "Slashing"}
	Bludgeoning = DamageType{name: "Bludgeoning"}
)
