package main

type Token rune

const (
	ParenOpen   Token = '('
	ParenClose        = ')'
	SquareOpen        = '['
	SquareClose       = ']'
	CurlyOpen         = '{'
	CurlyClose        = '}'
	AngleOpen         = '<'
	AngleClose        = '>'
)

//const?
var TokenMatches = map[Token]Token{
	ParenOpen:   ParenClose,
	ParenClose:  ParenOpen,
	SquareOpen:  SquareClose,
	SquareClose: SquareOpen,
	CurlyOpen:   CurlyClose,
	CurlyClose:  CurlyOpen,
	AngleOpen:   AngleClose,
	AngleClose:  AngleOpen,
}

//opening brackets
func IsTokenOpen(token Token) bool {
	l := []Token{ParenOpen, SquareOpen, CurlyOpen, AngleOpen}
	for _, t := range l {
		if t == token {
			return true
		}
	}
	return false
}

//closing brackets
func IsTokenClose(token Token) bool {
	l := []Token{ParenClose, SquareClose, CurlyClose, AngleClose}
	for _, t := range l {
		if t == token {
			return true
		}
	}
	return false
}

//find complementary match
func GetTokenMatch(t Token) Token {
	for k, v := range TokenMatches {
		if k == t {
			return v
		}
	}
	panic("unknown token")
}

//generate score for part 1
func ScoreToken(t Token) int {
	switch t {
	case ParenClose:
		return 3
	case SquareClose:
		return 57
	case CurlyClose:
		return 1197
	case AngleClose:
		return 25137
	}
	panic("extra corrupted")
}

//generate score for part 2
func ScoreCompletion(s string) int {
	sum := 0
	for _, r := range s {
		t := Token(r)
		sum = sum * 5

		switch t {
		case ParenClose:
			sum += 1
		case SquareClose:
			sum += 2
		case CurlyClose:
			sum += 3
		case AngleClose:
			sum += 4
		default:
			panic("bad completion string")
		}
	}

	return sum
}
