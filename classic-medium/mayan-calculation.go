package main

import "fmt"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type num struct {
	lines []string
	val   int64
}

func getNums(table []string, L int) []num {
	nums := []num{}
	for i := 0; i < 20; i++ {
		lines := []string{}
		for _, line := range table {
			lines = append(lines, line[i*L:(i*L)+L])
		}
		nums = append(nums, num{lines, int64(i)})
	}
	return nums
}

func getVal(lines []string, nums []num, H int) int64 {
	for _, num := range nums {
		valid := true
		for i := 0; i < H; i++ {
			if num.lines[i] != lines[i] {
				valid = false
			}
		}
		if valid {
			return num.val
		}
	}
	return -1
}

func getStrings(num1 int64, nums []num, H int) []string {
	strings := make([]string, H, H)
	for _, num2 := range nums {
		if num1 == num2.val {
			for i := 0; i < H; i++ {
				strings[i] = num2.lines[i]
			}
		}
	}
	return strings
}

func getNumberVal(nums []num) int64 {
	total := int64(0)
	size := len(nums) - 1
	for i := size; i >= 0; i-- {
		fmt.Fprintf(os.Stderr, "%v - %v\n", int64(math.Pow(20, float64(size-i))), nums[i].val)
		total += int64(math.Pow(20, float64(size-i))) * nums[i].val
	}
	return total
}

func getStringVal(num int64, nums []num, H int) [][]string {
	strings := [][]string{}
	temp := num
	tnums := []int64{}
	for temp > 0 {
		tnums = append(tnums, temp%20)
		temp = temp / 20
	}
	fmt.Fprintf(os.Stderr, "%v\n", tnums)
	tnums = reverseInts(tnums)
	if len(tnums) == 0 {
		tnums = append(tnums, 0)
	}
	for _, tnum := range tnums {
		strings = append(strings, getStrings(tnum, nums, H))
	}
	return strings
}

func calc(num1 int64, num2 int64, op string) int64 {
	switch op {
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	}
	return -1
}

func reverseInts(input []int64) []int64 {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func main() {
	var L, H int64
	fmt.Scan(&L, &H)

	table := make([]string, H)
	nums := []num{}

	for i := int64(0); i < H; i++ {
		var numeral string
		fmt.Scan(&numeral)
		table[i] = numeral
	}
	nums = getNums(table, int(L))

	fmt.Fprintf(os.Stderr, "%v\n", nums)

	var S1 int64
	fmt.Scan(&S1)
	S1num := []num{}
	for i := int64(0); i < S1/H; i++ {
		lines := []string{}
		for j := int64(0); j < L; j++ {
			var num1Line string
			fmt.Scan(&num1Line)
			lines = append(lines, num1Line)
		}
		S1num = append(S1num, num{lines, getVal(lines, nums, int(H))})
	}

	fmt.Fprintf(os.Stderr, "S1 - %v\n", S1num)
	S1val := getNumberVal(S1num)

	var S2 int64
	fmt.Scan(&S2)
	S2num := []num{}
	for i := int64(0); i < S2/H; i++ {
		lines := []string{}
		for j := int64(0); j < L; j++ {
			var num2Line string
			fmt.Scan(&num2Line)
			lines = append(lines, num2Line)
		}
		S2num = append(S2num, num{lines, getVal(lines, nums, int(H))})
	}
	fmt.Fprintf(os.Stderr, "S2 - %v\n", S2num)
	S2val := getNumberVal(S2num)

	var operation string
	fmt.Scan(&operation)

	fmt.Fprintf(os.Stderr, "%v\n%v\n%v\n", S1val, S2val, operation)

	result := calc(S1val, S2val, operation)
	strings := getStringVal(result, nums, int(H))

	fmt.Fprintf(os.Stderr, "%v\n", strings)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	for _, inner := range strings {
		for _, s := range inner {
			fmt.Println(s)
		}
	}
	//fmt.Println(result)// Write answer to stdout
}
