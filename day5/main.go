package main

import (
	"bufio"
	//"fmt"
	"os"
)

func ReadInput(file string) []Vect {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	var vects []Vect

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		v := NewVect(text)
		if v.Direction != Invalid {
			vects = append(vects, v)
		}
	}

	return vects
}

func main() {
	//sample
	vects := ReadInput("sample.txt")
	board := Board{}
	board.Part1(vects)

	vects = ReadInput("input.txt")
	board = Board{}
	board.Part1(vects)
}
