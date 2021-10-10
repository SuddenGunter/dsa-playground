package solver

import lls "github.com/emirpasic/gods/stacks/linkedliststack"

func IsValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	stackP := lls.New()

	for _, char := range s {
		if isClose(char) {
			v, ok := stackP.Pop()
			if !ok {
				return false
			}

			if v.(rune) != getOpen(char) {
				return false
			}

			continue
		}

		stackP.Push(char)
	}

	return stackP.Empty()
}

func getOpen(char rune) rune {
	switch char {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return '-'
	}
}

func isClose(char rune) bool {
	if char == ')' || char == '}' || char == ']' {
		return true
	}

	return false
}
