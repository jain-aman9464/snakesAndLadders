package snakes_ladders

import (
	"math/rand"
	"time"
)

type Dice struct {
	numberOfDice int
}

func NewDice(numofdice int) Dice {
	return Dice{numberOfDice: numofdice}
}

func (d *Dice) rollDice() int {
	rand.Seed(time.Now().UnixNano())
	upperBound := 6 * d.numberOfDice
	lowerBound := d.numberOfDice

	return rand.Intn(upperBound-lowerBound) + lowerBound // Generates a number between 1 to 6
}
