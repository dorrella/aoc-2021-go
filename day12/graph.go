package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Node struct {
	Name    string
	Nodes   map[string]*Node
	Visited bool
}

func NewNode(name string) *Node {
	return &Node{
		Name:    name,
		Nodes:   make(map[string]*Node),
		Visited: false,
	}
}

func (n *Node) IsBig() bool {
	if n.Name == "start" || n.Name == "end" {
		return false
	}

	r := rune(n.Name[0])
	if !unicode.IsLetter(r) {
		panic("bad name")
	}

	return unicode.IsUpper(r)
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{make(map[string]*Node)}
}

func (g *Graph) Update(s string) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 2 {
		panic("bad input")
	}

	//create from
	from_str := tokens[0]
	from, ok := g.Nodes[from_str]
	if !ok {
		from = NewNode(from_str)
		g.Nodes[from_str] = from
	}

	//create to
	to_str := tokens[1]
	to, ok := g.Nodes[to_str]
	if !ok {
		to = NewNode(to_str)
		g.Nodes[to_str] = to
	}

	//add to to from
	_, ok = from.Nodes[to_str]
	if !ok {
		from.Nodes[to_str] = to
	}

	//add from to to
	_, ok = to.Nodes[from_str]
	if !ok {
		to.Nodes[from_str] = from
	}
}

func (g *Graph) Walk(n *Node, path string) int {
	//fmt.Printf("%s: %s", n.Name, path)
	if !n.IsBig() {
		n.Visited = true
	}

	var new_path string
	if path == "" {
		new_path = n.Name
	} else {
		new_path = fmt.Sprintf("%s,%s", path, n.Name)
	}

	ret := 0
	if n.Name == "end" {
		//fmt.Println(new_path)
		ret = 1
	} else {
		for _, child := range n.Nodes {
			if child.IsBig() || child.Visited == false {
				ret += g.Walk(child, new_path)
			}
		}
	}

	if !n.IsBig() {
		n.Visited = false
	}
	return ret
}

func (g *Graph) Part1() {
	start, ok := g.Nodes["start"]
	if !ok {
		panic("bad input")
	}

	sum := g.Walk(start, "")
	fmt.Printf("total: %d\n", sum)
}
