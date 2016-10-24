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

func calcNext(XI, YI, EX int, POS string, g [][]int64) string {
	block := g[YI][XI]
	fmt.Fprintf(os.Stderr, "%v\n", block)
	switch block {
	case 1, 3, 7, 8, 9, 12, 13:
		return fmt.Sprintf("%d %d", XI, YI+1)
	case 2:
		{
			if POS == "LEFT" {
				return fmt.Sprintf("%d %d", XI+1, YI)
			}
			return fmt.Sprintf("%d %d", XI-1, YI)
		}
	case 4:
		{
			if POS == "TOP" {
				return fmt.Sprintf("%d %d", XI-1, YI)
			}
			return fmt.Sprintf("%d %d", XI, YI+1)
		}
	case 5:
		{
			if POS == "TOP" {
				return fmt.Sprintf("%d %d", XI+1, YI)
			}
			return fmt.Sprintf("%d %d", XI, YI+1)
		}
	case 6:
		{
			if POS == "TOP" {
				if EX-XI > 0 {
					return fmt.Sprintf("%d %d", XI-1, YI)
				}
				return fmt.Sprintf("%d %d", XI+1, YI)
			} else if POS == "RIGHT" {
				return fmt.Sprintf("%d %d", XI-1, YI)
			}
			return fmt.Sprintf("%d %d", XI+1, YI)
		}
	case 10:
		return fmt.Sprintf("%d %d", XI-1, YI)
	case 11:
		return fmt.Sprintf("%d %d", XI+1, YI)
	}
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// W: number of columns.
	// H: number of rows.
	var W, H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W, &H)

	grid := make([][]int64, H, H)
	for i := range grid {
		grid[i] = make([]int64, W, W)
	}

	for i := 0; i < H; i++ {
		scanner.Scan()
		LINE := strings.Split(scanner.Text(), " ") // represents a line in the grid and contains W integers. Each integer represents one room of a given type.
		fmt.Fprintf(os.Stderr, "%s\n", LINE)
		for j, val := range LINE {
			grid[i][j], _ = strconv.ParseInt(val, 10, 32)
		}
	}
	//fmt.Fprintf(os.Stderr, "%s\n", grid)
	// EX: the coordinate along the X axis of the exit (not useful for this first mission, but must be read).
	var EX int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &EX)

	for {
		var XI, YI int
		var POS string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &XI, &YI, &POS)

		output := calcNext(XI, YI, EX, POS, grid)

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// One line containing the X Y coordinates of the room in which you believe Indy will be on the next turn.
		fmt.Println(output)
	}
}
