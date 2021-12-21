package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Nodes
type Node struct {
	Name    string
	Nodes   map[string]*Node
	Visited int
}

func NewNode(name string) *Node {
	return &Node{
		Name:    name,
		Nodes:   make(map[string]*Node),
		Visited: 0,
	}
}

// Checks is Node qualifies as a big cave
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

// Graph
type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{make(map[string]*Node)}
}

//parse line and create/update nodes to build the graph
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

// walk the graph for part 1
func (g *Graph) Walk(n *Node, path string) int {
	if !n.IsBig() {
		n.Visited++
	}

	//update path
	var new_path string
	if path == "" {
		new_path = n.Name
	} else {
		new_path = fmt.Sprintf("%s,%s", path, n.Name)
	}

	ret := 0
	if n.Name == "end" {
		//check for end condition

		//fmt.Println(new_path)
		ret = 1
	} else {

		//maybe walk children
		for _, child := range n.Nodes {
			if child.IsBig() || child.Visited == 0 {
				ret += g.Walk(child, new_path)
			}
		}
	}

	if !n.IsBig() {
		n.Visited--
	}
	return ret
}

func (g *Graph) Walk2(n *Node, path string, doubled bool) int {
	//special exit case for not double visiting start
	//should not need one for "end"
	if n.Name == "start" && n.Visited > 0 {
		return 0
	}

	if !n.IsBig() {
		n.Visited++
	}

	//update path
	var new_path string
	if path == "" {
		new_path = n.Name
	} else {
		new_path = fmt.Sprintf("%s,%s", path, n.Name)
	}

	ret := 0
	if n.Name == "end" {
		//check for end condition

		//fmt.Println(new_path)
		ret = 1
	} else {
		//maybe walk children
		for _, child := range n.Nodes {
			big := child.IsBig()
			if big || child.Visited == 0 {
				ret += g.Walk2(child, new_path, doubled)
			} else if !big && !doubled {
				if child.Visited < 2 {
					ret += g.Walk2(child, new_path, true)
				}
			}
		}
	}

	if !n.IsBig() {
		n.Visited--
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

func (g *Graph) Part2() {
	start, ok := g.Nodes["start"]
	if !ok {
		panic("bad input")
	}

	sum := g.Walk2(start, "", false)
	fmt.Printf("deeper total: %d\n", sum)
}
