package main

import (
	"fmt"
)

// todo
type FishManager struct {
	fish []uint
}

func NewFishManager(fish []uint) *FishManager {
	//needs 0-8, so size 9
	fm := FishManager{make([]uint, 9)}

	//only care about the number of fish in each bucket
	for _, f := range fish {
		fm.fish[f]++
	}

	return &fm
}

func (fm *FishManager) SimulateDay() {
	// these will breed fish at 8 and reset to 6
	breeders := fm.fish[0]

	// only go to 8 since we are doing i and i+1
	for i := 0; i < 8; i++ {
		fm.fish[i] = fm.fish[i+1]
	}

	//new fish
	fm.fish[8] = breeders
	//reset breeders timer on breeders
	fm.fish[6] += breeders
}

func (fm *FishManager) Part1(days int) {
	if days < 0 {
		panic("bad days")
	}
	for i := 0; i < days; i++ {
		fm.SimulateDay()
	}

	var sum uint = 0
	for _, num := range fm.fish {
		sum += num
	}

	fmt.Printf("there are %d fish\n", sum)
}
