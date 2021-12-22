package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}
type Fold struct {
	Vert bool
	Axis int
}

func NewFold(fold string) Fold {
	//strips "fold along "
	str := fold[11:]
	tokens := strings.Split(str, "=")
	if len(tokens) != 2 {
		panic("bad fold")
	}

	vert := false
	if tokens[0] == "y" {
		vert = true
	}

	axis, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	return Fold{vert, axis}
}

type Paper struct {
	Matrix [][]bool
	Steps  []Fold
	Rows   int
	Cols   int
}

func NewPaper(points []Point, folds []string, max_row, max_col int) *Paper {
	p := &Paper{
		Matrix: make([][]bool, max_row+1),
		Steps:  make([]Fold, 0),
		Rows:   max_row + 1,
		Cols:   max_col + 1,
	}

	for i := 0; i < max_row+1; i++ {
		p.Matrix[i] = make([]bool, max_col+1)
	}

	for _, point := range points {
		p.Matrix[point.Y][point.X] = true
	}

	for _, fold := range folds {
		p.Steps = append(p.Steps, NewFold(fold))
	}

	return p
}

func (p *Paper) PrintPaper() {
	for row := 0; row < len(p.Matrix); row++ {
		for col := 0; col < len(p.Matrix[row]); col++ {
			if p.Matrix[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (p *Paper) PrintPaperH(axis int) {
	for row := 0; row < len(p.Matrix); row++ {
		for col := 0; col < len(p.Matrix[row]); col++ {
			if col == axis {
				fmt.Printf("|")
				continue
			}
			if p.Matrix[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (p *Paper) PrintPaperV(axis int) {
	for row := 0; row < len(p.Matrix); row++ {
		for col := 0; col < len(p.Matrix[row]); col++ {
			if row == axis {
				fmt.Printf("-")
				continue
			}
			if p.Matrix[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (p *Paper) FoldHorizontal(axis int) {
	new_rows := p.Rows
	new_cols := axis
	new_matrix := make([][]bool, new_rows)
	//init new matrix
	for row := 0; row < new_rows; row++ {
		new_matrix[row] = make([]bool, new_cols)
		for col := 0; col < new_cols; col++ {
			new_matrix[row][col] = p.Matrix[row][col]

		}
	}

	//fold in the rest
	for row := 0; row < new_rows; row++ {
		for col := 1; col+axis < p.Cols; col++ {
			if p.Matrix[row][axis+col] {
				new_matrix[row][axis-col] = true
			}
		}
	}

	p.Cols = axis
	p.Matrix = new_matrix

}

func (p *Paper) FoldVertical(axis int) {
	new_rows := axis
	new_cols := p.Cols
	new_matrix := make([][]bool, new_rows)

	//init new matrix
	for row := 0; row < new_rows; row++ {
		new_matrix[row] = make([]bool, new_cols)
		for col := 0; col < new_cols; col++ {
			new_matrix[row][col] = p.Matrix[row][col]

		}
	}

	//fold in the rest
	for row := 1; row+axis < p.Rows; row++ {
		for col := 0; col < new_cols; col++ {
			if p.Matrix[axis+row][col] {
				new_matrix[axis-row][col] = true
			}
		}
	}

	p.Rows = axis
	p.Matrix = new_matrix

}

func (p *Paper) Score() int {
	sum := 0
	for row := 0; row < len(p.Matrix); row++ {
		for col := 0; col < len(p.Matrix[row]); col++ {
			if p.Matrix[row][col] {
				//only count misses
				sum += 1
			}
		}
	}
	return sum
}

func (p *Paper) Part2(debug bool) {
	for _, step := range p.Steps {
		if step.Vert {
			if debug {
				p.PrintPaperV(step.Axis)
			}
			p.FoldVertical(step.Axis)
		} else {
			if debug {
				p.PrintPaperH(step.Axis)
			}
			p.FoldHorizontal(step.Axis)
		}
		if debug {
			fmt.Println()
			p.PrintPaper()
			fmt.Println()
		}
	}
	p.PrintPaper()
}

func (p *Paper) Part1(debug bool) {
	step := p.Steps[0]
	if step.Vert {
		if debug {
			p.PrintPaperV(step.Axis)
		}
		p.FoldVertical(step.Axis)
	} else {
		if debug {
			p.PrintPaperH(step.Axis)
		}
		p.FoldHorizontal(step.Axis)
	}
	if debug {
		fmt.Println()
		p.PrintPaper()
		fmt.Println()
	}

	sum := p.Score()
	fmt.Printf("total: %d\n", sum)
}
