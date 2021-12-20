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

func IsTokenOpen(token Token) bool {
	l := []Token{ParenOpen, SquareOpen, CurlyOpen, AngleOpen}
	for _, t := range l {
		if t == token {
			return true
		}
	}
	return false
}

func IsTokenClose(token Token) bool {
	l := []Token{ParenClose, SquareClose, CurlyClose, AngleClose}
	for _, t := range l {
		if t == token {
			return true
		}
	}
	return false
}

func GetTokenMatch(t Token) Token {
	for k, v := range TokenMatches {
		if k == t {
			return v
		}
	}
	panic("unknown token")
}

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
