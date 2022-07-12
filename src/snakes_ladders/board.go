package snakes_ladders

import (
	"container/list"
	"fmt"
)

type GameBoard struct {
	dice                   Dice
	snakes                 list.List
	ladders                list.List
	nextTurn               []Player
	playersCurrentPosition map[int]int
	boardSize              int
}

func NewGameBoard(dice Dice,
	nextTurn []Player,
	snakes list.List,
	ladders list.List,
	playersCurrentPos map[int]int,
	boardSize int) *GameBoard {

	return &GameBoard{
		dice:                   dice,
		snakes:                 snakes,
		ladders:                ladders,
		nextTurn:               nextTurn,
		playersCurrentPosition: playersCurrentPos,
		boardSize:              boardSize,
	}
}

func (board *GameBoard) StartGame() {

	for len(board.nextTurn) > 1 {
		// Take the first player and let him throw
		player := board.nextTurn[0]
		board.nextTurn = board.nextTurn[1:]                                   // Remove the first player from the queue
		currentPosition := board.playersCurrentPosition[player.getPlayerID()] // find the position of the first player
		nextCell := currentPosition + board.dice.rollDice()
		temp := nextCell
		if nextCell > board.boardSize {
			board.nextTurn = append(board.nextTurn, player)
		} else if nextCell == board.boardSize {
			fmt.Println(player.getPlayerName() + " won the game")
		} else {
			nextCell, _ := board.findNextCell(board.snakes, nextCell, false)
			if temp != nextCell {
				fmt.Printf("%s Bitten by Snake present at: %d \n", player.getPlayerName(), temp)
			}

			nextCell, isLadder := board.findNextCell(board.ladders, nextCell, true)
			if temp != nextCell && isLadder {
				fmt.Printf("%s Got ladder present at: %d \n", player.getPlayerName(), temp)
			}

			if nextCell == board.boardSize {
				fmt.Printf("%s won the game", player.getPlayerName())
			} else {
				board.playersCurrentPosition[player.getPlayerID()] = nextCell
				fmt.Printf("%s is at position %d \n", player.getPlayerName(), nextCell)
				board.nextTurn = append(board.nextTurn, player)
			}
		}
	}
}

func (board *GameBoard) findNextCell(l list.List, nextCell int, isLadder bool) (int, bool) {
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		val := ele.Value.(Jumper)
		if val.startPoint == nextCell {
			nextCell = val.endPoint
			return nextCell, isLadder == true
		}
	}

	return nextCell, false
}
