package main

import (
	"fmt"
	"strings"
)

func main() {
	fs := NewStack[string](100)
	for _, x := range input() {
		if x != "-" {
			fs.Push(x)
		} else if !fs.Empty() {
			fmt.Printf("%v ", fs.Pop())
		}
	}

	fmt.Printf("(%v left on stack)", fs.Size())

	fmt.Printf("\niterator: ")
	for iter := fs.Iteartor(); iter.HasNext(); {
		fmt.Printf("%v ", iter.Next())
	}
}

func input() []string {
	return strings.Split("to be or not to - be - - that - - - is", " ")
}
