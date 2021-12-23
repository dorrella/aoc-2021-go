package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file and return a list of Vectors
func ReadInput(file string) *Polymer {
	p := NewPolymer()
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	if !scanner.Scan() {
		panic("bad input")
	}
	p.Template = scanner.Text()
	//blank line
	if !scanner.Scan() {
		panic("bad input")
	}
	_ = scanner.Text()
	for scanner.Scan() {
		//g.Update(scanner.Text())
		p.AddRule(scanner.Text())
	}
	return p
}

func main() {
	var p *Polymer

	//sample
	p = ReadInput("sample.txt")
	p.Run(10, false)
	p.Run(40, false)
	//g.Part2()

	fmt.Println()

	//real
	//p = ReadInput("input.txt")
	//p.Run(40, false)
	//p.Part2()

}
