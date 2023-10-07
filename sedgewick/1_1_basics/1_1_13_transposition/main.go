package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")

	prettyPrint("input:", input)

	result := make([][]int, len(input[0]))
	for i := range result {
		result[i] = make([]int, len(input))
	}

	for i := range input {
		for j, x := range input[i] {
			result[j][i] = x
		}
	}

	prettyPrint("result:", result)
}

func readInput(filename string) [][]int {
	f, err := os.ReadFile(filename)
	if err != nil {
		fatal(err)
	}

	lines := strings.Split(string(f), "\n")

	result := make([][]int, 0, len(lines))

	for i := range lines {

		l := strings.Split(lines[i], " ")
		nums := make([]int, len(l))
		for j, x := range l {
			nums[j], err = strconv.Atoi(x)
			if err != nil {
				fatal(err)
			}
		}

		result = append(result, nums)
	}

	return result
}

func prettyPrint(label string, array [][]int) {
	fmt.Println(label)
	for i := range array {
		for j := range array[i] {
			fmt.Printf("%v ", array[i][j])
		}
		fmt.Print("\n")
	}
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}
