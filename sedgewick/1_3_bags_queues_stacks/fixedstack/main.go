package main

import (
	"fmt"
	"strings"
)

func main1() {
	fs := NewFixedCapacityStackOfStrings(100)
	for _, x := range input() {
		if x != "-" {
			fs.Push(x)
		} else if !fs.Empty() {
			fmt.Printf("%v ", fs.Pop())
		}
	}

	fmt.Printf("(%v left on stack)", fs.Size())
}

func main2() {
	fs := NewFixedCapacityStack[string](100)
	for _, x := range input() {
		if x != "-" {
			fs.Push(x)
		} else if !fs.Empty() {
			fmt.Printf("%v ", fs.Pop())
		}
	}

	fmt.Printf("(%v left on stack)", fs.Size())
}

func main() {
	fmt.Println()
	main1()
	fmt.Println()
	main2()
}

// public static void main(String[] args)
// {
// FixedCapacityStackOfStrings s;
// s = new FixedCapacityStackOfStrings(100);
// while (!StdIn.isEmpty())
// {
// String item = StdIn.readString();
// if (!item.equals("-"))
// s.push(item);
// else if (!s.isEmpty()) StdOut.print(s.pop() + " ");
// }
// StdOut.println("(" + s.size() + " left on stack)");
// }

func input() []string {
	return strings.Split("to be or not to - be - - that - - - is", " ")
}
