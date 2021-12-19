package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
)

// read file and return a list of Vectors
func ReadInput(file string) *Decoder {
	decoder := Decoder{
		Signals:    make([]*Signal, 0),
		DecoderMap: make(map[string]int),
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sig := Parse(scanner.Text())
		decoder.AddSignal(sig)
	}

	return &decoder
}

func main() {

	//sample
	d := ReadInput("sample.txt")
	//d.PrintScrambled()
	//fmt.Println()
	//d.PrintDecoded()
	//d.Part1()
	d.Signals[0].Decode()

	fmt.Println()

	//real
	d = ReadInput("input.txt")
	//d.PrintScrambled()
	//fmt.Println()
	//d.PrintDecoded()
	d.Part1()
}
