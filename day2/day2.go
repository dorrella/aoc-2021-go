package main

/*
--- Day 2: Dive! ---

Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

    forward X increases the horizontal position by X units.
    down X increases the depth by X units.
    up X decreases the depth by X units.

Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2

Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

    forward 5 adds 5 to your horizontal position, a total of 5.
    down 5 adds 5 to your depth, resulting in a value of 5.
    forward 8 adds 8 to your horizontal position, a total of 13.
    up 3 decreases your depth by 3, resulting in a value of 2.
    down 8 adds 8 to your depth, resulting in a value of 10.
    forward 2 adds 2 to your horizontal position, a total of 15.

After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Forward int = 0
	Down        = 1
	Up          = 2
)

type Direction struct {
	Direction int
	Spaces    int
}

// aim is used only for part 2
type Position struct {
	X   int
	Y   int
	Aim int
}

func ReadInput(c chan<- Direction, r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) != 2 {
			panic("failed to parse line")
		}

		var dir int = Forward
		switch tokens[0] {
		case "forward":
			dir = Forward
		case "down":
			dir = Down
		case "up":
			dir = Up
		default:
			panic("bad line")
		}

		spaces, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}

		c <- Direction{dir, spaces}
	}

	close(c)
}

// process input for part 1
func ProcessInput(c <-chan Direction) Position {
	p := Position{0, 0, 0}
	for dir := range c {
		switch dir.Direction {
		case Forward:
			p.X += dir.Spaces
		case Down:
			p.Y += dir.Spaces
		case Up:
			p.Y -= dir.Spaces
		}
	}

	return p
}

// day2 process input
func ProcessInput2(c <-chan Direction) Position {
	p := Position{0, 0, 0}
	for dir := range c {
		switch dir.Direction {
		case Forward:
			p.X += dir.Spaces
			p.Y = p.Y + (p.Aim * dir.Spaces)
		case Down:
			p.Aim += dir.Spaces
		case Up:
			p.Aim -= dir.Spaces
		}
	}

	return p
}

// scan file and find puzzle answer
// when aim is true, do part 2 puzzle input
func ScanFile(file string, aim bool) {
	c := make(chan Direction)
	var pos Position

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	go ReadInput(c, f)

	if aim {
		pos = ProcessInput2(c)

	} else {
		pos = ProcessInput(c)
	}

	fmt.Printf("range: %d depth: %d\n", pos.X, pos.Y)
	fmt.Printf("multiplicand: %d\n", pos.X*pos.Y)
}

func main() {
	ScanFile("sample.txt", false)
	fmt.Println()
	ScanFile("input.txt", false)
	fmt.Println()

	//part 2
	ScanFile("sample.txt", true)
	fmt.Println()
	ScanFile("input.txt", true)
}
