package main

import (
	"fmt"
)

type Board struct {
	board [1000][1000]int
}

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
	}
}

func (b *Board) Part1(vects []Vect) {
	for _, v := range vects {
		b.Update(v)
	}

	//now too big :(
	//b.PrintBoard()

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
