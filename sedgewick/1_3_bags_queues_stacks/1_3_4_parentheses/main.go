package main

import "fmt"

func main() {
	fmt.Printf("test for [()]{}{[()()]()}. expected: true. result: %v\n", valid("[()]{}{[()()]()}"))
	fmt.Printf("test for [(]). expected: false. result: %v\n", valid("[(])"))
}

func valid(str string) bool {
	stack := NewStack[rune]()

	for _, x := range str {
		if open(x) {
			stack.Push(x)
			continue
		}

		if stack.Empty() {
			return false
		}

		openPair := stack.Pop()

		if !matchPair(openPair, x) {
			return false
		}
	}

	if !stack.Empty() {
		return false
	}

	return true
}

func open(x rune) bool {
	switch x {
	case '{', '[', '(':
		return true
	default:
		return false
	}
}

func matchPair(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	default:
		return false
	}
}
