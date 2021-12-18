package main

import (
	"strconv"
	"strings"
)

type Direction int

const (
	Invalid Direction = 0
	Up                = 1
	Right             = 2
	Down              = 3
	Left              = 4
)

type Point struct {
	X int
	Y int
}

type Vect struct {
	Start     Point
	Distance  int
	Direction Direction
}

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

func NewVect(s string) Vect {
	tokens := strings.Fields(s)
	if len(tokens) != 3 {
		panic("bad input")
	}

	p1 := NewPoint(tokens[0])
	p2 := NewPoint(tokens[2])
	dir, dist := p1.GetVector(p2)

	return Vect{p1, dist, dir}
}

func (p *Point) GetVector(other Point) (Direction, int) {
	var dir Direction
	var dist int

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
	} else {
		//ignore (for now)?
		if p.Y != other.Y {
			dir = Invalid
			dist = -1
		} else if p.X > other.X {
			dir = Left
			dist = p.X - other.X
		} else {
			dir = Right
			dist = other.X - p.X
		}
	}
	return dir, dist
}
