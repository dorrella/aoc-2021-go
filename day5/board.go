package main

import (
	"fmt"
)

//board
//hardcoded to 1000x1000
//todo dynamic size
type Board struct {
	board [1000][1000]int
}

//prints the board, but input is too large to be usable
//
//hard code for sample input?
func (b *Board) PrintBoard() {
	for row := 0; row < 1000; row++ {
		for col := 0; col < 1000; col++ {
			if b.board[row][col] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", b.board[row][col])
			}
		}
		fmt.Println()
	}
}

//update board with vector
func (b *Board) Update(vect Vect) {
	var i int
	x := vect.Start.X
	y := vect.Start.Y

	switch vect.Direction {
	case Up:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y-i][x]++
		}
	case Down:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y+i][x]++
		}
	case Left:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y][x-i]++
		}
	case Right:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y][x+i]++
		}
	case UpLeft:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y-i][x-i]++
		}
	case UpRight:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y-i][x+i]++
		}
	case DownLeft:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y+i][x-i]++
		}
	case DownRight:
		for i = 0; i <= vect.Distance; i++ {
			b.board[y+i][x+i]++
		}
	}
}

// solve input for part 1
func (b *Board) Part1(vects []Vect) {
	for _, v := range vects {
		if !v.IsDiag() {
			b.Update(v)
		}
	}

	var row, col, sum int
	sum = 0

	for row = 0; row < 1000; row++ {
		for col = 0; col < 1000; col++ {
			if b.board[row][col] >= 2 {
				sum++
			}
		}
	}

	fmt.Printf("danger zones: %d\n", sum)
}

// solve input for part 2
func (b *Board) Part2(vects []Vect) {
	for _, v := range vects {
		b.Update(v)
	}

	var row, col, sum int
	sum = 0

	for row = 0; row < 1000; row++ {
		for col = 0; col < 1000; col++ {
			if b.board[row][col] >= 2 {
				sum++
			}
		}
	}

	fmt.Printf("super danger zones: %d\n", sum)
}
