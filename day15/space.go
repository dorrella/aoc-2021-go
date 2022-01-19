package main

type Space struct {
	Value    int
	Distance int
	Row      int
	Col      int
}

func NewSpace(row, col, value int) *Space {
	return &Space{
		Value:    value,
		Distance: 0,
		Row:      row,
		Col:      col,
	}
}
