package main

import "fmt"

type Parser struct {
	data []string
}

func NewParser() *Parser {
	return &Parser{make([]string, 0)}
}

func (p *Parser) LoadString(s string) {
	p.data = append(p.data, s)
}

func (p *Parser) Part1(debug bool) {
	sum := 0

	for _, in := range p.data {
		if debug {
			fmt.Println(in)
		}
		stack := NewStack()

		for _, r := range in {
			t := Token(r)
			if debug {
				stack.Print()
				fmt.Printf("%c\n", t)
			}
			if IsTokenClose(t) {
				match := GetTokenMatch(t)
				last, ok := stack.Peek()
				if !ok {
					panic("corrupted")
				}
				if last == match {
					_, _ = stack.Pop()
				} else {
					sum += ScoreToken(t)
					break
				}
				continue
			}

			stack.Push(t)
		}
	}
	fmt.Printf("total corruption: %d\n", sum)
}
