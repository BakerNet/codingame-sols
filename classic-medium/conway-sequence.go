package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func getConwayLine(start int, line int) []int {
	last := []int{start}
	next := []int{}

	for i := 2; i <= line; i++ {
		temp := last[0]
		count := 0
		for _, curr := range last {
			if temp != curr {
				next = append(next, count, temp)
				count = 0
				temp = curr
			}
			count++
		}
		next = append(next, count, temp)
		last = next
		next = []int{}

		//fmt.Fprintf(os.Stderr, "%v\n", last)
	}
	return last
}

func main() {
	var R int
	fmt.Scan(&R)

	var L int
	fmt.Scan(&L)

	answer := fmt.Sprintf("%v\n", getConwayLine(R, L))

	fmt.Fprintf(os.Stderr, "%v\n", answer)
	fmt.Println(answer[1 : len(answer)-2]) // Write answer to stdout
}
