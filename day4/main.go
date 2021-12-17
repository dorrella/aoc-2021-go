package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadInput(c chan<- string, r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		c <- scanner.Text()
	}

	close(c)
}

// load input into game manager object
func LoadBoard(file string) *GameManager {
	c := make(chan string)

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	go ReadInput(c, f)

	//draws
	draws := <-c
	// blank line before boards
	_ = <-c

	count := 0
	boards := make([]Board, 0)
	board_strings := make([]string, 5)
	for line := range c {
		if count == 5 {
			//this line is blank, so load the board
			boards = append(boards, NewBoard(board_strings))
			count = 0
			board_strings = make([]string, 5)
			continue
		}

		board_strings[count] = line
		count++
	}

	//last board is skipped by range :(
	boards = append(boards, NewBoard(board_strings))
	count = 0
	board_strings = make([]string, 0)

	return NewGameManager(draws, boards)
}

func main() {
	//sample
	gm := LoadBoard("sample.txt")
	gm.Part1()
	gm = LoadBoard("sample.txt")
	gm.Part2()

	fmt.Println()

	//real
	gm = LoadBoard("input.txt")
	gm.Part1()
	gm = LoadBoard("input.txt")
	gm.Part2()
}
