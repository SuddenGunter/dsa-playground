package main

import (
	"fmt"
	"strings"
)

// public static void main(String[] args)
// { // Create a queue and enqueue/dequeue strings.
// 	Queue<String> q = new Queue<String>();
// 	while (!StdIn.isEmpty())
// 	{
// 	String item = StdIn.readString();
// 	if (!item.equals("-"))
// 	q.enqueue(item);
// 	else if (!q.isEmpty()) StdOut.print(q.dequeue() + " ");
// 	}
// 	StdOut.println("(" + q.size() + " left on queue)");
// 	}

func main() {
	q := NewQueue[string]()
	for _, x := range input() {
		if x != "-" {
			q.Enqueue(x)
		} else if !q.Empty() {
			fmt.Printf("%v ", q.Dequeue())
		}
	}

	fmt.Printf("(%v left on stack)", q.Size())

	fmt.Printf("\niterator: ")
	for iter := q.Iteartor(); iter.HasNext(); {
		fmt.Printf("%v ", iter.Next())
	}
}

func input() []string {
	return strings.Split("to be or not to - be - - that - - - is", " ")
}
