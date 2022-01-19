package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *Cave {

	c := NewCave()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(scanner.Text())
		c.LoadRow(text)
	}

	return c
}

func main() {
	var c *Cave

	//sample
	c = ReadInput("sample.txt")
	c.Part2(true)

	fmt.Println()

	c = ReadInput("sample.txt")
	c.Part2(false)

	fmt.Println()

	//real
	c = ReadInput("input.txt")
	c.Part1(false)

	fmt.Println()
	c = ReadInput("input.txt")
	c.Part2(false)

}
