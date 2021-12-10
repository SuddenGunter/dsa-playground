package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/28/kmp/solver"
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

func TestExample4(t *testing.T) {
	haystack := "out-of-ra"
	needle := "out-of-range"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, -1, res)
}

func TestExample5(t *testing.T) {
	haystack := "abc"
	needle := "b"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 1, res)
}

func TestExample6(t *testing.T) {
	haystack := "hellhello"
	needle := "hello"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 4, res)
}

func TestExample7(t *testing.T) {
	haystack := "qweqweqweout-of-ra"
	needle := "out-of-range"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, -1, res)
}

func TestExample8(t *testing.T) {
	haystack := "mississippi"
	needle := "issip"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 4, res)
}

func TestExample9(t *testing.T) {
	haystack := "mississippi"
	needle := "pi"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 9, res)
}

func TestExample10(t *testing.T) {
	haystack := "mississippi"
	needle := "issipi"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, -1, res)
}

func TestExample11(t *testing.T) {
	haystack := "mississippi"
	needle := "sippi"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 6, res)
}

func TestExample12(t *testing.T) {
	haystack := "aabaaabaaac"
	needle := "aabaaac"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 4, res)
}

func TestExample13(t *testing.T) {
	haystack := "xacacabacacabacacacx"
	needle := "acacabacacabacacac"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, 1, res)
}

func TestExample14(t *testing.T) {
	haystack := "xacacabacacaxbacacacx"
	needle := "acacabacacabacacac"

	res := solver.StrStr(haystack, needle)

	assert.Equal(t, -1, res)
}
