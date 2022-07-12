package main

import (
	"container/list"
	"github.com/tokopedia/test/snakesAndLadders/src/snakes_ladders"
)

func main() {
	dice := snakes_ladders.NewDice(1)
	p1 := snakes_ladders.NewPlayer("Albert", 1)
	p2 := snakes_ladders.NewPlayer("Pintoss", 2)
	allPlayers := []snakes_ladders.Player{*p1, *p2}

	snake1 := snakes_ladders.NewJumper(10, 2)
	snake2 := snakes_ladders.NewJumper(99, 12)
	snakes := list.New()
	snakes.PushBack(snake1)
	snakes.PushBack(snake2)

	ladder1 := snakes_ladders.NewJumper(5, 25)
	ladder2 := snakes_ladders.NewJumper(40, 89)

	ladders := list.New()
	ladders.PushBack(ladder1)
	ladders.PushBack(ladder2)

	playersCurrentPosition := map[int]int{
		1: 0,
		2: 0,
	}

	gb := snakes_ladders.NewGameBoard(dice, allPlayers, *snakes, *ladders, playersCurrentPosition, 100)

	gb.StartGame()
}
