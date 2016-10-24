package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type vals map[rune]int

var val = vals{
	'e': 1, 'a': 1, 'i': 1, 'o': 1, 'n': 1, 'r': 1, 't': 1, 'l': 1, 's': 1, 'u': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10}

type word string

func (w word) getPoints() int {
	sum := 0
	for _, char := range w {
		sum += val[char]
	}
	return sum
}

func (w word) isValid(v map[rune]int) bool {
	t := map[rune]int{}
	for _, char := range w {
		if _, ok := v[char]; !ok {
			return false
		} else if _, ok := t[char]; !ok {
			t[char] = 1
		} else {
			t[char]++
		}
		if t[char] > v[char] {
			return false
		}
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	words := []word{}
	valid := map[rune]int{}

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		W := word(scanner.Text())
		words = append(words, W)
	}
	scanner.Scan()

	LETTERS := scanner.Text()
	for _, char := range LETTERS {
		if _, ok := valid[char]; !ok {
			valid[char] = 0
		}
		valid[char]++
	}

	maxPoints := 0
	var bestWord word

	for _, word := range words {
		if word.isValid(valid) {
			p := word.getPoints()
			if p > maxPoints {
				maxPoints = p
				bestWord = word
			}
		}
	}

	fmt.Fprintf(os.Stderr, "%v\n", valid)
	fmt.Println(bestWord) // Write answer to stdout
}
