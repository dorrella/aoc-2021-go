package main

import (
	"fmt"
	"strconv"
)

// counts bits
type BitCount struct {
	Ones  int
	Zeros int
}

// wrapper for list of bits
type BitCounter struct {
	Bits []BitCount
}

// get int of greatest value bits
func (bc *BitCounter) Gamma() int64 {
	bit_string := bc.GammaString()

	ret, err := strconv.ParseInt(bit_string, 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

// get string of greatest value bits
func (bc *BitCounter) GammaString() string {
	ret := ""
	for _, bit := range bc.Bits {
		if bit.Ones >= bit.Zeros {
			ret = fmt.Sprintf("%s%c", ret, '1')
		} else {
			ret = fmt.Sprintf("%s%c", ret, '0')
		}
	}

	return ret
}

// get int of least value bits
func (bc *BitCounter) Epsilon() int64 {
	bit_string := bc.EpsilonString()

	ret, err := strconv.ParseInt(bit_string, 2, 64)
	if err != nil {
		panic("can't convert gamma to int")
	}

	return ret
}

// Get string of least common bits
func (bc *BitCounter) EpsilonString() string {
	ret := ""
	for _, bit := range bc.Bits {
		if bit.Ones < bit.Zeros {
			ret = fmt.Sprintf("%s%c", ret, '1')
		} else {
			ret = fmt.Sprintf("%s%c", ret, '0')
		}
	}

	return ret
}
