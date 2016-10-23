package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	last := int64(0)
	var dir byte
	lmax := int64(0)
	finalLoss := int64(0)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")

	fmt.Fprintf(os.Stderr, "%d\n%v\n", n, inputs)

	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 32)
		if v > last {
			if dir != 'U' {
				dir = 'U'
				tempLoss := last - lmax
				if tempLoss < finalLoss {
					finalLoss = tempLoss
				}
			}
		} else if v < last {
			if dir != 'D' {
				dir = 'D'
				if last > lmax {
					lmax = last
				}
			}
		}
		last = v
	}
	if (last - lmax) < finalLoss {
		finalLoss = last - lmax
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(finalLoss) // Write answer to stdout
}
