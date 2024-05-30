package dice

import (
	"math"
	"math/rand"
)

type Dice struct {
	Sides int
	Count int
}

func (d Dice) Roll(mod int) int {
	res := mod
	for i := 0; i < d.Count; i++ {
		res += 1 + rand.Intn(d.Sides)
	}
	return res
}

func RollAdvantage(mod int, count int, d *Dice) int {
	return int(math.Max(float64(d.Roll(mod)), float64(d.Roll(mod))))
}

func RollDisadvantage(mod int, count int, d *Dice) int {
	return int(math.Min(float64(d.Roll(mod)), float64(d.Roll(mod))))
}
