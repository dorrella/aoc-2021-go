package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *HeightMap {
	var hm *HeightMap
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	hm = NewHeightMap()

	for scanner.Scan() {
		hm.AddRow(scanner.Text())
	}

	return hm
}

func main() {
	var hm *HeightMap

	//sample
	hm = ReadInput("sample.txt")
	hm.Part1()

	fmt.Println()

	//real
	hm = ReadInput("input.txt")
	hm.Part1()
}
