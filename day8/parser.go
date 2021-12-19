package main

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	chars := []rune(s)
	f := func(i, j int) bool {
		return chars[i] < chars[j]
	}

	sort.Slice(chars, f)
	return string(chars)
}

func Parse(s string) *Signal {
	tokens := strings.Split(s, "|")
	if len(tokens) != 2 {
		panic("could not split")
	}

	if len(tokens) != 2 {
		panic("bad input")
	}

	sig := NewSignal()

	in := ParseInput(tokens[0])
	out := ParseOutput(tokens[1])

	//sort first to prevent conflicts
	var sorted string
	for _, i := range in {
		sorted = SortString(i)
		sig.Input = append(sig.Input, sorted)
	}
	for _, o := range out {
		sorted = SortString(o)
		sig.Output = append(sig.Output, sorted)
	}
	//sig.Decode()

	return sig
}

func ParseInput(s string) []string {
	tokens := strings.Fields(s)
	if len(tokens) != 10 {
		panic("wrong input")
	}

	return tokens
}

func ParseOutput(s string) []string {
	tokens := strings.Fields(s)
	if len(tokens) != 4 {
		panic("wrong output")
	}

	return tokens
}
