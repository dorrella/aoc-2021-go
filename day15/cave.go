package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Cave struct {
	matrix [][]*Space
	heads  []*Space
}

func NewCave() *Cave {
	m := make([][]*Space, 0)
	h := make([]*Space, 0)
	return &Cave{
		matrix: m,
		heads:  h,
	}
}

func (c *Cave) LoadRow(s string) {
	arr := make([]*Space, 0)
	row := len(c.matrix)
	for col, char := range s {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		s := NewSpace(row, col, num)
		arr = append(arr, s)
	}
	c.matrix = append(c.matrix, arr)
}

//expand to 5x cave size
//could dynamically find values, but
//our data structures mean we have to expand the whole thing
func (cave *Cave) Expand() {
	old_rows := len(cave.matrix)
	old_cols := len(cave.matrix[0])
	new_rows := old_rows * 5
	new_cols := old_cols * 5
	m := make([][]*Space, new_rows)
	for r := 0; r < new_rows; r++ {
		m[r] = make([]*Space, new_cols)
		for c := 0; c < new_cols; c++ {
			if r < old_rows && c < old_cols {
				m[r][c] = cave.matrix[r][c]
			} else {
				off_r := r / old_rows
				orig_r := r % old_rows

				off_c := c / old_cols
				orig_c := c % old_cols

				offset := off_c + off_r

				val := cave.matrix[orig_r][orig_c].Value + offset
				if val > 9 {
					val = val % 9
				}
				m[r][c] = NewSpace(r, c, val)
			}
		}
	}
	cave.matrix = m
}

//check if we will segault on access
func (c *Cave) IsValid(row, col int) bool {
	if row < 0 || col < 0 {
		return false
	}

	if row >= len(c.matrix) || col >= len(c.matrix[0]) {
		return false
	}
	return true
}

//checks for bottom right space
func (c *Cave) IsEnd(row, col int) bool {
	if row != len(c.matrix)-1 {
		return false
	}

	if col != len(c.matrix[0])-1 {
		return false
	}

	return true
}

//sorts the heads for path finding
func (c *Cave) SortHeads() {
	f := func(i, j int) bool {
		return c.heads[i].Distance < c.heads[j].Distance
	}

	sort.Slice(c.heads, f)
}

// something vaguely like djikstra's
func (c *Cave) Part1(debug bool) {
	head := c.matrix[0][1]
	head.Distance = head.Value
	c.heads = append(c.heads, head)

	head = c.matrix[1][0]
	head.Distance = head.Value
	c.heads = append(c.heads, head)

	for !c.Step(debug) {
		if debug {
			fmt.Print("heads: ")
			for _, h := range c.heads {
				fmt.Printf("%d ", h.Distance)
			}
			fmt.Println()
		}
	}

	fmt.Printf("safest path %d\n", c.heads[0].Distance)
}

// expand first
// then run part1
func (c *Cave) Part2(debug bool) {
	c.Expand()
	c.Part1(debug)

}

//steps a single step, returns true if found end
//distance is in c.head[0].Distance
func (c *Cave) Step(debug bool) bool {
	head := c.heads[0]
	c.heads = c.heads[1:]

	spaces := []struct{ row, col int }{
		{head.Row - 1, head.Col},
		{head.Row + 1, head.Col},
		{head.Row, head.Col - 1},
		{head.Row, head.Col + 1},
	}

	for _, s := range spaces {
		if !c.IsValid(s.row, s.col) {
			if debug {
				fmt.Printf("%d, %d not valid\n", s.row, s.col)
			}
			continue
		}

		if c.matrix[s.row][s.col].Distance == 0 {
			if debug {
				fmt.Printf("%d, %d found\n", s.row, s.col)
			}
			space := c.matrix[s.row][s.col]
			val := head.Distance + space.Value
			space.Distance = val

			if c.IsEnd(s.row, s.col) {
				if debug {
					fmt.Printf("%d, %d found end\n", s.row, s.col)
				}
				c.heads = c.heads[:1]
				c.heads[0] = space
				return true
			} else {
				c.heads = append(c.heads, space)
			}

		} else {
			if debug {
				fmt.Printf("%d, %d skipped\n", s.row, s.col)
			}
		}
	}

	c.SortHeads()
	return false
}
