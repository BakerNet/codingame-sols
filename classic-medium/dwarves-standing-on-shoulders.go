package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type node struct {
	id       int
	adj      map[int]bool
	isSource bool
}

type graph map[int]*node

func (n *node) addAdj(g graph, i int) {
	n.adj[i] = true
	g[i].isSource = false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (n *node) traverse(g graph) int {
	len := 0
	for i, _ := range n.adj {
		len = max(len, g[i].traverse(g))
	}
	return len + 1
}

func main() {
	// n: the number of relationships of influence

	graph := make(graph)

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		if graph[x] == nil {
			graph[x] = &node{x, make(map[int]bool), true}
		}
		if graph[y] == nil {
			graph[y] = &node{y, make(map[int]bool), false}
		}
		graph[x].addAdj(graph, y)
	}

	var root *node

	for i, val := range graph {
		fmt.Fprintf(os.Stderr, "%v - %v\n", i, *val)
		if val.isSource {
			root = val
		}
	}

	len := root.traverse(graph)

	// The number of people involved in the longest succession of influences
	fmt.Println(len)
}
