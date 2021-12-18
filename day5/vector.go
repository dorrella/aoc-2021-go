package main

import (
	"strconv"
	"strings"
)

// wrapper for using directions
type Direction int

//diagonals are not valid for part 1
// todo change naming convention from UpLeft to DirUpLeft
const (
	Invalid   Direction = 0
	Up                  = 1
	Right               = 2
	Down                = 3
	Left                = 4
	UpLeft              = 5
	UpRight             = 6
	DownLeft            = 7
	DownRight           = 8
)

// simple point
type Point struct {
	X int
	Y int
}

// vector
type Vect struct {
	Start     Point
	Distance  int
	Direction Direction
}

// create a point from string
func NewPoint(s string) Point {
	tokens := strings.Split(s, ",")
	if len(tokens) != 2 {
		panic("bad point")
	}

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic("bad point")
	}

	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic("bad point")
	}

	return Point{x, y}
}

// create vector from line of input file
func NewVect(s string) Vect {
	tokens := strings.Fields(s)
	if len(tokens) != 3 {
		panic("bad input")
	}

	//tokens[1] is "->"
	p1 := NewPoint(tokens[0])
	p2 := NewPoint(tokens[2])
	dir, dist := p1.GetVector(p2)

	return Vect{p1, dist, dir}
}

// Dir and Magnitude of two points
func (p *Point) GetVector(other Point) (Direction, int) {
	var dir Direction
	var dist, dx, dy int

	dx = 0
	dy = 0

	if p.X == other.X {
		if p.Y == other.Y {
			panic("didn't move?")
		}
		if p.Y > other.Y {
			dir = Up
			dist = p.Y - other.Y
		} else {
			dir = Down
			dist = other.Y - p.Y
		}
	} else if p.Y == other.Y {
		if p.X > other.X {
			dir = Left
			dist = p.X - other.X
		} else {
			dir = Right
			dist = other.X - p.X
		}
	} else {
		if p.X > other.X && p.Y > other.Y {
			dir = UpLeft
			dx = p.X - other.X
			dy = p.Y - other.Y

		} else if p.X > other.X && p.Y <= other.Y {
			dir = DownLeft
			dx = p.X - other.X
			dy = other.Y - p.Y
		} else if p.X <= other.X && p.Y > other.Y {
			dir = UpRight
			dx = other.X - p.X
			dy = p.Y - other.Y

		} else if p.X <= other.X && p.Y <= other.Y {
			dir = DownRight
			dx = other.X - p.X
			dy = other.Y - p.Y
		}

		if dx != dy {
			panic("bad diag vector")
		}
		dist = dx
	}
	return dir, dist
}

//used to filter input for part 1
func (v *Vect) IsDiag() bool {
	switch v.Direction {
	case Up:
		return false
	case Down:
		return false
	case Left:
		return false
	case Right:
		return false
	}

	return true
}
