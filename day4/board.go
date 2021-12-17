package main

import (
	"strconv"
	"strings"
)


// wrapper around Square object
type Square struct {
	Hit bool
	Num int64
}

// wrapper around 2d array of Squares
type Board struct {
	Values [5][5]Square
}

func NewBoard(in []string) Board {
	ret := Board{}
	for i, line := range in {
		tokens := strings.Fields(line)
		for j, token := range tokens {
			num, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				panic("can't convert num to int")
			}

			ret.Values[i][j].Num = num
		}
	}

	return ret
}

// Check board for victory condition
//
// returns true if bingo
func (b *Board) CheckBoard() bool {

	var winning_vert bool
	var winning_hor bool

	for i := 0; i < 5; i++ {

		winning_vert = true
		winning_hor = true

		for j := 0; j < 5; j++ {
			if !winning_vert || !b.Values[i][j].Hit {
				winning_vert = false
			}

			//swap i and j to check horizontally
			if !winning_hor || !b.Values[j][i].Hit {
				winning_hor = false
			}
		}

		if winning_hor {
			//won across
			return true
		}

		if winning_vert {
			return true
		}
	}

	return false
}

// check of number
//
// num, number to mark as hit, if able
//
// returns true if number found
func (b *Board) MarkNumber(num int64) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b.Values[row][col].Num == num {
				b.Values[row][col].Hit = true
				return true
			}
		}
	}

	return false
}

// Returns sum of non-hit numbers
func (b *Board) Evaluate() int64 {
	var sum int64 = 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b.Values[row][col].Hit {
				sum += b.Values[row][col].Num
			}
		}
	}
	return sum
}
