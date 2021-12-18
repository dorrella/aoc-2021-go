package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) []Vect {
	var vects []Vect
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		vects = append(vects, NewVect(text))
	}

	return vects
}

func main() {
	//sample
	sample_vects := ReadInput("sample.txt")
	board := Board{}
	board.Part1(sample_vects)
	board = Board{}
	board.Part2(sample_vects)

	fmt.Println()

	//real input
	real_vects := ReadInput("input.txt")
	board = Board{}
	board.Part1(real_vects)
	board = Board{}
	board.Part2(real_vects)
}
