package main

import (
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

var leaves = make(map[*node]bool)
var relays = 0

type node struct {
	id     int
	adj    map[int]bool
	isLeaf bool
}

func (n *node) addAdj(i int) {
	n.adj[i] = true
}

func (n *node) checkLeaf() bool {
	if len(n.adj) == 1 {
		n.isLeaf = true
	}
	return n.isLeaf
}

func (n *node) removeFromGraph(graph map[int]*node) {
	for id, _ := range n.adj {
		delete(graph[id].adj, n.id)
		fmt.Fprintf(os.Stderr, "%d removed from %d - %v\n", n.id, graph[id].id, graph[id].adj)
	}
}

func removeLeaves(leaves map[*node]bool, graph map[int]*node) {
	for node, _ := range leaves {
		node.removeFromGraph(graph)
		delete(graph, node.id)
		delete(leaves, node)
	}
}

func findLeaves(graph map[int]*node) {
	for _, node := range graph {
		if node.checkLeaf() {
			leaves[node] = true
		}
	}
}

func main() {
	// n: the number of adjacency relations
	graph := make(map[int]*node)

	var n int
	fmt.Scan(&n)

	fmt.Fprintf(os.Stderr, "test\n")

	for i := 0; i < n; i++ {
		// xi: the ID of a person which is adjacent to yi
		// yi: the ID of a person which is adjacent to xi
		var xi, yi int
		fmt.Scan(&xi, &yi)
		if graph[xi] == nil {
			graph[xi] = &node{xi, make(map[int]bool), false}
		}
		if graph[yi] == nil {
			graph[yi] = &node{yi, make(map[int]bool), false}
		}
		graph[xi].addAdj(yi)
		graph[yi].addAdj(xi)
	}

	findLeaves(graph)
	for len(leaves) > 0 {
		fmt.Fprintf(os.Stderr, "%v\n", leaves)
		removeLeaves(leaves, graph)
		relays++

		findLeaves(graph)
	}
	// The minimal amount of steps required to completely propagate the advertisement
	fmt.Println(relays)
}
