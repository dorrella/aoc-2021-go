package main

import "strconv"

func ParseString(s string) []int {
	ret := make([]int, 0)
	for _, r := range s {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			panic("bad number")
		}
		ret = append(ret, num)
	}

	return ret
}
