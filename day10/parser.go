package main

import (
	"fmt"
	"sort"
)

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
					//only happens when first char is closing token
					panic("bad input")
				}
				if last == match {
					//remove matching tokens
					_, _ = stack.Pop()
				} else {
					sum += ScoreToken(t)
					break
				}
				continue
			}

			//otherwise push onto the stack
			stack.Push(t)
		}
	}
	fmt.Printf("total corruption: %d\n", sum)
}

func (p *Parser) Part2(debug bool) {
	scores := make([]int, 0)

	for _, in := range p.data {
		if debug {
			fmt.Println(in)
		}

		stack := NewStack()
		corrupted := false

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
					panic("bad input")
				}
				if last == match {
					// remove matching braces
					_, _ = stack.Pop()
				} else {
					corrupted = true
					break
				}
				continue
			}

			//otherwise push onto the stack
			stack.Push(t)
		}

		if !corrupted {
			missing := stack.CompletionString()
			score := ScoreCompletion(missing)
			scores = append(scores, score)
		}
	}

	//sort the scores
	f := func(i, j int) bool {
		return scores[i] < scores[j]
	}
	sort.Slice(scores, f)
	index := len(scores) / 2

	fmt.Printf("total corruption: %d\n", scores[index])
}
