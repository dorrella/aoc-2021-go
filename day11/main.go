package main

import (
	"bufio"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *OctopiManager {
	om := NewOctopiManager()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		om.AddOctopus(scanner.Text())
	}

	return om
}

func main() {
	var o *OctopiManager

	//sample
	o = ReadInput("sample.txt")
	o.Part1(false)
	o = ReadInput("sample.txt")
	o.Part2(false)

	//fmt.Println()

	//real
	o = ReadInput("input.txt")
	o.Part1(false)
	o = ReadInput("input.txt")
	o.Part2(false)
}
