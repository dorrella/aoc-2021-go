package main

import "fmt"

type HeightMap struct {
	HM [][]int
}

func NewHeightMap() *HeightMap {
	hm := make([][]int, 0)
	return &HeightMap{hm}
}

func (hm *HeightMap) AddRow(row string) {
	hm.HM = append(hm.HM, ParseString(row))
}

func (hm *HeightMap) ValidPos(row, col int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(hm.HM) {
		return false
	}
	if col >= len(hm.HM[0]) {
		return false
	}

	return true
}

//return risk level
func (hm *HeightMap) CheckPos(row, col int) int {
	r := row - 1
	c := col
	if hm.ValidPos(r, c) && hm.HM[row][col] >= hm.HM[r][c] {
		return 0
	}

	r = row + 1
	if hm.ValidPos(r, c) && hm.HM[row][col] >= hm.HM[r][c] {
		return 0
	}

	r = row
	c = col - 1
	if hm.ValidPos(r, c) && hm.HM[row][col] >= hm.HM[r][c] {
		return 0
	}

	c = col + 1
	if hm.ValidPos(r, c) && hm.HM[row][col] >= hm.HM[r][c] {
		return 0
	}

	return hm.HM[row][col] + 1
}

func (hm *HeightMap) Part1() {
	sum := 0
	for row := range hm.HM {
		for col := range hm.HM[0] {
			sum += hm.CheckPos(row, col)
		}
	}

	fmt.Printf("total risk: %d\n", sum)
}
