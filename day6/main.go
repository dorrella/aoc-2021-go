package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read file and return a list of Vectors
func ReadInput(file string) []uint {
	nums := make([]uint, 0)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		for _, t := range tokens {
			num, err := strconv.Atoi(t)
			if err != nil {
				panic("bad number")
			}

			nums = append(nums, uint(num))
		}
	}

	return nums
}

func main() {
	//sample
	sample_fish := ReadInput("sample.txt")
	gm := NewFishManager(sample_fish)
	gm.Part1(80)
	gm = NewFishManager(sample_fish)
	gm.Part1(256)

	fmt.Println()

	//real
	real_fish := ReadInput("input.txt")
	gm = NewFishManager(real_fish)
	gm.Part1(80)
	gm = NewFishManager(real_fish)
	gm.Part1(256)
}
