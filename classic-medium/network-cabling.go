package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type house struct {
	x int
	y int
}

type houseStats struct {
	maxX int
	minX int
	maxY int
	minY int
	aveY int
}

func (h *houseStats) getStats(houses []house) {
	h.maxX = -1073741824
	h.minX = 1073741824
	h.maxY = -1073741824
	h.minY = 1073741824
	sumY := 0
	for _, house := range houses {
		if house.x > h.maxX {
			h.maxX = house.x
		}
		if house.x < h.minX {
			h.minX = house.x
		}
		if house.y > h.maxY {
			h.maxY = house.y
		}
		if house.y < h.minY {
			h.minY = house.y
		}
		sumY += house.y
	}
	h.aveY = sumY / len(houses)
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func findMainCable(stats houseStats, houses []house) (y int, len int) {
	len = abs(stats.maxX - stats.minX)
	minDiff := abs(stats.aveY - houses[0].y)
	optimalY := houses[0].y
	for _, house := range houses {
		if abs(stats.aveY-house.y) < minDiff {
			minDiff = abs(stats.aveY - house.y)
			optimalY = house.y
		}
	}
	return optimalY, len
}

func findSideCables(houses []house, y int) int {
	sumLen := 0
	for _, house := range houses {
		sumLen += abs(y - house.y)
	}
	return sumLen
}

func main() {
	var N int
	fmt.Scan(&N)

	houses := make([]house, N, N)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		houses[i] = house{X, Y}
	}

	houseStats := houseStats{}
	houseStats.getStats(houses)

	optimalY, mainLen := findMainCable(houseStats, houses)
	sumLens := findSideCables(houses, optimalY)

	fmt.Fprintf(os.Stderr, "%v\n%v\n%v - %v - %v\n", houses, houseStats, optimalY, mainLen, sumLens)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(mainLen + sumLens) // Write answer to stdout
}
