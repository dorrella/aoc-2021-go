package main

import (
	"bufio"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *Graph {
	g := NewGraph()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		g.Update(scanner.Text())
	}
	return g
}

func main() {
	var g *Graph

	//sample
	g = ReadInput("sample.txt")
	g.Part1()
	g = ReadInput("sample2.txt")
	g.Part1()
	g = ReadInput("sample3.txt")
	g.Part1()

	//fmt.Println()

	//real
	g = ReadInput("input.txt")
	g.Part1()
	//g = ReadInput("input.txt")
	//g.Part2(false)
}
