package main

import (
	"fmt"
	"os"
	"sort"
)

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int
	fmt.Scan(&N)

	budgets := make([]int, N)
	payments := make([]int, N)

	var C int
	fmt.Scan(&C)

	for i := 0; i < N; i++ {
		var B int
		fmt.Scan(&B)
		budgets[i] = B
	}
	sort.Ints(budgets)

	for i, val := range budgets {
		evenC := C / (len(budgets) - i)
		fmt.Fprintf(os.Stderr, "%d - %d\n", evenC, val)
		if val < evenC {
			payments[i] = val
			C = C - val
		} else {
			payments[i] = evenC
			C = C - evenC
		}

	}

	impossible := false

	if C > 0 {
		impossible = true
	}

	//fmt.Fprintln(os.Stderr, "Debug messages...")
	if impossible {
		fmt.Println("IMPOSSIBLE") // Write answer to stdout
	} else {
		for _, val := range payments {
			fmt.Printf("%d\n", val)
		}
	}
}
