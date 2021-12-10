package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/28/naive/solver"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	haystack := "hello"
	needle := "ll"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 2, res)
}

func TestExample2(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, -1, res)
}

func TestExample3(t *testing.T) {
	haystack := ""
	needle := ""

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 0, res)
}
