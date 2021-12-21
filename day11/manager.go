package main

import (
	"fmt"
	"strconv"
)

type OctopiManager struct {
	Octs [][]*Octopus
}

func NewOctopiManager() *OctopiManager {
	return &OctopiManager{make([][]*Octopus, 0)}
}

func (om *OctopiManager) AddOctopus(s string) {
	line := make([]*Octopus, len(s))

	for i, r := range s {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			panic("bad input")
		}

		line[i] = &Octopus{
			Energy:     val,
			HasFlashed: false,
		}
	}

	om.Octs = append(om.Octs, line)
}

func (om *OctopiManager) PrintBoard() {
	for row := 0; row < len(om.Octs); row++ {
		for col := 0; col < len(om.Octs[row]); col++ {
			fmt.Printf("%d", om.Octs[row][col].Energy)
		}
		fmt.Println()
	}
}

func (om *OctopiManager) ValidPos(r, c int) bool {
	if r < 0 || c < 0 {
		return false
	}

	if r >= len(om.Octs) || c >= len(om.Octs[r]) {
		return false
	}
	return true
}

func (om *OctopiManager) Flash(r, c int) int {
	oct := om.Octs[r][c]

	if oct.HasFlashed {
		return 0
	}

	if oct.Energy < 10 {
		return 0
	}

	//must flash
	oct.Energy = 0
	oct.HasFlashed = true
	flashes := 1
	for row := r - 1; row <= r+1; row++ {
		for col := c - 1; col <= c+1; col++ {
			if om.ValidPos(row, col) && !om.Octs[row][col].HasFlashed {
				om.Octs[row][col].Energy++
				flashes += om.Flash(row, col)
			}
		}
	}
	return flashes
}

func (om *OctopiManager) SimulateRound() int {
	flashes := 0
	for row := 0; row < len(om.Octs); row++ {
		for col := 0; col < len(om.Octs[row]); col++ {
			om.Octs[row][col].Energy++
			om.Octs[row][col].HasFlashed = false
		}
	}

	//this could be done in 1 go, but the implementation is
	for row := 0; row < len(om.Octs); row++ {
		for col := 0; col < len(om.Octs[row]); col++ {
			flashes += om.Flash(row, col)
		}
	}

	return flashes
}

func (om *OctopiManager) Part1(debug bool) {
	flashes := 0
	for i := 1; i <= 100; i++ {
		flashes += om.SimulateRound()
		if debug {
			fmt.Printf("step %d flashes: %d\n", i, flashes)
			om.PrintBoard()
		}

	}

	fmt.Printf("flashes: %d\n", flashes)
}

func (om *OctopiManager) Part2(debug bool) {
	//assume uniform size
	size := len(om.Octs) * len(om.Octs[0])
	for round := 1; true; round++ {
		flashes := om.SimulateRound()
		if debug {
			fmt.Printf("step %d flashes: %d\n", round, flashes)
			om.PrintBoard()
		}

		if flashes == size {
			fmt.Printf("all flashes (%d) on round: %d\n", flashes, round)
			return
		}

	}
	panic("bad input")
}
