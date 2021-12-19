package main

import "fmt"

type Signal struct {
	Input      []string
	Output     []string
	DecoderMap map[string]int
}

func NewSignal() *Signal {
	return &Signal{
		Input:      make([]string, 0),
		Output:     make([]string, 0),
		DecoderMap: make(map[string]int),
	}
}

func (s *Signal) decodeEasy() []string {
	signals := append(s.Input, s.Output...)

	decoded := make([]string, 10)
	//not sure if needed
	for i := 0; i < 10; i++ {
		decoded[i] = ""
	}

	for _, signal := range signals {
		switch len(signal) {
		case 2:
			val, ok := s.DecoderMap[signal]
			if ok {
				if val != 1 {
					panic("double input")
				}
			} else {
				s.DecoderMap[signal] = 1
				decoded[1] = signal
			}
		case 3:
			val, ok := s.DecoderMap[signal]
			if ok {
				if val != 7 {
					panic("double input")
				}
			} else {
				s.DecoderMap[signal] = 7
				decoded[7] = signal
			}
		case 4:
			val, ok := s.DecoderMap[signal]
			if ok {
				if val != 4 {
					panic("double input")
				}
			} else {
				s.DecoderMap[signal] = 4
				decoded[4] = signal
			}
		case 7:
			val, ok := s.DecoderMap[signal]
			if ok {
				if val != 8 {
					panic("double input")
				}
			} else {
				s.DecoderMap[signal] = 8
				decoded[8] = signal
			}
		default:
			val, ok := s.DecoderMap[signal]
			if ok {
				if val != -1 {
					panic("double input")
				}
			} else {
				s.DecoderMap[signal] = -1
			}
		}
	}

	return decoded
}

func (s *Signal) decodeHard(cheat []string) {
	codex := map[rune]rune{
		'a': 'z',
		'b': 'z',
		'c': 'z',
		'd': 'z',
		'e': 'z',
		'f': 'z',
		'g': 'z',
	}

	codex['a'] = s.decodeA(cheat[7], cheat[1])
	fmt.Printf("found %c is %c\n", codex['a'], 'a')

	codex['c'], codex['d'], codex['e'] = s.decodeCDE(cheat[4], cheat[7], cheat[8])
	fmt.Printf("found %c is %c\n", codex['c'], 'c')
	fmt.Printf("found %c is %c\n", codex['d'], 'd')
	fmt.Printf("found %c is %c\n", codex['e'], 'e')

	codex['g'] = s.decodeG(codex['a'])
	fmt.Printf("found %c is %c\n", codex['g'], 'g')

	// set two
	tmp_str := fmt.Sprintf("%s%s%s%s%s", codex['a'], codex['c'], codex['d'], codex['e'], codex['g'])
	tmp_str = SortString(tmp_str)
	s.DecoderMap[tmp_str] = 2

	codex['b'], codex['f'] = s.decodeBF(tmp_str)
	fmt.Printf("found %c is %c\n", codex['b'], 'b')
	fmt.Printf("found %c is %c\n", codex['f'], 'f')

	//set three
	tmp_str = fmt.Sprintf("%s%s%s%s%s", codex['a'], codex['c'], codex['d'], codex['f'], codex['g'])
	tmp_str = SortString(tmp_str)
	s.DecoderMap[tmp_str] = 3

	//set five
	tmp_str = fmt.Sprintf("%s%s%s%s%s", codex['a'], codex['b'], codex['d'], codex['f'], codex['g'])
	tmp_str = SortString(tmp_str)
	s.DecoderMap[tmp_str] = 5

	for k, v := range codex {
		fmt.Printf("%c %c\n", k, v)
	}
}

func (s *Signal) decodeA(seven, one string) rune {
	runes_s := []rune(seven)
	runes_o := []rune(one)

	for _, r := range runes_s {
		// if not matched, it must be the top, a.
		if !isInRunes(runes_o, r) {
			return r
		}
	}

	panic("could not decode a")
}

func (s *Signal) decodeCDE(four, seven, eight string) (rune, rune, rune) {
	runes_f := []rune(four)
	runes_s := []rune(seven)
	runes_e := []rune(eight)

	d := 'z'
	e := 'z'
	c := 'z'

	for k := range s.DecoderMap {
		runes_k := []rune(k)
		if len(k) == 6 {
			// 0, 6, or 9
			if isSubRune(runes_k, runes_f) {
				// is 9, meaning the 1 extra rune is e
				s.DecoderMap[k] = 9
				er := extraRunes(runes_e, runes_k)
				e = er[0]
			} else {
				if isSubRune(runes_k, runes_s) {
					// is zero, so extra rune is d
					s.DecoderMap[k] = 6
					er := extraRunes(runes_e, runes_k)
					d = er[0]
				} else {
					// is six, so extra rune is c
					s.DecoderMap[k] = 0
					er := extraRunes(runes_e, runes_k)
					c = er[0]
				}
			}
		}
	}

	if d == 'z' || e == 'z' || c == 'z' {
		panic("could not decode signal")
	}

	return c, d, e
}

// decod value of a
func (s *Signal) decodeG(a rune) rune {
	four := ""
	nine := ""

	for k, v := range s.DecoderMap {
		if v == 4 {
			four = k
		} else if v == 9 {
			nine = k
		}
	}

	if four == "" || nine == "" {
		panic("can't find g")
	}

	//append a to make an almost 9
	r_nine := []rune(nine)
	r_four := []rune(four)
	r_four = append(r_four, a)

	er := extraRunes(r_nine, r_four)
	return er[0]
}

func (s *Signal) decodeBF(two string) (rune, rune) {
	one := ""
	eight := ""

	for k, v := range s.DecoderMap {
		if v == 1 {
			one = k
		} else if v == 8 {
			eight = k
		}
	}

	if one == "" || eight == "" {
		panic("lost one or eight")
	}

	r_two := []rune(two)
	r_eight := []rune(eight)
	r_one := []rune(one)

	er := extraRunes(r_eight, r_two)
	fmt.Printf("extra runes %c %c\n", er[0], er[1])
	if isInRunes(r_one, er[0]) {
		return er[1], er[0]
	}

	// 0 is b
	return er[0], er[1]

}

func (s *Signal) Decode() {
	cheat_sheet := s.decodeEasy()
	s.decodeHard(cheat_sheet)

}

// todo generics
func isInRunes(a []rune, b rune) bool {
	for _, r := range a {
		if b == r {
			return true
		}
	}
	return false
}

//little is a substring of big
func isSubRune(big, little []rune) bool {
	for _, r := range little {
		if !isInRunes(big, r) {
			return false
		}
	}
	return true
}

func extraRunes(extra, other []rune) []rune {
	ret := make([]rune, 0)
	for _, r := range extra {
		if !isInRunes(other, r) {
			ret = append(ret, r)
		}
	}
	return ret
}
