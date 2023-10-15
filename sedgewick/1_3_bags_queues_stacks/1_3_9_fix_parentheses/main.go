package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )"
	expResult := "( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )"

	result := formatWithSpaces(fixExpression(strings.ReplaceAll(input, " ", "")))
	fmt.Printf("input: %v\n", input)
	fmt.Printf("expected result: %v\n", expResult)
	fmt.Printf("result (%v): %v\n", resultMatch(result, expResult), result)
}

func resultMatch(a, b string) string {
	if a == b {
		return "OK"
	} else {
		return "DOES NOT MATCH"
	}
}

func fixExpression(exp string) string {
	operators := NewStack[rune]()
	operands := NewStack[string]()

	for _, x := range exp {
		if operator(x) {
			operators.Push(x)
			continue
		}

		if x == ')' {
			op := operators.Pop()
			r, l := operands.Pop(), operands.Pop()
			operands.Push(fmt.Sprintf("(%v%v%v)", l, string(op), r))
			continue
		}
		operands.Push(string(x))
	}

	return operands.Peek()
}

func operator(x rune) bool {
	switch x {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

func formatWithSpaces(exp string) string {
	sb := &strings.Builder{}
	for _, r := range exp {
		_, err := sb.WriteRune(r)
		if err != nil {
			panic(err)
		}

		_, err = sb.WriteRune(' ')
		if err != nil {
			panic(err)
		}
	}

	return strings.TrimSpace(sb.String())
}
