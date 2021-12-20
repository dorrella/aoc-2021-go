package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *Parser {
	parser := NewParser()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parser.LoadString(scanner.Text())
	}

	return parser
}

func main() {
	var p *Parser

	//sample
	p = ReadInput("sample.txt")
	p.Part1(false)
	p.Part2(false)

	fmt.Println()

	//real
	p = ReadInput("input.txt")
	p.Part1(false)
	p.Part2(false)
}
