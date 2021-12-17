package main

import (
	"fmt"
	"strconv"
	"strings"
)

// game manager object
type GameManager struct {
	Draws  []int64
	Boards []Board
}

// init game manager
//
// draws: comma separated list of numbers to draw
// boards: list of Boards to check
func NewGameManager(draws string, boards []Board) *GameManager {
	d := make([]int64, 0)
	tokens := strings.Split(draws, ",")

	for _, token := range tokens {
		num, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			panic("can't convert num to int")
		}
		d = append(d, num)
	}

	return &GameManager{d, boards}

}


// draw and return one number
func (gm *GameManager) Draw() int64 {
	ret := gm.Draws[0]
	gm.Draws = gm.Draws[1:]
	return ret
}


// find "final score" for part 1
func (gm *GameManager) Part1() {
	for true {
		num := gm.Draw()
		for i := 0; i < len(gm.Boards); i++ {
			if gm.Boards[i].MarkNumber(num) {
				if gm.Boards[i].CheckBoard() {
					sum := gm.Boards[i].Evaluate()
					fmt.Printf("final score: %d\n", num*sum)
					return
				}
			}
		}
	}
}

// find "final score" for part 2
func (gm *GameManager) Part2() {
	var sum int64 = 0
	var new_boards []Board

	for true {
		num := gm.Draw()
		new_boards = make([]Board, 0)

		// mark boards and save boards that have not won
		for i := 0; i < len(gm.Boards); i++ {
			if gm.Boards[i].MarkNumber(num) && gm.Boards[i].CheckBoard() {
				sum = gm.Boards[i].Evaluate()
			} else {
				new_boards = append(new_boards, gm.Boards[i])
			}
		}

		// whatever the last sum is, is the last board to win
		if len(new_boards) < 1 {
			fmt.Printf("final score: %d\n", num*sum)
			return
		}

		// only iterate over unwon boards
		gm.Boards = new_boards
	}
}
