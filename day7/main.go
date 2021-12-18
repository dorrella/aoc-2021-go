package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read file and return a list of Vectors
func ReadInput(file string) []int {
	nums := make([]int, 0)
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

			nums = append(nums, num)
		}
	}

	return nums
}

func main() {
	//sample
	crabs := ReadInput("sample.txt")
	cm := NewCrabManager(crabs)
	cm.Part1()
	cm = NewCrabManager(crabs)
	cm.Part2()

	fmt.Println()

	//real
	crabs = ReadInput("input.txt")
	cm = NewCrabManager(crabs)
	cm.Part1()
	cm = NewCrabManager(crabs)
	cm.Part2()
}
