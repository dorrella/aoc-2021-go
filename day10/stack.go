package main

import "fmt"

func ValidateToken(t Token) {
	if IsTokenOpen(t) || IsTokenClose(t) {
		return
	}
	panic("rune not valid")
}

type Stack struct {
	stack []Token
}

func NewStack() *Stack {
	return &Stack{make([]Token, 0)}
}

func (s *Stack) GetLen() int {
	return len(s.stack)
}

func (s *Stack) Push(t Token) {
	ValidateToken(t)
	s.stack = append(s.stack, t)
}

//ignore token when bool is false
func (s *Stack) Peek() (Token, bool) {
	l := s.GetLen()
	if l == 0 {
		return 'a', false
	}

	ret := s.stack[l-1]
	return ret, true
}

//ignore token when bool is false
func (s *Stack) Pop() (Token, bool) {
	l := s.GetLen() - 1
	if l == 0 {
		return ParenOpen, false
	}
	ret := s.stack[l]
	//actually remove item
	s.stack = s.stack[:l]
	return ret, true
}

func (s *Stack) Print() {
	out := ""
	for _, r := range s.stack {
		out = fmt.Sprintf("%s%c", out, r)
	}
	fmt.Println(out)
}
