package main

import (
	"fmt"
)

//magic
const (
	MaxUint uint = ^uint(0)
	MaxInt  int  = int(MaxUint >> 1)
)

type CrabManager struct {
	crabs []int
}

func NewCrabManager(crabs []int) *CrabManager {
	return &CrabManager{crabs}
}

// calculate Cost
//
// stops and returns -1 if cost exceeds limit
// calcs cost based on target depth
func (cm *CrabManager) Cost(target int, limit int) int {
	cost := 0
	for i := 0; i < len(cm.crabs); i++ {
		fuel := cm.crabs[i] - target
		if fuel < 0 {
			fuel = target - cm.crabs[i]
		}

		cost += fuel
		if cost > limit {
			return -1
		}
	}

	return cost
}

// calculate Cost for part 2
//
// stops and returns -1 if cost exceeds limit
// calcs cost based on target depth
func (cm *CrabManager) SlidingCost(target int, limit int) int {
	var cost int = 0
	for i := 0; i < len(cm.crabs); i++ {
		fuel := cm.crabs[i] - target
		if fuel < 0 {
			fuel = target - cm.crabs[i]
		}

		// ignore not moving
		if fuel == 0 {
			continue
		}

		fuel = fuel * (fuel + 1) / 2
		cost += fuel
		if cost > limit {
			return -1
		}
	}

	return cost
}

func (cm *CrabManager) Part1() {
	pos := 0
	min := MaxInt

	var max_crab, min_crab int
	max_crab = 0
	min_crab = MaxInt
	for _, c := range cm.crabs {
		if c > max_crab {
			max_crab = c
		}
		if c < min_crab {
			min_crab = c
		}
	}

	for i := min_crab; i <= max_crab; i++ {
		cost := cm.Cost(i, min)
		if cost >= 0 && cost < min {
			pos = i
			min = cost
		}
	}

	fmt.Printf("ideal depth is %d, cost %d\n", cm.crabs[pos], min)

}

func (cm *CrabManager) Part2() {
	pos := 0
	min := MaxInt

	var max_crab, min_crab int
	max_crab = 0
	min_crab = MaxInt
	for _, c := range cm.crabs {
		if c > max_crab {
			max_crab = c
		}
		if c < min_crab {
			min_crab = c
		}
	}

	for i := min_crab; i <= max_crab; i++ {
		cost := cm.SlidingCost(i, min)
		if cost >= 0 && cost < min {
			pos = i
			min = cost
		}
	}

	fmt.Printf("ideal depth for sliding cost is %d, cost %d\n", cm.crabs[pos], min)

}
