package solver

import "fmt"

func prettyPrint(state [size][size]rune) {
	for i := range state {
		for j := range state[i] {
			fmt.Printf("%v ", state[i][j])
		}
		fmt.Println()
	}
}
