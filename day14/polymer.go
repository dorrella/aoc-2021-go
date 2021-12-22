package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	First  rune
	Second rune
	Insert rune
}

type Polymer struct {
	Template string
	Elements map[rune]int
	Rules    []Rule
}

func NewPolymer() *Polymer {
	return &Polymer{
		Template: "",
		Elements: make(map[rune]int),
		Rules:    make([]Rule, 0),
	}
}

func (p *Polymer) AddRule(s string) {
	tokens := strings.Split(s, " -> ")
	if len(tokens) != 2 {
		panic("bad rule")
	}
	if len(tokens[0]) != 2 {
		panic("bad rule")
	}
	if len(tokens[1]) != 1 {
		panic("bad rule")
	}

	r := Rule{
		First:  []rune(tokens[0])[0],
		Second: []rune(tokens[0])[1],
		Insert: []rune(tokens[1])[0],
	}

	p.Rules = append(p.Rules, r)
}

func (p *Polymer) MatchRule(first, second rune) (rune, bool) {
	for _, r := range p.Rules {
		if first == r.First && second == r.Second {
			return r.Insert, true
		}
	}

	return 'a', false
}

func (p *Polymer) Step(in []rune) []rune {
	p.Elements = make(map[rune]int)
	ret := make([]rune, 0)

	for i := 0; i < len(in); i++ {
		first := in[i]
		_, found := p.Elements[first]
		if found {
			p.Elements[first]++
		} else {
			p.Elements[first] = 1
		}
		ret = append(ret, first)
		if i == len(in)-1 {
			break
		}

		second := in[i+1]
		match, ok := p.MatchRule(first, second)
		if ok {
			_, found := p.Elements[match]
			if found {
				p.Elements[match]++
			} else {
				p.Elements[match] = 1
			}
			ret = append(ret, match)
		}

	}

	return ret
}

func (p *Polymer) Part1(debug bool) {
	in := []rune(p.Template)
	if debug {
		fmt.Printf("token: %s\n", string(in))
	}
	for i := 1; i <= 10; i++ {
		in = p.Step(in)
		if debug {
			fmt.Printf("step %d: %s\n", i, string(in))
		}
	}

	if debug {
		fmt.Println(p.Elements)
	}

	max := -1
	min := 999999
	for _, v := range p.Elements {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Printf("%d - %d = %d\n", max, min, max-min)

}
