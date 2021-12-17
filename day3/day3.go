package main

/*
The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010

Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

So, the gamma rate is the binary number 10110, or 22 in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.)
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//read from file and pass text over channel
func ReadInput(c chan<- string, r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		c <- scanner.Text()
	}

	close(c)
}

//[]string is used for part 2
func ProcessInput(c <-chan string) (BitCounter, []string) {
	bc := BitCounter{make([]BitCount, 0)}
	out := make([]string, 0)

	for line := range c {
		out = append(out, line)

		//set bitcount to correct length, initializing all to zero
		if len(line) > len(bc.Bits) {
			bc.Bits = make([]BitCount, len(line))
		}

		//count values for epsilon/gamma
		for i, bit := range line {
			switch bit {
			case '0':
				bc.Bits[i].Zeros += 1
			case '1':
				bc.Bits[i].Ones += 1
			default:
				panic("bad input")
			}
		}
	}

	return bc, out
}

// part 2 process input
// in list of string read from wherever
// match is the string to match per the rules
// place is the idividual place to match
//
// could be done faster by tracking number of matches
func LifeSupport(in []string, match string, place int) (BitCounter, []string) {
	bc := BitCounter{make([]BitCount, 0)}
	out := make([]string, 0)

	for _, line := range in {
		//set bitcount to correct length, initializing all to zero
		if len(line) > len(bc.Bits) {
			bc.Bits = make([]BitCount, len(line))
		}

		if line[place] != match[place] {
			continue
		}

		out = append(out, line)

		//only count values for epsilon/gamma for taken lines
		for i, bit := range line {
			switch bit {
			case '0':
				bc.Bits[i].Zeros += 1
			case '1':
				bc.Bits[i].Ones += 1
			default:
				panic("bad input")
			}
		}
	}

	return bc, out
}

// filter for oxygen for part 2
func GetOxygen(in []string, match string) int64 {
	lines := in
	place := 0
	var bc BitCounter
	match_str := match

	// trust the input
	for len(lines) > 1 {
		bc, lines = LifeSupport(lines, match_str, place)
		match_str = bc.GammaString()
		place++
	}

	ret, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

//filter for co2 for part2
func GetCO2(in []string, match string) int64 {
	lines := in
	place := 0
	var bc BitCounter
	match_str := match

	// trust the input
	for len(lines) > 1 {
		bc, lines = LifeSupport(lines, match_str, place)
		match_str = bc.EpsilonString()
		place++
	}

	//make is usable
	ret, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

// scan file and find puzzle answer part 1
func GetPower(file string) {
	c := make(chan string)
	var bc BitCounter

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	go ReadInput(c, f)

	bc, _ = ProcessInput(c)

	fmt.Printf("power: %d\n", bc.Gamma()*bc.Epsilon())
}

// scan file and find puzzle answer part 1
func PrintLifeSupport(file string) {
	c := make(chan string)
	var bc BitCounter

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	go ReadInput(c, f)
	bc, lines := ProcessInput(c)

	oxygen := GetOxygen(lines, bc.GammaString())
	carbon := GetCO2(lines, bc.EpsilonString())

	fmt.Printf("life support: %d\n", oxygen*carbon)
}

func main() {
	GetPower("sample.txt")
	PrintLifeSupport("sample.txt")
	fmt.Println()

	GetPower("input.txt")
	PrintLifeSupport("input.txt")

}
