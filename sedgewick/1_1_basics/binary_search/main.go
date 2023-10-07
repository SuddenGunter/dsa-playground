package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	allowList := readInts("allow.txt")
	input := readInts("input.txt")

	sort.Ints(allowList)

	for _, x := range input {
		if rank(x, allowList) == -1 {
			fmt.Println(x)
		}
	}
}

func readInts(filename string) []int {
	f, err := os.ReadFile(filename)
	if err != nil {
		fatal(err)
	}

	lines := strings.Split(string(f), "\n")
	nums := make([]int, len(lines))

	for i := range lines {
		nums[i], err = strconv.Atoi(lines[i])
		if err != nil {
			fatal(err)
		}
	}

	return nums
}

func fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func rank(x int, allowList []int) int {
	lo, hi := 0, len(allowList)-1
	mid := (hi-lo)/2 + lo
	for hi >= lo {
		if allowList[mid] == x {
			return mid
		}

		if allowList[mid] > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}

		mid = (hi-lo)/2 + lo
	}

	return -1
}
