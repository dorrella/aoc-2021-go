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

type BitCount struct {
	Ones  int
	Zeros int
}

type BitCounter struct {
	Bits []BitCount
}

func (bc *BitCounter) Gamma() int64 {
	bit_string := ""
	for _, bit := range bc.Bits {
		if bit.Ones > bit.Zeros {
			bit_string = fmt.Sprintf("%s%c", bit_string, '1')
		} else {
			bit_string = fmt.Sprintf("%s%c", bit_string, '0')
		}
	}

	ret, err := strconv.ParseInt(bit_string, 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

func (bc *BitCounter) Epsilon() int64 {
	bit_string := ""
	for _, bit := range bc.Bits {
		if bit.Ones < bit.Zeros {
			bit_string = fmt.Sprintf("%s%c", bit_string, '1')
		} else {
			bit_string = fmt.Sprintf("%s%c", bit_string, '0')
		}
	}

	ret, err := strconv.ParseInt(bit_string, 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

func ReadInput(c chan<- string, r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		c <- scanner.Text()
	}

	close(c)
}

func ProcessInput(c <-chan string) BitCounter {
	bc := BitCounter{make([]BitCount, 0)}

	for line := range c {
		fmt.Println(line)
		if len(line) > len(bc.Bits) {
			bc.Bits = make([]BitCount, len(line))
		}

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
	fmt.Println(bc)
	return bc
}

// scan file and find puzzle answer
// when aim is true, do part 2 puzzle input
func ScanFile(file string, aim bool) {
	c := make(chan string)
	var bc BitCounter

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	go ReadInput(c, f)

	if aim {
		//_ = ProcessInput2(c)

	} else {
		bc = ProcessInput(c)
	}

	fmt.Printf("power: %d\n", bc.Gamma()*bc.Epsilon())
}

func main() {
	ScanFile("sample.txt", false)
	fmt.Println()
	ScanFile("input.txt", false)
	fmt.Println()

	/*
		//part 2
		ScanFile("sample.txt", true)
		fmt.Println()
		ScanFile("input.txt", true)
	*/
}
