package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read file and return a list of Vectors
func ReadInput(file string) *Paper {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	points := make([]Point, 0)
	folds := make([]string, 0)
	max_row := 0
	max_col := 0
	doing_points := true
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			//switch between points and folds
			doing_points = false
		} else if doing_points {
			tokens := strings.Split(text, ",")
			if len(tokens) != 2 {
				panic("bad input")
			}

			col, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic("bad input")
			}
			if col > max_col {
				max_col = col
			}

			row, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("bad input")
			}
			if row > max_row {
				max_row = row
			}

			points = append(points, Point{X: col, Y: row})
		} else {
			folds = append(folds, text)
		}
		//g.Update(scanner.Text())
		//fmt.Println(scanner.Text())
	}

	p := NewPaper(points, folds, max_row, max_col)
	return p
}

func main() {
	var p *Paper

	//sample
	p = ReadInput("sample.txt")
	p.Part1(false)
	p = ReadInput("sample.txt")
	p.Part2(false)

	fmt.Println()

	//real
	p = ReadInput("input.txt")
	p.Part1(false)
	p = ReadInput("input.txt")
	p.Part2(false)

}
