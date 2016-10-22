package main

import "fmt"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type node struct {
	val  byte
	next map[byte]*node
}

type stack []*node

func (s *stack) pop() *node {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *stack) push(n *node) {
	*s = append(*s, n)
}

func (n *node) size() int {
	size := 0
	nstack := stack{n}
	for len(nstack) > 0 {
		currnode := nstack.pop()
		for _, node := range currnode.next {
			nstack.push(node)
		}
		size++
	}
	return size
}

var numbers = make(map[byte]*node)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var telephone string
		fmt.Scan(&telephone)
		var nodes []*node
		for i := 0; i < len(telephone); i++ {
			currnode := node{telephone[i], make(map[byte]*node)}
			nodes = append(nodes, &currnode)
		}
		if numbers[nodes[0].val] == nil {
			numbers[nodes[0].val] = nodes[0]
		}
		currnode := numbers[nodes[0].val]
		for i, node := range nodes {
			if i != 0 {
				if currnode.next[node.val] == nil {
					currnode.next[node.val] = node
				}
				currnode = currnode.next[node.val]
			}
		}
	}

	total := 0
	for _, node := range numbers {
		total += node.size()
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// The number of elements (referencing a number) stored in the structure.
	fmt.Println(total)
}
